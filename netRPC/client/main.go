package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	N int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := Args{7}
	var reply int
	err = client.Call("Calculator.Square", args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	fmt.Printf("The square of %d is %d", args.N, reply) // Outputs: The square of 7 is 49
}
