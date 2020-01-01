package main

/*
#cgo CFLAGS: -I/opt/project/gobgp_test/gowork/src/srx_grpc/srx/test_install/include/ -I/opt/project/gobgp_test/gowork/src/srx_grpc/srx/src/ -I/opt/project/gobgp_test/gowork/src/srx_grpc/srx/src/../extras/local/include -I/opt/project/srx_test1/srx/../_inst//include
#cgo LDFLAGS: /opt/project/gobgp_test/gowork/src/srx_grpc/srx/src/.libs/log.o -L/opt/project/gobgp_test/gowork/src/srx_grpc/srx/src/.libs -lgrpc_client_service -Wl,-rpath -Wl,/opt/project/gobgp_test/gowork/src/srx_grpc/srx/src/.libs -Wl,--unresolved-symbols=ignore-all

#include <stdlib.h>
#include "shared/srx_packets.h"
#include "client/grpc_client_service.h"
*/
import "C"

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	pb "srx_grpc"
	"time"
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

var client Client

type ProxyVerifyClient struct {
	stream pb.SRxApi_ProxyVerifyClient
}

//export InitSRxGrpc
func InitSRxGrpc(addr string) bool {

	/* Disable Logging */
	log.SetFlags(0)               // skip all formatting
	log.SetOutput(ioutil.Discard) // using this as io.Writer to skip logging
	os.Stdout = nil               // to suppress fmt.Print

	log.Printf("InitSRxGrpc Called \n")
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
		return false
	}
	//log.Printf("conn: %#v, err: %#v\n", conn, err)
	log.Printf("gRPC Client Initiated and Connected Server Address: %s\n", addr)
	//defer conn.Close()
	cli := pb.NewSRxApiClient(conn)

	client.conn = conn
	client.cli = cli

	//fmt.Printf("cli : %#v\n", cli)
	//fmt.Printf("client.cli : %#v\n", client.cli)
	//fmt.Println()
	return true
}

//export Run
func Run(data []byte) uint32 {
	// Set up a connection to the server.
	cli := client.cli
	fmt.Printf("client : %#v\n", client)
	fmt.Printf("client.cli data: %#v\n", client.cli)
	fmt.Println()

	// Contact the server and print out its response.
	//ctx, _ := context.WithTimeout(context.Background(), time.Second)
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()

	if data == nil {
		fmt.Println("#############")
		data = []byte(defaultName)
	}

	fmt.Printf("input data: %#v\n", data)

	r, err := cli.SendPacketToSRxServer(context.Background(), &pb.PduRequest{Data: data, Length: uint32(len(data))})
	if err != nil {
		log.Printf("could not receive: %v", err)
	}

	fmt.Printf("data : %#v\n", r.Data)
	fmt.Printf("size : %#v\n", r.Length)
	fmt.Printf("status: %#v\n", r.ValidationStatus)
	fmt.Println()

	//return r.ValidationStatus, err
	return uint32(r.ValidationStatus)
}

//export RunProxyHello
func RunProxyHello(data []byte) (*C.uchar, uint32) {

	cli := client.cli
	fmt.Println()
	fmt.Printf("client : %#v\n", client)
	fmt.Printf("client.cli : %#v\n", client.cli)
	fmt.Println(cli)
	fmt.Printf("input data: %#v\n", data)
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
		log.Printf("could not receive: %v", err)
	}

	fmt.Printf("+ HelloRequest	: %#v\n", req)
	fmt.Printf("+ response		: %#v\n", resp)

	rp := C.SRXPROXY_HELLO_RESPONSE{
		//version:         C.ushort(resp.Version), // --> TODO: need to pack/unpack for packed struct in C
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

	fmt.Printf("+ cb: %#v\n", cb)
	//return (*C.uchar)(unsafe.Pointer(&buf[0]))
	return &cb[0], resp.ProxyIdentifier
}

type Go_ProxySyncRequest struct {
	_type     uint8
	_reserved uint16
	_zero     uint8
	_length   uint32
}

func (g *Go_ProxySyncRequest) Pack(out unsafe.Pointer) {
	buf := &bytes.Buffer{}
	binary.Write(buf, binary.LittleEndian, g)
	l := buf.Len()
	o := (*[1 << 20]C.uchar)(out)

	for i := 0; i < l; i++ {
		b, _ := buf.ReadByte()
		o[i] = C.uchar(b)
	}
}

type Go_ProxyVerifyNotify struct {
	_type         uint8
	_resultType   uint8
	_roaResult    uint8
	_bgpsecResult uint8
	_length       uint32
	_requestToken uint32
	_updateID     uint32
}

