package main

import (
	"fmt"
	"go-micro-demo/model"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(new(model.Arith))

	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Printf("listen error: %s", err)
	}

	fmt.Println("start server")

	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}

		go func(conn net.Conn) {
			fmt.Println("new client in coming")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
