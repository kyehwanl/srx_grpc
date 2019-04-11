package main

/*
#cgo CFLAGS: -I/opt/project/gobgp_test/gowork/src/srx_grpc/srx/test_install/include/ -I/opt/project/gobgp_test/gowork/src/srx_grpc/srx/src/ -I/opt/project/gobgp_test/gowork/src/srx_grpc/srx/src/../extras/local/include -I/opt/project/srx_test1/srx/../_inst//include

//#include <stdlib.h>
#include "shared/srx_packets.h"
*/
import "C"

import (
	"encoding/binary"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"runtime"
	pb "srx_grpc"
	"unsafe"
)

const (
	address     = "localhost:50000"
	defaultName = "RPKI_DATA"
)

type Client struct {
	conn *grpc.ClientConn
	cli  pb.SRxApiClient
}

type ProxyVerifyClient struct {
	stream pb.SRxApi_ProxyVerifyClient
}

var client Client

//export Init
func Init(address string) uint32 {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	cli := pb.NewSRxApiClient(conn)

	client.conn = conn
	client.cli = cli

	return 0
}

//export Run
func Run(data []byte) uint32 {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSRxApiClient(conn)
	client.conn = conn
	client.cli = c

	// Contact the server and print out its response.
	//ctx, _ := context.WithTimeout(context.Background(), time.Second)
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()

	if data == nil {
		fmt.Println("#############")
		data = []byte(defaultName)
	}

	fmt.Printf("input data: %#v\n", data)

	r, err := c.SendPacketToSRxServer(context.Background(), &pb.PduRequest{Data: data, Length: uint32(len(data))})
	if err != nil {
		log.Fatalf("could not receive: %v", err)
	}

	fmt.Printf("data : %#v\n", r.Data)
	fmt.Printf("size : %#v\n", r.Length)
	fmt.Printf("status: %#v\n", r.ValidationStatus)

	//return r.ValidationStatus, err
	return uint32(r.ValidationStatus)
}

/* test 2: request HelloRequest */
/*
   hdr->type            = PDU_SRXPROXY_HELLO;
   hdr->version         = htons(SRX_PROTOCOL_VER);
   hdr->length          = htonl(length);
   hdr->proxyIdentifier = htonl(5);   // htonl(proxy->proxyID);
   hdr->asn             = htonl(65005);
   hdr->noPeers         = htonl(noPeers);
*/

//export RunProxyHello
func RunProxyHello(data []byte) *C.uchar {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	cli := pb.NewSRxApiClient(conn)

	fmt.Printf("input data: %#v\n", data)
	/*
		client.conn = conn
		client.cli = cli
	*/

	/*
		retData := C.RET_DATA{}
		hData := C.SRXPROXY_HELLO{}
		fmt.Println("temp:", retData, hData)
	*/

	req := pb.ProxyHelloRequest{
		Type:            uint32(data[0]),
		Version:         uint32(binary.BigEndian.Uint16(data[1:3])),
		Zero:            uint32(data[3]),
		Length:          binary.BigEndian.Uint32(data[4:8]),
		ProxyIdentifier: binary.BigEndian.Uint32(data[8:12]),
		Asn:             binary.BigEndian.Uint32(data[12:16]),
		NoPeerAS:        binary.BigEndian.Uint32(data[16:20]),
	}

	resp, err := cli.ProxyHello(context.Background(), &req)
	if err != nil {
		log.Fatalf("could not receive: %v", err)
	}

	fmt.Printf("+ HelloRequest	: %#v\n", req)
	fmt.Printf("+ response		: %#v\n", resp)

	rp := C.SRXPROXY_HELLO_RESPONSE{
		//version:         C.ushort(resp.Version), // --> TODO: need to pack/unpack for packed struct in C
		length:          C.uint(resp.Length),
		proxyIdentifier: C.uint(resp.ProxyIdentifier),
	}
	rp._type = C.uchar(resp.Type)

	//fmt.Println("rp:", rp)

	buf := make([]byte, C.sizeof_SRXPROXY_HELLO_RESPONSE)
	//buf := make([]byte, 12)
	buf[0] = byte(resp.Type)
	binary.BigEndian.PutUint16(buf[1:3], uint16(resp.Version))
	binary.BigEndian.PutUint32(buf[4:8], resp.Length)
	binary.BigEndian.PutUint32(buf[8:12], resp.ProxyIdentifier)

	cb := (*[C.sizeof_SRXPROXY_HELLO_RESPONSE]C.uchar)(C.malloc(C.sizeof_SRXPROXY_HELLO_RESPONSE))
	// TODO: defer C.free(unsafe.Pointer(cb)) at caller side --> DONE
	cstr := (*[C.sizeof_SRXPROXY_HELLO_RESPONSE]C.uchar)(unsafe.Pointer(&buf[0]))

	for i := 0; i < C.sizeof_SRXPROXY_HELLO_RESPONSE; i++ {
		cb[i] = cstr[i]
	}

	//return (*C.uchar)(unsafe.Pointer(&buf[0]))
	return &cb[0]
}

