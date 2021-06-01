package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"go-micro-demo/micro/pb"
	"time"
)

var (
	ConsulClusterClient = []string{"127.0.0.1:2379"}
)

func main() {
	service := micro.NewService(
		micro.Name("rpc_client"),
		micro.Selector(selector.NewSelector(selector.SetStrategy(selector.Random))),
		micro.Registry(etcd.NewRegistry(registry.Addrs(ConsulClusterClient...))),
	)

	service.Init()

	studentService := pb.NewStudentService("rpc_server", service.Client())

	res, err := studentService.GetStudent(
		context.TODO(),
		&pb.StudentRequest{Name: "jack"},
		func(o *client.CallOptions) {
			o.RequestTimeout = time.Second * 30
			o.DialTimeout = time.Second * 30
		},
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("res:%+v\n", res)
}
