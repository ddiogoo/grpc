package main

import (
	"context"
	"fmt"
	"math"

	pb "github.com/ddiogoo/grpc/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	fmt.Printf("Sqrt function was invoked with %v\n", in)
	number := in.Number
	if number < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Received a negative number: %v", number))
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}
