package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, name string, msg *string) error {
	*msg = "Hello " + name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("json_server"),
		micro.Registry(etcd.NewRegistry(registry.Addrs("iopening.cn:2379"))),
	)
	service.Init()

	micro.RegisterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
