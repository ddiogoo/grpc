package main

import (
	"fmt"
	"log"

	pb "github.com/ddiogoo/grpc/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream grpc.ServerStreamingServer[pb.GreetResponse]) error {
	log.Printf("GreetManyTimes function was invoked with %v\n", in)
	for i := 0; i < 10; i++ {
		stream.Send(&pb.GreetResponse{
			Result: fmt.Sprintf("Hello %s, this is response number %d", in.FirstName, i),
		})
	}
	return nil
}
