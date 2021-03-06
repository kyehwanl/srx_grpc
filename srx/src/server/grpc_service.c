

#include <stdio.h>
#include "server/grpc_service.h"
#include "server/command_handler.h"

#define HDR  "([0x%08X] GRPC_ServiceHandler): "


__attribute__((always_inline)) inline void printHex(int len, unsigned char* buff) 
{                                                                                 
  int i;                                                                          
  for(i=0; i < len; i++ )                                                         
  {                                                                               
      if(i%16 ==0) printf("\n");                                                  
      printf("%02x ", buff[i]);                                                   
  }                                                                               
  printf("\n");                                                                   
}  

        
static bool processHandshake_grpc(unsigned char *data, RET_DATA *rt)
{
  LOG(LEVEL_DEBUG, HDR "[%s] function called \n", __FUNCTION__);
  LOG(LEVEL_DEBUG, HDR "++ grpcServiceHandler : %p  \n", &grpcServiceHandler );
  LOG(LEVEL_DEBUG, HDR "++ grpcServiceHandler.CommandQueue   : %p  \n", grpcServiceHandler.cmdQueue);
  LOG(LEVEL_DEBUG, HDR "++ grpcServiceHandler.CommandHandler : %p  \n", grpcServiceHandler.cmdHandler );
  LOG(LEVEL_DEBUG, HDR "++ grpcServiceHandler.UpdateCache    : %p  \n", grpcServiceHandler.updCache);
  LOG(LEVEL_DEBUG, HDR "++ grpcServiceHandler.svrConnHandler : %p  \n", grpcServiceHandler.svrConnHandler);

  SRXPROXY_HELLO* hdr  = (SRXPROXY_HELLO*)data;
  uint32_t proxyID     = 0;
  uint8_t  clientID    = 0;
  
  
  /* 
   * To make clientID in accordance with porxyID with proxyMap 
   */
  ClientThread* cthread;
  cthread = (ClientThread*)appendToSList(&grpcServiceHandler.svrConnHandler->svrSock.cthreads, sizeof (ClientThread));

  cthread->active          = true;
  cthread->initialized     = false;
  cthread->goodByeReceived = false;

  cthread->proxyID  = 0; // will be changed for srx-proxy during handshake
  cthread->routerID = 0; // Indicates that it is currently not usable, 
  //cthread->clientFD = cliendFD;
  cthread->svrSock  = &grpcServiceHandler.svrConnHandler->svrSock;
  //cthread->caddr	  = caddr;


  LOG(LEVEL_DEBUG, HDR "[SRx server] Obtained cthread: %p \n", cthread);


  if (ntohs(hdr->version) != SRX_PROTOCOL_VER)
  {
    RAISE_ERROR("Received Hello packet is of protocol version %u but expected "
                "is a Hello packet of protocol version %u",
                ntohs(hdr->version), SRX_PROTOCOL_VER);
    sendError(SRXERR_WRONG_VERSION, NULL, NULL, false);
    //sendGoodbye(item->serverSocket, item->client, false);
  }
  else
  {
    // Figure out the proxyID if it can be used or not. If not, answer with a new proxy ID.
    proxyID = ntohl(hdr->proxyIdentifier);

    clientID = findClientID(grpcServiceHandler.svrConnHandler, proxyID);
    if (clientID == 0)
    {
      clientID = createClientID(grpcServiceHandler.svrConnHandler);
    }


    if (clientID > 0)
    {
      if (!addMapping(grpcServiceHandler.svrConnHandler, proxyID, clientID, cthread, true))
      {
        clientID = 0; // FAIL HANDSHAKE
      }

      LOG(LEVEL_INFO, HDR "[SRx server] proxyID: %d --> mapping[clientID:%d] cthread: %p\n", 
          proxyID,  clientID, grpcServiceHandler.svrConnHandler->proxyMap[clientID].socket);
    }

    LOG (LEVEL_INFO, "Handshake: Connection to proxy[0x%08X] accepted. Proxy "
        "registered as internal client[0x%02X]", proxyID, clientID);

    cthread->proxyID  = proxyID;
    cthread->routerID = clientID;

    grpcServiceHandler.cmdHandler->grpcEnable = true;

    // TODO: client registration should be followed
    //  _processHandshake()
    //  command_handler.c: 233 -


    /* Send Hello Response */
    bool retVal = true;
    uint32_t length = sizeof(SRXPROXY_HELLO_RESPONSE);
    SRXPROXY_HELLO_RESPONSE* pdu = malloc(length);
    memset(pdu, 0, length);

    pdu->type    = PDU_SRXPROXY_HELLO_RESPONSE;
    pdu->version = htons(SRX_PROTOCOL_VER);
    pdu->length  = htonl(length);
    pdu->proxyIdentifier = htonl(proxyID);


    rt->size = length;
    rt->data = (unsigned char*) malloc(length);
    memcpy(rt->data, pdu, length);
    free(pdu);



    // TODO: send goodbye in case there is error 
    //

  }

  return (rt->size == 0) ? false : true;
}

