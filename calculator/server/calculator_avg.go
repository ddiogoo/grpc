package main

import (
	"io"
	"log"

	pb "github.com/ddiogoo/grpc/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg was invoked")
	var sum int32 = 0
	var count int32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(count),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		log.Printf("Receiving req: %v\n", req.Number)
		sum += req.Number
		count += 1
	}
}