func (g *Go_ProxyVerifyNotify) Pack(out unsafe.Pointer) {
	buf := &bytes.Buffer{}
	binary.Write(buf, binary.LittleEndian, g)
	l := buf.Len()
	o := (*[1 << 20]C.uchar)(out)

	for i := 0; i < l; i++ {
		b, _ := buf.ReadByte()
		o[i] = C.uchar(b)
	}
}

type Go_ProxyGoodBye struct {
	_type       uint8
	_keepWindow uint16
	_zero       uint8
	_length     uint32
}

func (g *Go_ProxyGoodBye) Pack(out unsafe.Pointer) {

	buf := &bytes.Buffer{}
	binary.Write(buf, binary.LittleEndian, g)
	l := buf.Len()
	o := (*[1 << 20]C.uchar)(out)

	for i := 0; i < l; i++ {
		b, _ := buf.ReadByte()
		o[i] = C.uchar(b)
	}
}

func (g *Go_ProxyGoodBye) Unpack(i *C.SRXPROXY_GOODBYE) {

	cdata := C.GoBytes(unsafe.Pointer(i), C.sizeof_SRXPROXY_GOODBYE)
	buf := bytes.NewBuffer(cdata)
	binary.Read(buf, binary.BigEndian, &g._type)
	binary.Read(buf, binary.BigEndian, &g._keepWindow)
	binary.Read(buf, binary.BigEndian, &g._zero)
	binary.Read(buf, binary.BigEndian, &g._length)
}

//export RunProxyGoodBye
func RunProxyGoodBye(in C.SRXPROXY_GOODBYE, grpcClientID uint32) bool {
	cli := client.cli

	fmt.Printf("+ [grpc client] Goobye function: input parameter: %#v \n", in)
	fmt.Printf("+ [grpc client] Goobye function: size: %d \n", C.sizeof_SRXPROXY_GOODBYE)

	goGB := Go_ProxyGoodBye{}
	goGB.Unpack(&in)
	//out := (*[C.sizeof_SRXPROXY_GOODBYE]C.uchar)(C.malloc(C.sizeof_SRXPROXY_GOODBYE))
	fmt.Printf("+ [grpc client] Goodbye out bytes: %#v\n", goGB)

	req := pb.ProxyGoodByeRequest{
		Type:         uint32(goGB._type),
		KeepWindow:   uint32(goGB._keepWindow),
		Zero:         uint32(goGB._zero),
		Length:       uint32(goGB._length),
		GrpcClientID: grpcClientID,
	}
	resp, err := cli.ProxyGoodBye(context.Background(), &req)
	if err != nil {
		log.Printf("could not receive: %v", err)
		return false
	}

	log.Printf("+ [grpc client] GoodByeRequest	: %#v\n", req)
	log.Printf("+ [grpc client] response		: %#v\n", resp)

	return resp.Status
}

//export RunProxyGoodByeStream
func RunProxyGoodByeStream(data []byte, grpcClientID uint32) uint32 {

	//fmt.Printf("+ [grpc client] Goobye Stream function Started : input parameter: %#v \n", data)
	cli := client.cli
	stream, err := cli.ProxyGoodByeStream(context.Background(), &pb.PduRequest{Data: data, Length: uint32(len(data))})
	ctx := stream.Context()
	if err != nil {
		log.Printf("open stream error %v", err)
	}

	go func() {
		<-ctx.Done()
		log.Printf("+ Client Context Done")
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		fmt.Printf("+ client context done\n")
		return
	}()

	resp, err := stream.Recv()
	fmt.Printf("+ [grpc client][GoodByeStream] Goodbye Stream function Received... \n")
	if err == io.EOF {
		log.Printf("[client] EOF close ")
		return 1
	}
	if err != nil {
		log.Printf("can not receive %v", err)
	}

	// NOTE : receive process here
	fmt.Printf("+ [grpc client][GoodByeStream] data : %#v\n", resp.Data)
	fmt.Printf("+ [grpc client][GoodByeStream] size : %#v\n", resp.Length)
	fmt.Println()

	go_gb := &Go_ProxyGoodBye{
		_type:       resp.Data[0],
		_keepWindow: *((*uint16)(unsafe.Pointer(&resp.Data[1]))),
		_zero:       resp.Data[3],
		_length:     *((*uint32)(unsafe.Pointer(&resp.Data[4]))),
	}

	gb := (*C.SRXPROXY_GOODBYE)(C.malloc(C.sizeof_SRXPROXY_GOODBYE))
	defer C.free(unsafe.Pointer(gb))
	go_gb.Pack(unsafe.Pointer(gb))

	log.Printf(" gb: %#v\n", gb)

	//void processGoodbye_grpc(SRXPROXY_GOODBYE* hdr)
	C.processGoodbye_grpc(gb)

	return 0
}

