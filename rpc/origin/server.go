package main

import (
	"go-micro-demo/model"
	"log"
	"net"
	"net/rpc"
)

func main() {
	server := rpc.NewServer()
	err := server.Register(new(model.Arith))
	if err != nil {
		log.Fatalln(err)
	}

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	server.Accept(listen)
}
