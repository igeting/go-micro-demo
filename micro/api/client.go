package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	mh "github.com/micro/go-plugins/client/http/v2"
	"io/ioutil"
	"net/http"
	"time"
)

func cli() {
	for {
		cli := mh.NewClient(
			client.Selector(
				selector.NewSelector(
					selector.SetStrategy(selector.Random),
					selector.Registry(
						etcd.NewRegistry(registry.Addrs("iopening.cn:2379")),
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

		time.Sleep(time.Second * 3)
	}
}

func api() {
	service, err := etcd.NewRegistry(registry.Addrs("iopening.cn:2379")).GetService("api_server")
	if err != nil {
		fmt.Println(err)
	}
	for {
		node, err := selector.Random(service)()
		if err != nil {
			fmt.Println(err)
		}

		res, err := http.Post("http://"+node.Address+"/prods", "application/json", nil)
		if err != nil {
			fmt.Println(err)
		}

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(data))
		}
		res.Body.Close()

		time.Sleep(time.Second * 3)
	}
}

func main() {
	go cli()
	go api()
	select {}
}
