package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ddiogoo/grpc/calculator/proto"
)

func doCalculatorAvg(c pb.CalculatorServiceClient) {
	log.Println("doCalculatorAvg was invoked")
	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}
	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error while calling CalculatorAvg: %v\n", err)
	}
	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from CalculatorAvg: %v\n", err)
	}
	log.Printf("CalculatorAvg: %f\n", res.Result)
}
