package main

import (
	"context"
	"io"
	"log"

	pb "github.com/ddiogoo/grpc/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlogs(c pb.BlogServiceClient) {
	log.Println("--- listBlogs was invoked ---")
	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error while reading blogs: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading blog: %v", err)
		}
		log.Printf("Blog has been read: %v\n", res)
	}
}