//export RunStream
func RunStream(data []byte) uint32 {

	// TODO: how to persistently obtain grpc Dial object

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSRxApiClient(conn)

	if data == nil {
		fmt.Println("#############")
		data = []byte(defaultName)
	}

	fmt.Printf("input data for stream response: %#v\n", data)

	stream, err := c.SendAndWaitProcess(context.Background(), &pb.PduRequest{Data: data, Length: uint32(len(data))})
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	ctx := stream.Context()
	done := make(chan bool)
	var r pb.PduResponse

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				log.Fatalf("[client] EOF close ")
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}

			// TODO: receive process here
			fmt.Printf("data : %#v\n", resp.Data)
			fmt.Printf("size : %#v\n", resp.Length)
			fmt.Printf("status: %#v\n", resp.ValidationStatus)
			r = *resp

			if resp.Data == nil && resp.Length == 0 {
				_, _, line, _ := runtime.Caller(0)
				log.Fatalf("[client:%d] close stream ", line+1)
				//stream.CloseSend()
				close(done)
			}
		}

	}()

	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		fmt.Printf("+ client context done\n")
		close(done)
	}()

	<-done
	log.Printf("Finished with Resopnse valie: %d", uint32(r.ValidationStatus))
	log.Fatalf("Finished with Resopnse valie:")
	fmt.Printf("Finished with Resopnse valie: %d", uint32(r.ValidationStatus))
	return uint32(r.ValidationStatus)
}

func RunProxyVerify(data []byte) uint32 {

	// TODO: how to persistently obtain grpc Dial object

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSRxApiClient(conn)

	if data == nil {
		fmt.Println("#############")
		data = []byte(defaultName)
	}

	fmt.Printf("input data for stream response: %#v\n", data)

	stream, err := c.ProxyVerify(context.Background(), &pb.ProxyVerifyV4Request{})
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	ctx := stream.Context()
	done := make(chan bool)
	var r pb.ProxyVerifyNotify

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}

			// TODO: receive process here
			fmt.Printf("received data: %#v\n", resp)
			r = *resp

			/*
				if resp.Data == nil && resp.Length == 0 {
					log.Fatalf("close stream ")
					close(done)
				}
			*/
		}

	}()

	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()

	<-done
	return 0
}

func main() {
	/* FIXME XXX
	buff_hello_request :=
		[]byte{0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x14, 0x0, 0x0, 0x0, 0x5, 0x0, 0x0, 0xfd, 0xed, 0x0, 0x0, 0x0, 0x0}
	res := RunProxyHello(buff_hello_request)
	//r := Run(buff_hello_request)
	log.Printf("Transferred: %#v\n\n", res)
	*/

	/*
		buff_verify_req := []byte{
			0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x14, 0x0, 0x0, 0x0, 0x5, 0x0, 0x0, 0xfd, 0xed, 0x0, 0x0, 0x0, 0x0}
		RunProxyVerify(buff_verify_req)
	*/

	data := []byte(defaultName)
	data2 := []byte{0x10, 0x11, 0x40, 0x42}
	data3 := []byte{0x10, 0x11, 0x40, 0x42, 0xAB, 0xCD, 0xEF}

	r := Run(data)
	log.Printf("Transferred: %#v\n\n", r)

	r = Run(data2)
	log.Printf("Transferred: %#v\n\n", r)

	r = RunStream(data3)
	log.Printf("Transferred: %#v\n\n", r)

}

/* NOTE

TODO 1: init function - for receiving client (*grpc.ClientConn)
		--> maybe good to use a global variable for client

TODO 2

*/
