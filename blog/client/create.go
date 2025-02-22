package main

import (
	"context"
	"log"

	pb "github.com/ddiogoo/grpc/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("--- createBlog was invoked ---")
	blog := &pb.Blog{
		AuthorId: "Diogo",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}
	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Error while creating blog: %v", err)
	}
	log.Printf("Blog has been created: %v", res)
	return res.GetId()
}
