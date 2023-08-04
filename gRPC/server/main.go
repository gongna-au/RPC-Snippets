package main

import (
	"context"
	"log"
	"net"

	"github.com/RPC-Snippets/gRPC/pb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Square(ctx context.Context, in *pb.SquareRequest) (*pb.SquareResponse, error) {
	return &pb.SquareResponse{Result: in.Number * in.Number}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
