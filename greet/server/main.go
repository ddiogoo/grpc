package main

import (
	"log"
	"net"

	pb "github.com/ddiogoo/grpc/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	log.Printf("Server listening at %s", addr)

	s := grpc.NewServer()
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