//export RunProxyStream
func RunProxyStream(data []byte, grpcClientID uint32) uint32 {

	//fmt.Printf("+ [grpc client] Stream function Started : input parameter: %#v \n", data)
	cli := client.cli
	stream, err := cli.ProxyStream(context.Background(), &pb.PduRequest{Data: data, Length: uint32(len(data))})
	ctx := stream.Context()
	if err != nil {
		log.Printf("open stream error %v", err)
	}

	go func() {
		<-ctx.Done()
		log.Printf("+ Client Context Done")
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		fmt.Printf("+ client context done\n")
		return
	}()

	for {
		resp, err := stream.Recv()
		fmt.Printf("+ [grpc client][ProxyStream] Proxy_Stream function Received... \n")
		if err == io.EOF {
			log.Printf("[client] EOF close ")
			return 1
		}
		if err != nil {
			log.Printf("can not receive %v", err)
		}

		// NOTE : receive process here
		fmt.Printf("+ [grpc client][ProxyStream] data : %#v\n", resp.Data)
		fmt.Printf("+ [grpc client][ProxyStream] size : %#v\n", resp.Length)
		fmt.Println()

		if resp.Data == nil || resp.Length == 0 {
			_, _, line, _ := runtime.Caller(0)
			log.Printf("[client:%d] not available message", line+1)

		} else {

			switch resp.Data[0] {
			case C.PDU_SRXPROXY_SYNC_REQUEST:
				log.Printf("[client] Sync Request\n")

				go_sr := &Go_ProxySyncRequest{
					_type:     resp.Data[0],
					_reserved: *((*uint16)(unsafe.Pointer(&resp.Data[1]))),
					_zero:     resp.Data[3],
					_length:   *((*uint32)(unsafe.Pointer(&resp.Data[4]))),
				}
				sr := (*C.SRXPROXY_SYNCH_REQUEST)(C.malloc(C.sizeof_SRXPROXY_SYNCH_REQUEST))
				defer C.free(unsafe.Pointer(sr))
				go_sr.Pack(unsafe.Pointer(sr))
				log.Printf(" sr: %#v\n", sr)

				//void processSyncRequest_grpc(SRXPROXY_SYNCH_REQUEST* hdr)
				C.processSyncRequest_grpc(sr)

			case C.PDU_SRXPROXY_SIGN_NOTIFICATION:
				log.Printf("[client] Sign Notification\n")

				//void processSignNotify_grpc(SRXPROXY_SIGNATURE_NOTIFICATION* hdr)
				C.processSignNotify_grpc(nil)

			}
		}
	}

	return 0
}

//export RunStream
func RunStream(data []byte) uint32 {

	cli := client.cli
	fmt.Printf("client data: %#v\n", client)

	if data == nil {
		fmt.Println("#############")
		data = []byte(defaultName)
	}

	fmt.Printf("input data for stream response: %#v\n", data)

	stream, err := cli.SendAndWaitProcess(context.Background(), &pb.PduRequest{Data: data, Length: uint32(len(data))})
	if err != nil {
		log.Printf("open stream error %v", err)
	}

	ctx := stream.Context()
	done := make(chan bool)
	//var r pb.PduResponse

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				log.Printf("[client] EOF close ")
				return
			}
			if err != nil {
				log.Printf("can not receive %v", err)
			}

			// NOTE : receive process here
			fmt.Printf("+ data : %#v\n", resp.Data)
			fmt.Printf("+ size : %#v\n", resp.Length)
			fmt.Printf("+ status: %#v\n", resp.ValidationStatus)
			fmt.Println()
			//r = resp

			if resp.Data == nil && resp.Length == 0 {
				_, _, line, _ := runtime.Caller(0)
				log.Printf("[client:%d] close stream ", line+1)
				//done <- true
				//stream.CloseSend()
				close(done)
			}
		}
	}()

	go func() {
		<-ctx.Done()
		log.Printf("+ Client Context Done")
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		fmt.Printf("+ client context done\n")
		//close(done)
	}()

	<-done
	//log.Printf("Finished with Resopnse valie: %d", uint32(resp.ValidationStatus))
	log.Printf("Finished with Resopnse valie")
	//fmt.Printf("Finished with Resopnse valie: %d", uint32(resp.ValidationStatus))
	//close(ctx.Done)

	return 0
	//return uint32(resp.ValidationStatus)
}

