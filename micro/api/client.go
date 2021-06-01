package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-plugins/client/http/v2"
)

func main() {
	cli := http.NewClient(
		client.Selector(
			selector.NewSelector(
				selector.SetStrategy(selector.Random),
				selector.Registry(
					etcd.NewRegistry(registry.Addrs("127.0.0.1:2379")),
				),
			),
		),
		client.ContentType("application/json"),
	)

	req := cli.NewRequest("api_server", "/prods", nil)

	var res map[string]interface{}

	if err := cli.Call(context.TODO(), req, &res); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
