package main

import "C"

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
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

	fmt.Printf("status: %#v\n", r.ValidationStatus)

	//return r.ValidationStatus, err
	return uint32(r.ValidationStatus)
}

func main() {
	data := []byte(defaultName)
	data2 := []byte{0x10, 0x11, 0x40, 0x42}
	//Run(data)
	r := Run(data)
	log.Printf("Transferred: %#v\n", r)

	r = Run(data2)
	log.Printf("Transferred: %#v\n", r)
}
