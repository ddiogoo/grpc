package main

import (
	"context"
	"log"

	pb "github.com/ddiogoo/grpc/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("--- readBlog was invoked ---")
	blogId := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), blogId)
	if err != nil {
		log.Fatalf("Error while reading blog: %v", err)
	}
	log.Printf("Blog has been read: %v\n", res)
	return res
}
