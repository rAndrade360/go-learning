package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A int
	B int
}

type Arith int

func (a *Arith) Sum(args *Args, reply *int) error {
	log.Println("Request received. Args: ", args.A, args.B)
	*reply = args.A + args.B
	return nil
}

func (a *Arith) Multiply(args *Args, reply *int) error {
	log.Println("Request received. Args: ", args.A, args.B)
	*reply = args.A * args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	http.Serve(listen, nil)
}
