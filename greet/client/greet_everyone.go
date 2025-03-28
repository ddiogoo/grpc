package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/ddiogoo/grpc/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}
	reqs := []*pb.GreetRequest{
		{FirstName: "Diogo"},
		{FirstName: "Martins"},
		{FirstName: "Assis"},
	}
	waitc := make(chan struct{})
	go func() {
		for _, req := range reqs {
			log.Printf("Sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}
			log.Printf("Received: %v\n", res.Result)
		}
		close(waitc)
	}()
	<-waitc
}
