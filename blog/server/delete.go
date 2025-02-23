package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ddiogoo/grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Printf("DeleteBlog was invoked with: %v\n", in)
	blogId := in.GetId()
	oid, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Error while converting id: %v\n", err))
	}
	filter := bson.M{"_id": oid}
	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Error while deleting blog: %v\n", err))
	}
	if res.DeletedCount == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Blog not found: %v\n", err))
	}
	return &emptypb.Empty{}, nil
}
