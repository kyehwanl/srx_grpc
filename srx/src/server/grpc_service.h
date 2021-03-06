#ifndef GRPC_SERVICE_H
#define	GRPC_SERVICE_H

#include "server/command_queue.h"
#include "server/command_handler.h"
#include "server/update_cache.h"
#include "shared/srx_packets.h"
#include "util/log.h"
#include "server/server_connection_handler.h"

typedef struct {
  // Arguments (create)
  CommandQueue*             cmdQueue;
  CommandHandler*           cmdHandler;
  ServerConnectionHandler*  svrConnHandler;
  //BGPSecHandler*            bgpsecHandler;
  //RPKIHandler*              rpkiHandler;
  UpdateCache*              updCache;

  // Argument (start)
  //CommandQueue*             queue;

} GRPC_ServiceHandler;

GRPC_ServiceHandler     grpcServiceHandler;


typedef struct {
    unsigned int size;
    unsigned char *data;
} RET_DATA;

//int responseGRPC (int size);
//int responseGRPC (int size, unsigned char* data);
RET_DATA responseGRPC (int size, unsigned char* data, unsigned int grpcClientID);



#endif /* GRPC_SERVICE_H */




