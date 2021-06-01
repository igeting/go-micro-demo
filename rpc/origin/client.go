package main

import (
	"go-micro-demo/model"
	"log"
	"net/rpc"
)

func main() {
	cli, err := rpc.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	req := model.Args{A: 9, B: 2}
	res := new(model.Quotient)

	err = cli.Call("Arith.Multiply", req, res)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}
