package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {
	service := micro.NewService(
		micro.Registry(etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))),
	)
	service.Init()

	cli := service.Client()

	req := cli.NewRequest("json_server", "Greeter.Hello", "Tom", client.WithContentType("application/json"))
	var res string
	if err := cli.Call(context.TODO(), req, &res); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
