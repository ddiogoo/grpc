package main

import (
	"context"
	"log"

	pb "github.com/ddiogoo/grpc/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("--- deleteBlog was invoked ---")
	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("Error while deleting blog: %v\n", err)
	}
	log.Printf("Blog was deleted with id: %v\n", id)
}
