package main

import (
	"context"
	"log"

	pb "github.com/ddiogoo/grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("UpdateBlog was invoked with: %v\n", in)
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "cannot parse ID")
	}
	data := &BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}
	filter := bson.M{"_id": oid}
	update := bson.M{"$set": data}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, status.Error(codes.Internal, "cannot update blog")
	}
	if res.MatchedCount == 0 {
		return nil, status.Error(codes.NotFound, "blog not found")
	}
	return &emptypb.Empty{}, nil
}
