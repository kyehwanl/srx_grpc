package main

/*

#cgo CFLAGS: -I/opt/project/gobgp_test/gowork/src/srx_grpc/srx/test_install/include/ -I/opt/project/gobgp_test/gowork/src/srx_grpc/srx/src/ -I/opt/project/gobgp_test/gowork/src/srx_grpc/srx/src/../extras/local/include -I/opt/project/srx_test1/srx/../_inst//include

#cgo LDFLAGS: -L/opt/project/gobgp_test/gowork/src/srx_grpc/srx/src/.libs -lgrpc_service -Wl,-rpath -Wl,/opt/project/gobgp_test/gowork/src/srx_grpc/srx/src/.libs -Wl,--unresolved-symbols=ignore-all

//#cgo LDFLAGS: -L/opt/project/gobgp_test/gowork/src/srx_grpc/srx/src/server -lgrpc_service -Wl,-rpath -Wl,/opt/project/gobgp_test/gowork/src/srx_grpc/srx/src/server -L/opt/project/gobgp_test/gowork/src/srx_grpc/srx/test_install/lib64/srx -lSRxProxy -Wl,-rpath -Wl,/opt/project/gobgp_test/gowork/src/srx_grpc/srx/test_install/lib64/srx -Wl,--unresolved-symbols=ignore-all
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
	"encoding/binary"
	_ "io"
	"runtime"
	_ "time"
	"unsafe"
)

var port = flag.Int("port", 50000, "The server port")
var gStream pb.SRxApi_SendAndWaitProcessServer
var gStream_verify pb.SRxApi_ProxyVerifyServer

var done chan bool

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

	if f == 0 && b == nil {
		_, _, line, _ := runtime.Caller(0)
		log.Printf("[server:%d] close stream ", line)
		//done <- true
		//return
	}

	//b := []byte{0x10, 0x11, 0x40, 0x42, 0xAB, 0xCD, 0xEF}
	resp := pb.PduResponse{
		Data:             b,
		Length:           uint32(len(b)),
		ValidationStatus: 2,
	}

	if gStream != nil {
		if resp.Data == nil && resp.Length == 0 {
			_, _, line, _ := runtime.Caller(0)
			log.Printf("[server:%d] close stream ", line)
			//close(done)
		} else {
			if err := gStream.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
			_, _, line, _ := runtime.Caller(0)
			log.Printf("[%d] sending stream data", line+1)
		}

		/*
			TODO: How to Send EOF from the server side through grpc ??

				if f == -1 {
					_, err := gStream.Recv()
					if err == io.EOF {
						log.Printf("end stream connection")
						log.Println(err)
						close(done)
					}
				}
		*/
	}

}

func (s *Server) SendPacketToSRxServer(ctx context.Context, pdu *pb.PduRequest) (*pb.PduResponse, error) {
	data := uint32(0x07)
	//C.setLogMode(3)
	fmt.Printf("server: %s %#v\n", pdu.Data, pdu)
	//C.setLogMode(7)
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
	ctx := stream.Context()
	done = make(chan bool)
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}

		_, _, line, _ := runtime.Caller(0)
		fmt.Printf("+ [%d] server context done\n", line+1)

		// FIXME : channel panic: close of closed channel
		///*
		_, ok := <-done
		if ok == true {
			fmt.Printf("+ server close the channel done here\n")
			close(done)
		}
		//*/
		/*
			fmt.Printf("+ done: %#v\n", done)
			if done != nil {
				_, _, line, _ := runtime.Caller(0)
				fmt.Printf("+ [%d] server close the channel done here\n", line+1)
				close(done)
			}
		*/
	}()

	data := uint32(0x09)
	//C.setLogMode(3)
	fmt.Printf("stream server: %s %#v\n", pdu.Data, pdu)
	//C.setLogMode(7)
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

	//time.Sleep(5 * time.Second)

	<-done
	log.Printf("Finished with RPC send [Send_Wait_Process] \n")

	return nil
}

func (s *Server) ProxyHello(ctx context.Context, pdu *pb.ProxyHelloRequest) (*pb.ProxyHelloResponse, error) {
	//data := uint32(0x07)
	//C.setLogLevel(0x07)
	fmt.Printf("server: %#v\n", pdu)
	fmt.Println("calling SRxServer server:ProxyHello()")

	fmt.Printf("input :  %#v\n", pdu.Type)
	fmt.Println("ProxyHelloRequest", pdu)
	fmt.Printf("ProxyHelloRequest: %#v", pdu)

	/* serialize */
	buf := make([]byte, C.sizeof_SRXPROXY_HELLO)
	buf[0] = byte(pdu.Type)
	binary.BigEndian.PutUint16(buf[1:3], uint16(pdu.Version))
	binary.BigEndian.PutUint32(buf[4:8], pdu.Length)
	binary.BigEndian.PutUint32(buf[8:12], pdu.ProxyIdentifier)
	binary.BigEndian.PutUint32(buf[12:16], pdu.Asn)
	binary.BigEndian.PutUint32(buf[16:20], pdu.NoPeerAS)

	retData := C.RET_DATA{}
	retData = C.responseGRPC(C.int(20), (*C.uchar)(unsafe.Pointer(&buf[0])))

	b := C.GoBytes(unsafe.Pointer(retData.data), C.int(retData.size))
	fmt.Printf("return size: %d \t data: %#v\n", retData.size, b)

	return &pb.ProxyHelloResponse{
		Type:            uint32(b[0]),
		Version:         uint32(binary.BigEndian.Uint16(b[1:3])),
		Zero:            uint32(b[3]),
		Length:          binary.BigEndian.Uint32(b[4:8]),
		ProxyIdentifier: binary.BigEndian.Uint32(b[8:12]),
	}, nil
}

func (s *Server) ProxyVerify(pdu *pb.ProxyVerifyV4Request, stream pb.SRxApi_ProxyVerifyServer) error {
	fmt.Println("calling SRxServer server:ProxyVerify()")

	gStream_verify = stream
	ctx := stream.Context()
	done := make(chan bool)
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()

	fmt.Printf("stream server: %#v\n", pdu)

	retData := C.RET_DATA{}
	retData = C.responseGRPC(C.int(0), (*C.uchar)(unsafe.Pointer(nil)))

	b := C.GoBytes(unsafe.Pointer(retData.data), C.int(retData.size))
	fmt.Printf("return size: %d \t data: %#v\n", retData.size, b)

	resp := pb.ProxyVerifyNotify{
		Type:       2,
		ResultType: 3,
		RoaResult:  3,
	}

	if err := stream.Send(&resp); err != nil {
		log.Printf("send error %v", err)
	}
	log.Printf("sending stream data")

	//time.Sleep(5 * time.Second)

	<-done
	log.Printf("Finished with RPC send \n")

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
		log.Printf("failed to listen: %v", err)
	}

	server := NewServer(grpc.NewServer())
	if err := server.grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}

func main() {
	//fmt.Println("size: ", C.sizeof_SRXPROXY_HELLO)
	Serve()
}
