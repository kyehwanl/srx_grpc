package main

/*
#cgo CFLAGS: -I/opt/project/gobgp_test/gowork/src/srx_grpc/srx/test_install/include/
#cgo LDFLAGS: -L/opt/project/gobgp_test/gowork/src/srx_grpc/srx/test_install/lib64/srx -lSRxProxy -Wl,-rpath -Wl,/opt/project/gobgp_test/gowork/src/srx_grpc/srx/test_install/lib64/srx

#include <stdio.h>
#include "srx/srx_api.h"

void PrintInternalCallTest(int i);
*/
import "C"

import (
	"flag"
	"fmt"
	"log"
	"net"
	pb "srx_grpc"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	//	"github.com/golang/protobuf/proto"
)

var port = flag.Int("port", 50000, "The server port")

type Server struct {
	grpcServer *grpc.Server
}

func (s *Server) SendPacketToSRxServer(ctx context.Context, pdu *pb.PduRequest) (*pb.PduResponse, error) {
	data := uint32(0x07)
	C.setLogMode(3)
	fmt.Printf("server: %s %#v\n", pdu.Data, pdu)
	//C.PrintInternalCallTest(12)
	C.setLogMode(7)
	return &pb.PduResponse{ValidationStatus: data}, nil
}

func NewServer(g *grpc.Server) *Server {
	grpc.EnableTracing = false
	server := &Server{
		grpcServer: g,
	}
	pb.RegisterSRxApiServer(g, server)
	return server
}

//export Serve
func Serve() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := NewServer(grpc.NewServer())
	if err := server.grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	//C.PrintInternalCallTest(2)
	Serve()
}
