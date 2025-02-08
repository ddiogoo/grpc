package main

import (
	"io"
	"log"

	pb "github.com/ddiogoo/grpc/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max was invoked")
	sentNumbers := make(map[int]struct{})
	var currentMax int
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		number := int(req.Number)
		if _, exists := sentNumbers[number]; exists {
			continue
		}
		sentNumbers[number] = struct{}{}
		if len(sentNumbers) == 1 || number > currentMax {
			currentMax = number
			err = stream.Send(&pb.MaxResponse{Result: req.Number})
			if err != nil {
				log.Fatalf("Error while sending data to client: %v\n", err)
			}
		}
	}
}
