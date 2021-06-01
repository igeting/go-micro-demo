package main

import (
	"fmt"
	"go-micro-demo/model"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	rpc.Register(new(model.Arith))
	rpc.HandleHTTP()

	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("listen error: %s", err)
	}

	fmt.Println("start server")

	if err := http.Serve(lis, nil); err != nil {
		fmt.Println(err)
	}
}
