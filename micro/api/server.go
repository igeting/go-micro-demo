package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"go-micro-demo/model"
)

func main() {
	router := gin.Default()
	router.POST("/prods", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": model.NewProdList(3),
		})
	})

	service := web.NewService(
		web.Name("api_server"),
		web.Address(":8001"),
		web.Advertise(":8001"),
		web.Handler(router),
		web.Registry(etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))),
		web.Metadata(map[string]string{"protocol": "http"}),
	)

	service.Init()

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
