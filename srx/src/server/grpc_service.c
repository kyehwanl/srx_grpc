

#include <stdio.h>
#include "server/grpc_service.h"
#include "server/command_handler.h"
#include "shared/srx_packets.h"
#include "util/log.h"

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
  printf("[%s] function called \n", __FUNCTION__);
  printf("++ grpcServiceHandler : %p  \n", &grpcServiceHandler );
  printf("++ grpcServiceHandler.CommandHandler : %p  \n", grpcServiceHandler.cmdHandler );
  printf("++ grpcServiceHandler.UpdateCache    : %p  \n", grpcServiceHandler.updCache);
  printf("++ grpcServiceHandler.svrConnHandler : %p  \n", grpcServiceHandler.svrConnHandler);

  SRXPROXY_HELLO* hdr  = (SRXPROXY_HELLO*)data;
  uint32_t proxyID     = 0;
  uint8_t  clientID    = 0;

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

  }

  return (rt->size == 0) ? false : true;
}

static bool processValidationRequest_grpc(unsigned char *data, RET_DATA *rt)
{
  printf("[%s] function called \n", __FUNCTION__);
  LOG(LEVEL_DEBUG, HDR "Enter processValidationRequest", pthread_self());
    
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
  uint8_t clientID = 100; //(doOriginVal || doPathVal) ? client->routerID : 0;

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


  //  3. Try to find the update, if it does not exist yet, store it.
  SRxResult        srxRes;
  SRxDefaultResult defResInfo;
  // The method getUpdateResult will initialize the result parameters and
  // register the client as listener (only if the update already exists)
  ProxyClientMapping* clientMapping = clientID > 0 ? &grpcServiceHandler.svrConnHandler->proxyMap[clientID]
                                                   : NULL;
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

  printf("+ from updata cache srxRes.roaResult : %02x\n", srxRes.roaResult);
  printf("+ from updata cache srxRes.bgpsecResult : %02x\n", srxRes.bgpsecResult);

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

  }
  return retVal;
}



//int responseGRPC (int size, unsigned char* data)
RET_DATA responseGRPC (int size, unsigned char* data)
{
    printf("[%s] calling - size: %d \n", __FUNCTION__, size);

    bool ret = _isSet(0x03, 0x01);
    printf("ret bool: %d \n", ret);

    printHex(size, data);

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

    // TODO: need handling functions for multiplexing the request
    SRXPROXY_BasicHeader* bhdr = (SRXPROXY_BasicHeader*)data;
          
    switch (bhdr->type)

    {                   
      case PDU_SRXPROXY_HELLO:
        // The mapping information will be maintained during the handshake
        processHandshake_grpc(data, &rt);
        break;

      case PDU_SRXPROXY_VERIFY_V4_REQUEST:
      case PDU_SRXPROXY_VERIFY_V6_REQUEST:
        processValidationRequest_grpc(data, &rt);
        //processUpdateValidation_grpc(data, &rt);
        break;

      case PDU_SRXPROXY_SIGN_REQUEST:
        //_processUpdateSigning(cmdHandler, item);
        break;
      case PDU_SRXPROXY_GOODBYE:
        /*
        gbhdr = (SRXPROXY_GOODBYE*)item->data;
        closeClientConnection(&cmdHandler->svrConnHandler->svrSock,
            item->client);
        clientID = ((ClientThread*)item->client)->routerID;
        //cmdHandler->svrConnHandler->proxyMap[clientID].isActive = false;
        // The deaktivation will also delete because it did not crash
        deactivateConnectionMapping(cmdHandler->svrConnHandler, clientID,
            false, htons(gbhdr->keepWindow));
        //delMapping(cmdHandler->svrConnHandler, clientID);

        deleteFromSList(&cmdHandler->svrConnHandler->clients,
            item->client);
        LOG(LEVEL_DEBUG, HDR "GoodBye!", pthread_self());
        */
        break;
      case PDU_SRXPROXY_DELTE_UPDATE:
        //_processDeleteUpdate(cmdHandler, item);
        break;
      case PDU_SRXPROXY_PEER_CHANGE:
        //_processPeerChange(cmdHandler, item);
        break;
      default:
        RAISE_ERROR("Unknown/unsupported pdu type: %d", bhdr->type);
        /*
        sendError(SRXERR_INVALID_PACKET, item->serverSocket,
            item->client, false);
        sendGoodbye(item->serverSocket, item->client, false);
        closeClientConnection(&cmdHandler->svrConnHandler->svrSock,
            item->client);

        clientID = ((ClientThread*)item->client)->routerID;
        // The deaktivatio will also delete the mapping because it was NOT
        // a crash.
        deactivateConnectionMapping(cmdHandler->svrConnHandler, clientID,
            false, cmdHandler->sysConfig->defaultKeepWindow);
        //cmdHandler->svrConnHandler->proxyMap[clientID].isActive = false;
        //delMapping(cmdHandler->svrConnHandler, clientID);

        deleteFromSList(&cmdHandler->svrConnHandler->clients,
            item->client);
            */
    }

    return rt;
}













