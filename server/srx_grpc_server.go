package main

/*

#cgo CFLAGS: -I/opt/project/gobgp_test/gowork/src/srx_grpc/srx/test_install/include/ -I/opt/project/gobgp_test/gowork/src/srx_grpc/srx/src/ -I/opt/project/gobgp_test/gowork/src/srx_grpc/srx/src/../extras/local/include -I/opt/project/srx_test1/srx/../_inst//include
#cgo LDFLAGS: -L/opt/project/gobgp_test/gowork/src/srx_grpc/srx/src/server -lgrpc_service -Wl,-rpath -Wl,/opt/project/gobgp_test/gowork/src/srx_grpc/srx/src/server -L/opt/project/gobgp_test/gowork/src/srx_grpc/srx/test_install/lib64/srx -lSRxProxy -Wl,-rpath -Wl,/opt/project/gobgp_test/gowork/src/srx_grpc/srx/test_install/lib64/srx -Wl,--unresolved-symbols=ignore-all

#include <stdio.h>
#include "srx/srx_api.h"
#include "server/grpc_service.h"



extern void cb_proxy(int f, void* user_data);
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
	_ "bytes"
	"time"
	"unsafe"
)

var port = flag.Int("port", 50000, "The server port")
var gStream pb.SRxApi_SendAndWaitProcessServer

type Server struct {
	grpcServer *grpc.Server
}

//export cb_proxy
func cb_proxy(f C.int, v unsafe.Pointer) {
	fmt.Printf("proxy callback function : arg[%d, %#v]\n", f, v)

	b := C.GoBytes(unsafe.Pointer(v), f)
	// call my callback
	MyCallback(int(f), b)
}

func MyCallback(f int, b []byte) {

	fmt.Printf("My callback function - received arg: %d, %#v \n", f, b)

	//b := []byte{0x10, 0x11, 0x40, 0x42, 0xAB, 0xCD, 0xEF}
	resp := pb.PduResponse{
		Data:             b,
		Length:           uint32(len(b)),
		ValidationStatus: 2,
	}

	if gStream != nil {
		if err := gStream.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
		log.Printf("sending stream data")
	}

}

func (s *Server) SendPacketToSRxServer(ctx context.Context, pdu *pb.PduRequest) (*pb.PduResponse, error) {
	data := uint32(0x07)
	C.setLogMode(3)
	fmt.Printf("server: %s %#v\n", pdu.Data, pdu)
	C.setLogMode(7)
	fmt.Println("calling SRxServer responseGRPC()")

	retData := C.RET_DATA{}
	retData = C.responseGRPC(C.int(pdu.Length), (*C.uchar)(unsafe.Pointer(&pdu.Data[0])))

	b := C.GoBytes(unsafe.Pointer(retData.data), C.int(retData.size))
	fmt.Printf("return size: %d \t data: %#v\n", retData.size, b)

	//C.setLogMode(3)
	return &pb.PduResponse{
		Data:             b,
		Length:           uint32(retData.size),
		ValidationStatus: data}, nil
}

func (s *Server) SendAndWaitProcess(pdu *pb.PduRequest, stream pb.SRxApi_SendAndWaitProcessServer) error {

	gStream = stream

	data := uint32(0x09)
	C.setLogMode(3)
	fmt.Printf("stream server: %s %#v\n", pdu.Data, pdu)
	C.setLogMode(7)
	fmt.Println("calling SRxServer responseGRPC()")

	retData := C.RET_DATA{}
	retData = C.responseGRPC(C.int(pdu.Length), (*C.uchar)(unsafe.Pointer(&pdu.Data[0])))

	b := C.GoBytes(unsafe.Pointer(retData.data), C.int(retData.size))
	fmt.Printf("return size: %d \t data: %#v\n", retData.size, b)

	resp := pb.PduResponse{
		Data:             b,
		Length:           uint32(retData.size),
		ValidationStatus: data,
	}

	if err := stream.Send(&resp); err != nil {
		log.Printf("send error %v", err)
	}
	log.Printf("sending stream data")

	time.Sleep(5 * time.Second)

	return nil
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
	Serve()
}
