package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Arith struct{}

type Args struct {
	A int
	B int
}

type Quotient struct {
	Pro int
	Quo int
	Rem int
}

func (a *Arith) Multiply(req Args, res *Quotient) error {
	res.Pro = req.A * req.B
	return nil
}

func (a *Arith) Divide(req Args, res *Quotient) error {
	if req.B == 0 {
		return errors.New("divide by zero")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}

func main() {
	server := rpc.NewServer()
	server.Register(new(Arith))

	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("listen error: %s", err)
	} else {
		log.Println("rpc listening...")
	}

	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Printf("rpc server accept: %s", err)
			return
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
