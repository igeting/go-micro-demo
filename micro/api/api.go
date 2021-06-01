package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"io/ioutil"
	"net/http"
)

func main() {
	service, err := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	).GetService("api_server")
	if err != nil {
		fmt.Println(err)
	}

	next := selector.Random(service)
	node, err := next()
	if err != nil {
		fmt.Println(err)
	}

	url := "http://" + node.Address + "/prods"
	res, err := http.Post(url, "application/json", nil)
	if err != nil {
		fmt.Println(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
	defer res.Body.Close()
}