//export RunProxyVerify
func RunProxyVerify(data []byte, grpcClientID uint32) uint32 {

	cli := client.cli
	fmt.Printf("client data: %#v\n", client)

	if data == nil {
		fmt.Println("#############")
		data = []byte(defaultName)
	}

	fmt.Printf("input data for stream response: %#v\n", data)

	stream, err := cli.ProxyVerifyStream(context.Background(),
		&pb.ProxyVerifyRequest{
			Data:         data,
			Length:       uint32(len(data)),
			GrpcClientID: grpcClientID,
		},
	)
	if err != nil {
		log.Printf("open stream error %v", err)
	}

	ctx := stream.Context()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				// NOTE: remove close below to prevent panic from calling close again
				//		 which was done by the case of type=0, length=0
				//close(done)
				log.Printf("[client] EOF close ")
				return
			}
			if err != nil {
				log.Printf("can not receive %v", err)
			}

			fmt.Println()
			fmt.Printf("+ data : %#v\n", resp)
			fmt.Printf("+ size : %#v\n", resp.Length)

			if resp.Type == 0 && resp.Length == 0 {
				_, _, line, _ := runtime.Caller(0)
				log.Printf("[client:%d] close stream \n", line+1)
				close(done)
			} else {

				go_vn := &Go_ProxyVerifyNotify{
					_type:         uint8(resp.Type),
					_resultType:   uint8(resp.ResultType),
					_roaResult:    uint8(resp.RoaResult),
					_bgpsecResult: uint8(resp.BgpsecResult),
					_length:       resp.Length,
					_requestToken: resp.RequestToken,
					_updateID:     resp.UpdateID,
				}
				vn := (*C.SRXPROXY_VERIFY_NOTIFICATION)(C.malloc(C.sizeof_SRXPROXY_VERIFY_NOTIFICATION))
				defer C.free(unsafe.Pointer(vn))
				go_vn.Pack(unsafe.Pointer(vn))
				log.Printf(" vn: %#v\n", vn)

				// to avoid runtime: address space conflict:
				//			and fatal error: runtime: address space conflict
				//	    NEED to make a shared library at the client side same way at server side
				C.processVerifyNotify_grpc(vn)

			}
		}
	}()

	go func() {
		<-ctx.Done()
		log.Printf("+ Client Context Done")
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		fmt.Printf("+ client context done\n")
		//close(done)
	}()

	<-done
	log.Printf("Finished with Resopnse value")
	return 0
}

