package main

import (
	"context"
	"io"
	"log"

	pb "github.com/ddiogoo/grpc/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doCalculator was invoked")
	stream, err := c.Primes(context.Background(), &pb.PrimesRequest{Number: 120})
	if err != nil {
		log.Fatalf("Error while calling Primes: %v\n", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}
		log.Printf("GreetManyTimes: %d\n", msg.Result)
	}
}
