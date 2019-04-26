
#include <stdio.h>
#include "client/grpc_client_service.h"
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


