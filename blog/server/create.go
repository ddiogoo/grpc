package main

import (
	"context"
	"log"

	pb "github.com/ddiogoo/grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("CreateBlog was invoked with: %v\n", in)
	data := &BlogItem{
		AuthorID: in.GetAuthorId(),
		Title:    in.GetTitle(),
		Content:  in.GetContent(),
	}
	res, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to insert blog: %v", err)
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, "Failed to convert ObjectID: %v", err)
	}
	return &pb.BlogId{Id: oid.Hex()}, nil
}
