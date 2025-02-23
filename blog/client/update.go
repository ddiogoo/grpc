package main

import (
	"context"
	"log"

	pb "github.com/ddiogoo/grpc/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("--- updateBlog was invoked ---")
	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "New Author",
		Title:    "New Title",
		Content:  "New Content",
	}
	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error while updating blog: %v", err)
	}
	log.Printf("Blog has been updated")
}