func main() {
	///* FIXME XXX

	log.Printf("main start Init(%s)\n", address)
	rv := InitSRxGrpc(address)
	if rv != true {
		log.Printf(" Init Error ")
		return
	}
	//defer client.conn.Close()

	/*
		// TODO: construct Proxy Verify Request data structure and nested structures too
		req := pb.ProxyVerifyV4Request{}
		req.Common = &pb.ProxyBasicHeader{
			Type:         0x03,
			Flags:        0x83,
			RoaResSrc:    0x01,
			BgpsecResSrc: 0x01,
			Length:       0xa9,
			RoaDefRes:    0x03,
			BgpsecDefRes: 0x03,
			PrefixLen:    0x18,
			RequestToken: binary.BigEndian.Uint32([]byte{0x01, 0x00, 0x00, 0x00}),
		}
		req.PrefixAddress = &pb.IPv4Address{
			AddressOneof: &pb.IPv4Address_U8{
				U8: []byte{0x064, 0x01, 0x00, 0x00},
			},
		}
		req.OriginAS = binary.BigEndian.Uint32([]byte{0x00, 0x00, 0xfd, 0xf3})
		req.BgpsecLength = binary.BigEndian.Uint32([]byte{0x00, 0x00, 0x00, 0x71})

		req.BgpsecValReqData = &pb.BGPSECValReqData{
			NumHops: binary.BigEndian.Uint32([]byte{0x00, 0x00, 0x00, 0x01}),
			AttrLen: binary.BigEndian.Uint32([]byte{0x00, 0x00, 0x00, 0x6d}),
			ValPrefix: &pb.SCA_Prefix{
				Afi:  binary.BigEndian.Uint32([]byte{0x00, 0x00, 0x00, 0x6d}),
				Safi: binary.BigEndian.Uint32([]byte{0x00, 0x00, 0x00, 0x6d}),
			},
			ValData: &pb.BGPSEC_DATA_PTR{
				LocalAs: binary.BigEndian.Uint32([]byte{0x00, 0x00, 0xfd, 0xed}),
			},
		}
		fmt.Printf(" request: %#v\n", req)
		log.Fatalf("terminate here")
	*/

	// NOTE: SRx Proxy Hello
	log.Printf("Hello Request\n")
	buff_hello_request :=
		[]byte{0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x14, 0x0, 0x0, 0x0, 0x5, 0x0, 0x0, 0xfd, 0xed, 0x0, 0x0, 0x0, 0x0}
	res, grpcClientID := RunProxyHello(buff_hello_request)
	//r := Run(buff_hello_request)
	log.Printf("Transferred: %#v, proxyID: %d\n\n", res, grpcClientID)
	//*/

	// NOTE: SRx Proxy GoodBye Stream Test
	go func() {
		log.Printf("GoodBye Stream Request\n")
		buff_goodbye_stream_request := []byte{0x02, 0x03, 0x84, 0x0, 0x0, 0x0, 0x0, 0x08}
		result := RunProxyGoodByeStream(buff_goodbye_stream_request, grpcClientID)
		log.Println("result:", result)
	}()

	//time.Sleep(2 * time.Second)

	// NOTE: SRx Proxy Verify
	log.Printf("Verify Request\n")
	buff_verify_req := []byte{0x03, 0x83, 0x01, 0x01, 0x00, 0x00, 0x00, 0xa9, 0x03, 0x03, 0x00, 0x18,
		0x00, 0x00, 0x00, 0x01, 0x64, 0x01, 0x00, 0x00, 0x00, 0x00, 0xfd, 0xf3, 0x00, 0x00, 0x00, 0x71,
		0x00, 0x01, 0x00, 0x6d, 0x00, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xfd, 0xed, 0x00, 0x00, 0xfd, 0xf3,
		0x90, 0x21, 0x00, 0x69, 0x00, 0x08, 0x01, 0x00, 0x00, 0x00, 0xfd, 0xf3, 0x00, 0x61, 0x01, 0xc3,
		0x04, 0x33, 0xfa, 0x19, 0x75, 0xff, 0x19, 0x31, 0x81, 0x45, 0x8f, 0xb9, 0x02, 0xb5, 0x01, 0xea,
		0x97, 0x89, 0xdc, 0x00, 0x48, 0x30, 0x46, 0x02, 0x21, 0x00, 0xbd, 0x92, 0x9e, 0x69, 0x35, 0x6e,
		0x7b, 0x6c, 0xfe, 0x1c, 0xbc, 0x3c, 0xbd, 0x1c, 0x4a, 0x63, 0x8d, 0x64, 0x5f, 0xa0, 0xb7, 0x20,
		0x7e, 0xf3, 0x2c, 0xcc, 0x4b, 0x3f, 0xd6, 0x1b, 0x5f, 0x46, 0x02, 0x21, 0x00, 0xb6, 0x0a, 0x7c,
		0x82, 0x7f, 0x50, 0xe6, 0x5a, 0x5b, 0xd7, 0x8c, 0xd1, 0x81, 0x3d, 0xbc, 0xca, 0xa8, 0x2d, 0x27,
		0x47, 0x60, 0x25, 0xe0, 0x8c, 0xda, 0x49, 0xf9, 0x1e, 0x22, 0xd8, 0xc0, 0x8e}
	RunProxyVerify(buff_verify_req, grpcClientID)
	//RunStream(buff_verify_req)

	// NOTE: SRx PROY GOODBYE
	goGB := &Go_ProxyGoodBye{
		_type:       0x02,
		_keepWindow: binary.BigEndian.Uint16([]byte{0x83, 0x03}), // 0x03 0x84 : 900
		_length:     binary.BigEndian.Uint32([]byte{0x8, 0x00, 0x00, 0x00}),
	}

	gb := (*C.SRXPROXY_GOODBYE)(C.malloc(C.sizeof_SRXPROXY_GOODBYE))
	defer C.free(unsafe.Pointer(gb))

	goGB.Pack(unsafe.Pointer(gb))
	log.Printf(" gb: %#v\n", gb)

	status := RunProxyGoodBye(*gb, uint32(grpcClientID))
	log.Printf(" GoodBye response status: %#v\n", status)

	/* FIXME
	data := []byte(defaultName)
	data2 := []byte{0x10, 0x11, 0x40, 0x42}
	data3 := []byte{0x10, 0x11, 0x40, 0x42, 0xAB, 0xCD, 0xEF}

	r := Run(data)
	log.Printf("Transferred: %#v\n\n", r)

	r = Run(data2)
	log.Printf("Transferred: %#v\n\n", r)

	r = RunStream(data3)
	log.Printf("Transferred: %#v\n\n", r)
	*/
}

/* NOTE

TODO 1: init function - for receiving client (*grpc.ClientConn)
		--> maybe good to use a global variable for client

TODO 2

*/
