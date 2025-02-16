package main

import (
	"log"
	"net"
	"os"

	pb "github.com/ddiogoo/grpc/greet/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	godotenv.Load()
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	log.Printf("Server listening at %s", addr)

	opts := []grpc.ServerOption{}
	tsl := os.Getenv("ENABLE_TLS") == "true"
	if tsl {
		certFile := os.Getenv("CERT_FILE_PATH")
		key := os.Getenv("KEY_FILE_PATH")
		creds, err := credentials.NewServerTLSFromFile(certFile, key)
		if err != nil {
			log.Fatalf("Failed loading certificates: %v", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}
	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
