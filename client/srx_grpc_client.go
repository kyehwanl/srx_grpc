package main

import "C"

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "srx_grpc"
)

const (
	address     = "localhost:50000"
	defaultName = "RPKI_DATA"
)

//export Run
func Run(data []byte) uint32 {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSRxApiClient(conn)

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
	log.Printf("Finished with Resopnse valie: %d", uint32(r.ValidationStatus))
	return uint32(r.ValidationStatus)
}

func main() {
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
