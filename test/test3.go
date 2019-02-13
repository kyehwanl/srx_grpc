package main

/*
#cgo CFLAGS: -I/opt/project/gobgp_test/gowork/src/srxapi/srx/srx_install/include/srx
#cgo LDFLAGS: -L/opt/project/gobgp_test/gowork/src/srxapi/srx/srx_install/lib64/srx -lSRxProxy -Wl,-rpath -Wl,/opt/project/gobgp_test/gowork/src/srxapi/srx/srx_install/lib64/srx

#include <stdio.h>
#include "srx_api.h"

void PrintInternalCall(int i) {
	printf(" testing - %d  calling from server driver module\n", i);
}
*/
import "C"

import (
	"fmt"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
}

func NewServer(g *grpc.Server) *Server {
	grpc.EnableTracing = false
	server := &Server{
		grpcServer: g,
	}
	return server
}

/* export Serve */
func Serve() {

}

func main() {
	//C.PrintInternalCall(12)
	fmt.Println("test")

	Serve()
}
