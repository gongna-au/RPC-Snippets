package main

import (
	"log"
	"net"
	"net/rpc"
)

type Args struct {
	N int
}

type Calculator int

func (c *Calculator) Square(args *Args, reply *int) error {
	*reply = args.N * args.N
	return nil
}

func main() {
	calculator := new(Calculator)
	server := rpc.NewServer()
	server.Register(calculator)

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	server.Accept(l)
}
