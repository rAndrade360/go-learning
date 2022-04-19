package main

import (
	"log"
	"net/rpc"
)

type Args struct {
	A int
	B int
}

func main() {
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("Log err:", err.Error())
	}

	args := &Args{1, 5}

	var reply int

	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Err on call", err.Error())
	}

	log.Printf("Arith: %d * %d = %d", args.A, args.B, reply)
}
