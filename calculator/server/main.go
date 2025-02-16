package main

import (
	"log"
	"net"

	pb "github.com/ddiogoo/grpc/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	log.Printf("Server listening at %s", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	reflection.Register(s)
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
