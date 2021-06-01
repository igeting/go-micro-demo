package main

import (
	"fmt"
	"go-micro-demo/python/grpc/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"reflect"
)

const (
	address = "127.0.0.1:8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFormatDataClient(conn)

	r, err := c.DoFormat(context.Background(), &pb.ActionRequest{Text: "test", Corpus: pb.ActionRequest_NEWS})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(r.Age)
	fmt.Println(reflect.TypeOf(r.Result))
	for k, v := range r.Result {
		fmt.Println(k, v)
		fmt.Println(v.Snippets)
	}
}