static bool processValidationRequest_grpc(unsigned char *data, RET_DATA *rt, unsigned int grpcClientID)
{
  LOG(LEVEL_DEBUG, HDR "[%s] function called, grpc clientID: %d \n", __FUNCTION__, grpcClientID);
  LOG(LEVEL_INFO, HDR "Enter processValidationRequest", pthread_self());
    
  bool retVal = true;
  SRXRPOXY_BasicHeader_VerifyRequest* hdr =
                                (SRXRPOXY_BasicHeader_VerifyRequest*)data;

  // Determine if a receipt is requested and a result packet must be send
  bool     receipt =    (hdr->flags & SRX_FLAG_REQUEST_RECEIPT)
                      == SRX_FLAG_REQUEST_RECEIPT;
  // prepare already the send flag. Later on, if this is > 0 send a response.
  uint8_t  sendFlags = hdr->flags & SRX_FLAG_REQUEST_RECEIPT;

  bool     doOriginVal = (hdr->flags & SRX_FLAG_ROA) == SRX_FLAG_ROA;
  bool     doPathVal   = (hdr->flags & SRX_FLAG_BGPSEC) == SRX_FLAG_BGPSEC;

  bool      v4     = hdr->type == PDU_SRXPROXY_VERIFY_V4_REQUEST;

  // 1. get an idea what validations are requested:
  //bool originVal = _isSet(hdr->flags, SRX_PROXY_FLAGS_VERIFY_PREFIX_ORIGIN);
  //bool pathVal   = _isSet(hdr->flags, SRX_PROXY_FLAGS_VERIFY_PATH);
  //SRxUpdateID updateID = (SRxUpdateID)item->dataID;


  uint32_t requestToken = receipt ? ntohl(hdr->requestToken)
                                  : DONOTUSE_REQUEST_TOKEN;
  uint32_t originAS = 0;
  SRxUpdateID collisionID = 0;
  SRxUpdateID updateID = 0;

  bool doStoreUpdate = false;
  IPPrefix* prefix = NULL;
  // Specify the client id as a receiver only when validation is requested.
  uint8_t clientID = (uint8_t)grpcClientID; //(doOriginVal || doPathVal) ? client->routerID : 0;

  // 1. Prepare for and generate the ID of the update
  prefix = malloc(sizeof(IPPrefix));
  memset(prefix, 0, sizeof(IPPrefix));
  prefix->length     = hdr->prefixLen;
  BGPSecData bgpData;
  memset (&bgpData, 0, sizeof(BGPSecData));

  uint8_t* valPtr = (uint8_t*)hdr;
  if (v4)
  {
    SRXPROXY_VERIFY_V4_REQUEST* v4Hdr = (SRXPROXY_VERIFY_V4_REQUEST*)hdr;
    valPtr += sizeof(SRXPROXY_VERIFY_V4_REQUEST);
    prefix->ip.version  = 4;
    prefix->ip.addr.v4  = v4Hdr->prefixAddress;
    originAS            = ntohl(v4Hdr->originAS);
    // The next two are in host format for convenience
    bgpData.numberHops  = ntohs(v4Hdr->bgpsecValReqData.numHops);
    bgpData.attr_length = ntohs(v4Hdr->bgpsecValReqData.attrLen);
    // Now in network format as required.
    bgpData.afi         = v4Hdr->bgpsecValReqData.valPrefix.afi;
    bgpData.safi        = v4Hdr->bgpsecValReqData.valPrefix.safi;
    bgpData.local_as    = v4Hdr->bgpsecValReqData.valData.local_as;
  }
  else
  {
    SRXPROXY_VERIFY_V6_REQUEST* v6Hdr = (SRXPROXY_VERIFY_V6_REQUEST*)hdr;
    valPtr += sizeof(SRXPROXY_VERIFY_V6_REQUEST);
    prefix->ip.version  = 6;
    prefix->ip.addr.v6  = v6Hdr->prefixAddress;
    originAS            = ntohl(v6Hdr->originAS);
    // The next two are in host format for convenience
    bgpData.numberHops  = ntohs(v6Hdr->bgpsecValReqData.numHops);
    bgpData.attr_length = ntohs(v6Hdr->bgpsecValReqData.attrLen);
    // Now in network format as required.
    bgpData.afi         = v6Hdr->bgpsecValReqData.valPrefix.afi;
    bgpData.safi        = v6Hdr->bgpsecValReqData.valPrefix.safi;
    bgpData.local_as    = v6Hdr->bgpsecValReqData.valData.local_as;
  }

  // Check if AS path exists and if so then set it
  if (bgpData.numberHops != 0)
  {
    bgpData.asPath = (uint32_t*)valPtr;
  }
  // Check if BGPsec path exits and if so then set it
  if (bgpData.attr_length != 0)
  {
    // BGPsec attribute comes after the as4 path
    bgpData.bgpsec_path_attr = valPtr + (bgpData.numberHops * 4);
  }

  // 2. Generate the CRC based updateID
  updateID = generateIdentifier(originAS, prefix, &bgpData);
  LOG(LEVEL_DEBUG, HDR "\n[SRx server] Generated Update ID: %08X, client ID:%d \n\n", updateID, clientID);


  //  3. Try to find the update, if it does not exist yet, store it.
  SRxResult        srxRes;
  SRxDefaultResult defResInfo;
  // The method getUpdateResult will initialize the result parameters and
  // register the client as listener (only if the update already exists)
  ProxyClientMapping* clientMapping = clientID > 0 ? &grpcServiceHandler.svrConnHandler->proxyMap[clientID]
                                                   : NULL;

  LOG(LEVEL_DEBUG, HDR "[SRx Server] proxyMap[clientID:%d]: %p\n", clientID, clientMapping);

  doStoreUpdate = !getUpdateResult (grpcServiceHandler.svrConnHandler->updateCache, &updateID,
                                    clientID, clientMapping,
                                    &srxRes, &defResInfo);

  if (doStoreUpdate)
  {
    defResInfo.result.roaResult    = hdr->roaDefRes;
    defResInfo.resSourceROA        = hdr->roaResSrc;

    defResInfo.result.bgpsecResult = hdr->bgpsecDefRes;
    defResInfo.resSourceBGPSEC     = hdr->bgpsecResSrc;

    if (!storeUpdate(grpcServiceHandler.svrConnHandler->updateCache, clientID, clientMapping,
                     &updateID, prefix, originAS, &defResInfo, &bgpData))
    {
      RAISE_SYS_ERROR("Could not store update [0x%08X]!!", updateID);
      free(prefix);
      return false;
    }

    // Use the default result.
    srxRes.roaResult    = defResInfo.result.roaResult;
    srxRes.bgpsecResult = defResInfo.result.bgpsecResult;
  }
  free(prefix);
  prefix = NULL;

  LOG(LEVEL_INFO, HDR "+ from updata cache srxRes.roaResult : %02x\n", srxRes.roaResult);
  LOG(LEVEL_INFO, HDR "+ from updata cache srxRes.bgpsecResult : %02x\n", srxRes.bgpsecResult);

  // Just check if the client has the correct values for the requested results
  if (doOriginVal && (hdr->roaDefRes != srxRes.roaResult))
  {
    sendFlags = sendFlags | SRX_FLAG_ROA;
  }
  if (doPathVal && (hdr->bgpsecDefRes != srxRes.bgpsecResult))
  {
    sendFlags = sendFlags | SRX_FLAG_BGPSEC;
  }


  if (sendFlags > 0) // a notification is needed. flags specifies the type
  {
    // TODO: Check specification if we can send a receipt without results, if
    // not the following 6 lines MUST be included, otherwise not.
    if (doOriginVal)
    {
      sendFlags = sendFlags | SRX_FLAG_ROA;
    }
    if (doPathVal)
    {
      sendFlags = sendFlags | SRX_FLAG_BGPSEC;
    }

    // Now send the results we know so far;
    
    /*
       sendVerifyNotification(svrSock, client, updateID, sendFlags,
       requestToken, srxRes.roaResult,
       srxRes.bgpsecResult,
       !self->sysConfig->mode_no_sendqueue);
       */

    uint32_t length = sizeof(SRXPROXY_VERIFY_NOTIFICATION);
    SRXPROXY_VERIFY_NOTIFICATION* pdu = malloc(length);
    memset(pdu, 0, length);

    pdu->type          = PDU_SRXPROXY_VERI_NOTIFICATION;
    pdu->resultType    = sendFlags;
    pdu->requestToken  = htonl(requestToken);
    pdu->roaResult     = srxRes.roaResult;
    pdu->bgpsecResult  = srxRes.bgpsecResult;
    pdu->length        = htonl(length);
    pdu->updateID      = htonl(updateID);

    if ((pdu->requestToken != 0) && (sendFlags < SRX_FLAG_REQUEST_RECEIPT))
    {
      LOG(LEVEL_NOTICE, "Send a notification of update 0x%0aX with request "
          "token 0x%08X but no receipt flag set!", updateID, requestToken);
    }

    pdu->length = htonl(length);

    // return value for response grpc
    rt->size = length;
    rt->data = (unsigned char*) malloc(length);
    memcpy(rt->data, pdu, length);
    free(pdu);


    if ((doOriginVal || doPathVal) && ((sendFlags & SRX_FLAG_ROA_AND_BGPSEC) > 0))
    {
      // Only keep the validation flags.
      hdr->flags = sendFlags & SRX_FLAG_ROA_AND_BGPSEC;

      // create the validation command!
      if (!queueCommand(grpcServiceHandler.cmdQueue, COMMAND_TYPE_SRX_PROXY, NULL, NULL,
            updateID, ntohl(hdr->length), (uint8_t*)hdr))
      {
        RAISE_ERROR("Could not add validation request to command queue!");
        retVal = false;
      }
    }

    LOG(LEVEL_INFO, HDR "Exit processValidationRequest", pthread_self());

  }
  return retVal;
}

