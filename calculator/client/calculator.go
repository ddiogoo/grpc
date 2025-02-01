package main

import (
	"context"
	"log"

	pb "github.com/ddiogoo/grpc/calculator/proto"
)

func doCalculator(c pb.CalculatorServiceClient) {
	log.Println("doCalculator was invoked")
	addResult, err := c.Add(context.Background(), &pb.CalculatorRequest{FirstNumber: 22, SecondNumber: 2})
	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}
	log.Printf("Sum: %d\n", addResult.Result)

	subResult, err := c.Subtract(context.Background(), &pb.CalculatorRequest{FirstNumber: 22, SecondNumber: 2})
	if err != nil {
		log.Fatalf("Could not subtract: %v\n", err)
	}
	log.Printf("Subtraction: %d\n", subResult.Result)

	mulResult, err := c.Multiply(context.Background(), &pb.CalculatorRequest{FirstNumber: 22, SecondNumber: 2})
	if err != nil {
		log.Fatalf("Could not multiply: %v\n", err)
	}
	log.Printf("Multiplication: %d\n", mulResult.Result)

	divResult, err := c.Divide(context.Background(), &pb.CalculatorRequest{FirstNumber: 22, SecondNumber: 2})
	if err != nil {
		log.Fatalf("Could not divide: %v\n", err)
	}
	log.Printf("Division: %d\n", divResult.Result)
}
