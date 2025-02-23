package main

import (
	pb "github.com/ddiogoo/grpc/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)
	id := createBlog(c)
	readBlog(c, id)
	updateBlog(c, id)
	readBlog(c, id)
}
