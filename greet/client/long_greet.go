package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ddiogoo/grpc/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")
	reqs := []*pb.GreetRequest{
		{FirstName: "Diogo"},
		{FirstName: "Martins"},
		{FirstName: "Assis"},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}
	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}
	log.Println(res.GetResult())
	log.Printf("LongGreet: %s\n", res.Result)
}
