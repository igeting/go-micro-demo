package main

import (
	"context"
	"go-micro-demo/python/grpc/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type action struct{}

func (s *action) DoFormat(ctx context.Context, req *pb.ActionRequest) (*pb.ActionResponse, error) {
	return &pb.ActionResponse{Text: req.Text, Age: 30, Result: []*pb.ActionResponse_Result{
		&pb.ActionResponse_Result{Title: "pkg", Url: "https://pkg.go.dev/"},
	}}, nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterFormatDataServer(s, new(action))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
