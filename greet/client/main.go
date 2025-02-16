package main

import (
	"log"
	"os"
	"time"

	pb "github.com/ddiogoo/grpc/greet/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50051"

func main() {
	godotenv.Load()
	opts := []grpc.DialOption{}
	tsl := os.Getenv("ENABLE_TLS") == "true"
	if tsl {
		certFile := os.Getenv("CERT_CLIENT_FILE_PATH")
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("Error while loading CA certificates: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)
	//doGreet(c)
	//doGreetManyTimes(c)
	//doLongGreet(c)
	// doGreetEveryone(c)
	//doGreetWithDeadline(c, 5*time.Second)
	doGreetWithDeadline(c, 1*time.Second)
}
