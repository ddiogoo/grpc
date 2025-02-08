package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/ddiogoo/grpc/calculator/proto"
)

func doCalculatorMax(c pb.CalculatorServiceClient) {
	log.Println("doCalculatorMax was invoked")
	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}
	reqs := []*pb.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}
	numbers := make([]int32, len(reqs))
	for i, req := range reqs {
		numbers[i] = req.Number
	}
	log.Printf("Sending numbers: %v\n", numbers)
	waitc := make(chan struct{})
	go func() {
		for _, req := range reqs {
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
