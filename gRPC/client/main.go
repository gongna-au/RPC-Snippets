package main

import (
	"log"

	"github.com/RPC-Snippets/gRPC/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorClient(conn)

	r, err := c.Square(context.Background(), &pb.SquareRequest{Number: 7})
	if err != nil {
		log.Fatalf("could not calculate square: %v", err)
	}
	log.Printf("The square of 7 is: %d", r.Result)
}
