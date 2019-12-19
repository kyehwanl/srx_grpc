
#include <stdio.h>
#include "client/grpc_client_service.h"
#include "client/client_connection_handler.h"
#define HDR  "([0x%08X] GRPC_Client_ServiceHandler): "


static void dispatchPackets_grpc(SRXPROXY_BasicHeader* packet, void* proxyPtr);

extern SRxProxy* g_proxy;


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



void processVerifyNotify_grpc(SRXPROXY_VERIFY_NOTIFICATION* hdr)
{
    printf("+++ [%s] called in proxy: %p \n", __FUNCTION__, g_proxy);
    SRxProxy* proxy = g_proxy;

    printHex(sizeof(SRXPROXY_VERIFY_NOTIFICATION), (unsigned char*)hdr);

    if (proxy)
    {
        if (proxy->resCallback != NULL)
        {
            bool hasReceipt = (hdr->resultType & SRX_FLAG_REQUEST_RECEIPT)
                == SRX_FLAG_REQUEST_RECEIPT;
            bool useROA     = (hdr->resultType & SRX_FLAG_ROA) == SRX_FLAG_ROA;
            bool useBGPSEC  = (hdr->resultType & SRX_FLAG_BGPSEC) == SRX_FLAG_BGPSEC;

            uint32_t localID = ntohl(hdr->requestToken);
            SRxUpdateID updateID = ntohl(hdr->updateID);

#ifdef BZ263
            ct++;
            printf("#%u - uid:0x%08x lid:0x%08X (%u)\n", ct, updateID, localID,
                    localID);
#endif

            if (localID > 0 && !hasReceipt)
            {
                printf(" -> ERROR, no receipt flag set.\n");
                LOG(LEVEL_WARNING, HDR "Unusual notification for update [0x%08X] with "
                        "local id [0x%08X] but receipt flag NOT SET!",
                        updateID, localID);
                localID = 0;
            }
            else
            {
                LOG(LEVEL_DEBUG, HDR "Update [0x%08X] with localID [0x%08X]: %d",
                        updateID, localID, localID);
            }

            uint8_t roaResult    = useROA ? hdr->roaResult : SRx_RESULT_UNDEFINED;
            uint8_t bgpsecResult = useBGPSEC ? hdr->bgpsecResult : SRx_RESULT_UNDEFINED;
            ValidationResultType valType = hdr->resultType & SRX_FLAG_ROA_AND_BGPSEC;

            // hasReceipt ? localID : 0 is result of BZ263
            proxy->resCallback(updateID, localID, valType, roaResult, bgpsecResult,
                    proxy->userPtr);
        }
        else
        {
            LOG(LEVEL_INFO, "processVerifyNotify: NO IMPLEMENTATION PROVIDED FOR "
                    "proxy->resCallback!!!\n");
        }
    }
    else
    {
        printf("this client doens't have a proxy pointer set, maybe due to simple test\n");
    }
}

void processGoodbye_grpc(SRXPROXY_GOODBYE* hdr)
{
    printf("+++ [%s] called in proxy: %p \n", __FUNCTION__, g_proxy);
    SRxProxy* proxy = g_proxy;

    if (proxy)
    {
        // The client connection handler
        ClientConnectionHandler* connHandler =
            (ClientConnectionHandler*)proxy->connHandler;
        LOG(LEVEL_DEBUG, HDR "Received Goodbye", pthread_self());
        // SERVER CLOSES THE CONNECTION. END EVERYTHING.
        connHandler->established = false;
        connHandler->stop = true;

        //releaseClientConnectionHandler(connHandler);

        // It is possible to receive a Goodbye during handshake in this case the 
        // connection handler is NOT initialized yet. The main process is still in 
        // init process and the init process has to cleanup.
        if (connHandler->initialized)
        {
            // Do not receive or try to connect anymore
            connHandler->stop = true;

            // Make sure the SIGINT does not terminate the program
            signal(SIGINT, SIG_IGN); // Ignore the signals


            // Reinstall the default signal handler
            signal(SIGINT, SIG_DFL);

            // Deallocate the send queue and lock
            acquireWriteLock(&connHandler->queueLock);
            releaseSList(&connHandler->sendQueue);
            releaseRWLock(&connHandler->queueLock);

            // Deallocate The packet receive monitor
            if (connHandler->rcvMonitor != NULL)
            {
                pthread_mutex_destroy(connHandler->rcvMonitor);
                free(connHandler->rcvMonitor);
                connHandler->rcvMonitor = NULL;
            }

            if (connHandler->cond != NULL)
            {
                pthread_cond_destroy(connHandler->cond);
                free(connHandler->cond);
                connHandler->cond       = NULL;
            }
        }
    }
    else
    {
        printf("this client doens't have a proxy pointer set, maybe due to simple test\n");
    }

}

void processSyncRequest_grpc(SRXPROXY_SYNCH_REQUEST* hdr)
{
    printf("+++ [%s] called in proxy: %p \n", __FUNCTION__, g_proxy);
    SRxProxy* proxy = g_proxy;

    if (proxy)
    {
        if (proxy->syncNotification != NULL)
        {
            proxy->syncNotification(proxy->userPtr);
        }
        else
        {
            LOG(LEVEL_INFO, "processSyncRequest: NO IMPLEMENTATION PROVIDED FOR "
                    "proxy->syncNotification!!!\n");
        }
    }
    else{

        printf("this client doens't have a proxy pointer set, maybe due to simple test\n");
    }
}

void processSignNotify_grpc(SRXPROXY_SIGNATURE_NOTIFICATION* hdr)
{
    printf("+++ [%s] called in proxy: %p \n", __FUNCTION__, g_proxy);
    SRxProxy* proxy = g_proxy;
    
    if (proxy)
    {
        // @TODO: Finishe the implementation with the correct data.
        if (proxy->sigCallback != NULL)
        {
            LOG(LEVEL_INFO, "processSignNotify: NOT IMPLEMENTED IN THIS PROTOTYPE!!!\n");
            SRxUpdateID updId = hdr->updateIdentifier;
            //TODO finish processSigNotify - especially the bgpsec data
            BGPSecCallbackData bgpsecCallback;
            bgpsecCallback.length = 0;
            bgpsecCallback.data = NULL;
            proxy->sigCallback(updId, &bgpsecCallback, proxy->userPtr);
        }
        else
        {
            LOG(LEVEL_INFO, "processSignNotify: NO IMPLEMENTATION PROVIDED FOR "
                    "proxy->sigCallback!!!\n");
        }
    }
}
