package main

import (
	"context"

	pb "github.com/ddiogoo/grpc/calculator/proto"
)

func (s *Server) Add(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	return &pb.CalculatorResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}

func (s *Server) Subtract(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	return &pb.CalculatorResponse{
		Result: in.FirstNumber - in.SecondNumber,
	}, nil
}

func (s *Server) Multiply(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	return &pb.CalculatorResponse{
		Result: in.FirstNumber * in.SecondNumber,
	}, nil
}

func (s *Server) Divide(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	return &pb.CalculatorResponse{
		Result: in.FirstNumber / in.SecondNumber,
	}, nil
}
