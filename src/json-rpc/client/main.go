package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Args struct {
	A int
	B int
}

type Quotient struct {
	Quo int
	Rem int
}

func main() {
	go2node()
}

func go2go() {
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)

	var quote Quotient
	err = client.Call("Arith.Divide", args, &quote)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d / %d = %d remainder %d\n", args.A, args.B, quote.Quo, quote.Rem)
}

func go2node() {
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := Args{17, 8}
	var reply int
	err = client.Call("add", [2]interface{}{1, 2}, &reply)
	if err != nil {
		log.Fatal("calling:", err)
	}
	fmt.Printf("add from node: %d + %d = %d\n", args.A, args.B, reply)
}
