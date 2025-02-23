package main

import (
	"context"
	"fmt"

	pb "github.com/ddiogoo/grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListBlogs(in *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {
	fmt.Println("--- ListBlogs was invoked ---")
	cur, err := collection.Find(context.Background(), primitive.D{})
	if err != nil {
		return status.Error(codes.NotFound, fmt.Sprintf("Error while reading blogs: %v\n", err))
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		data := &pb.Blog{}
		if err := cur.Decode(data); err != nil {
			return status.Error(codes.Internal, fmt.Sprintf("Error while decoding blog: %v\n", err))
		}
		if err := stream.Send(data); err != nil {
			return status.Error(codes.Internal, fmt.Sprintf("Error while sending blog: %v\n", err))
		}
	}
	if err = cur.Err(); err != nil {
		return status.Error(codes.Internal, fmt.Sprintf("Error while iterating blogs: %v\n", err))
	}
	return nil
}