static void _processUpdateSigning_grpc(unsigned char *data, RET_DATA *rt, unsigned int grpcClientID)
{
  // TODO Sign the data
  LOG(LEVEL_INFO, "Signing of updates is currently not supported!");
}

//static void _processDeleteUpdate(CommandHandler* cmdHandler, CommandQueueItem* item)
static void _processDeleteUpdate_grpc(unsigned char *data, RET_DATA *rt, unsigned int grpcClientID)
{

  CommandHandler* cmdHandler =  grpcServiceHandler.cmdHandler;

  // TODO: replace item with real pointer variable
  CommandQueueItem* item;
  // For now the delete will NOT remove the update from the cache. It will
  // remove the client - update association though or return an error in case
  // no association existed.
  SRxUpdateID   updateID = (SRxUpdateID)item->dataID;
  ClientThread* clThread = (ClientThread*)item->client;
  SRXPROXY_DELETE_UPDATE* duHdr = (SRXPROXY_DELETE_UPDATE*)item->data;

  if (deleteUpdateFromCache(cmdHandler->updCache, clThread->routerID,
                            &updateID, htons(duHdr->keepWindow)))
  {
    // Reduce the updates by one. BZ308
    cmdHandler->svrConnHandler->proxyMap[clThread->routerID].updateCount--;
  }
  else
  {
    // The update was either not found or the client was not associated to the
    // specified update.
    sendError(SRXERR_UPDATE_NOT_FOUND, item->serverSocket, item->client, false);
    LOG(LEVEL_NOTICE, "Deletion request for update [0x%08X] from client "
                      "[0x%02X] failed, update not found in update cache!");
  }
}

