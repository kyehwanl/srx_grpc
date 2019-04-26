#ifndef GRPC_CLIENT_SERVICE_H
#define GRPC_CLIENT_SERVICE_H

#include "client/srx_api.h"
#include "shared/srx_packets.h"
#include "util/log.h"
//#include "client/client_connection_handler.h"


typedef struct {
    unsigned int size;
    unsigned char *data;
} RET_DATA;

SRxProxy* g_proxy;

void processVerifyNotify_grpc(SRXPROXY_VERIFY_NOTIFICATION* hdr);



#endif