static void _processPeerChange_grpc(unsigned char *data, RET_DATA *rt, unsigned int grpcClientID)
{
  // TODO@ add code for deletion of peer data
  LOG(LEVEL_WARNING, "Peer Changes are not supported prior Version 0.4.0!");
}



//int responseGRPC (int size, unsigned char* data)
RET_DATA responseGRPC (int size, unsigned char* data, unsigned int grpcClientID)
{
    LOG(LEVEL_DEBUG, HDR "[SRx server] [%s] calling - size: %d, grpcClient ID: %02x  \n", __FUNCTION__, size, grpcClientID);
    //setLogLevel(LEVEL_DEBUG);

    /*
    bool ret = _isSet(0x03, 0x01);
    printf("ret bool: %d \n", ret);
    */

    LogLevel lv = getLogLevel();
    LOG(LEVEL_INFO, HDR "srx server log Level: %d\n", lv);

    if (lv >= LEVEL_INFO) {
      printHex(size, data);
    }

    RET_DATA rt;
    memset(&rt, 0x0, sizeof(RET_DATA));

    /* (examples)
     *
    rt.data = (unsigned char*) malloc(size);
    rt.size = size;

    memset(rt.data, 0x00, size);
    rt.data[0] = 0xAF; rt.data[1] = 0x11; rt.data[2] = 0x12; 
    rt.data[3] = 0x33; rt.data[4] = 0xAB; rt.data[5] = 0xCD; rt.data[6] = 0xEF;
    */

    SRXPROXY_BasicHeader* bhdr = (SRXPROXY_BasicHeader*)data;
    uint8_t clientID;
    ClientThread* cthread;
          
    uint32_t length = sizeof(SRXPROXY_GOODBYE);     
    uint8_t pdu[length];                            
    SRXPROXY_GOODBYE* hdr = (SRXPROXY_GOODBYE*)pdu; 

    switch (bhdr->type)
    {                   
      case PDU_SRXPROXY_HELLO:
        processHandshake_grpc(data, &rt);
        break;

      case PDU_SRXPROXY_VERIFY_V4_REQUEST:
      case PDU_SRXPROXY_VERIFY_V6_REQUEST:
        processValidationRequest_grpc(data, &rt, grpcClientID);
        break;

      case PDU_SRXPROXY_SIGN_REQUEST:
        _processUpdateSigning_grpc(data, &rt, grpcClientID);
        break;
      case PDU_SRXPROXY_GOODBYE:
        LOG(LEVEL_DEBUG, HDR "[SRx Server] Received GOOD BYE from proxyID: %d\n", grpcClientID);
        clientID = findClientID(grpcServiceHandler.svrConnHandler, grpcClientID);
      
        LOG(LEVEL_DEBUG, HDR "[SRx server] proxyID: %d --> mapping[clientID:%d] cthread: %p\n", 
          grpcClientID,  clientID, grpcServiceHandler.svrConnHandler->proxyMap[clientID].socket);

        cthread = (ClientThread*)grpcServiceHandler.svrConnHandler->proxyMap[clientID].socket;
        // in order to skip over terminating a client pthread which was not generated if grpc enabled
        cthread->active  = false;
        closeClientConnection(&grpcServiceHandler.cmdHandler->svrConnHandler->svrSock, cthread);

        //clientID = ((ClientThread*)item->client)->routerID;
        deactivateConnectionMapping(grpcServiceHandler.svrConnHandler, clientID, false, 0);
        deleteFromSList(&grpcServiceHandler.cmdHandler->svrConnHandler->clients, cthread);
        grpcServiceHandler.cmdHandler->grpcEnable = false;
        LOG(LEVEL_DEBUG, HDR "GoodBye!", pthread_self());
        break;

      case PDU_SRXPROXY_DELTE_UPDATE:
        //_processDeleteUpdate(cmdHandler, item);
        _processDeleteUpdate_grpc(data, &rt, grpcClientID);
        break;
      case PDU_SRXPROXY_PEER_CHANGE:
        //_processPeerChange(cmdHandler, item);
        _processPeerChange_grpc(data, &rt, grpcClientID);
        break;
      default:
        RAISE_ERROR("Unknown/unsupported pdu type: %d", bhdr->type);

        memset(pdu, 0, length);                         
        LOG(LEVEL_INFO, HDR" send Goodbye! called" );  
        hdr->type       = PDU_SRXPROXY_GOODBYE;         
        hdr->keepWindow = htons(900);            
        hdr->length     = htonl(length);                

        LOG(LEVEL_DEBUG, HDR "\n\nCalling CallBack function forGoodbye STREAM\n\n");         
        cb_proxyGoodBye(*hdr);
        
        // XXX: NOTE: do the same way in GoodBye above
        clientID = findClientID(grpcServiceHandler.svrConnHandler, grpcClientID);
        LOG(LEVEL_INFO, HDR "[SRx server] proxyID: %d --> mapping[clientID:%d] cthread: %p\n", 
          grpcClientID,  clientID, grpcServiceHandler.svrConnHandler->proxyMap[clientID].socket);
        cthread = (ClientThread*)grpcServiceHandler.svrConnHandler->proxyMap[clientID].socket;
        cthread->active  = false;
        closeClientConnection(&grpcServiceHandler.cmdHandler->svrConnHandler->svrSock, cthread);
        deactivateConnectionMapping(grpcServiceHandler.svrConnHandler, clientID, false, 0);
        deleteFromSList(&grpcServiceHandler.cmdHandler->svrConnHandler->clients, cthread);
        grpcServiceHandler.cmdHandler->grpcEnable = false;
        LOG(LEVEL_DEBUG, HDR "GoodBye!", pthread_self());
    }

    return rt;
}













