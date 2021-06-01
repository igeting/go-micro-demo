package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"go-micro-demo/micro/pb"
	"time"
)

var (
	ConsulClusterServer = []string{"127.0.0.1:2379"}
)

type StudentManager struct{}

func (s *StudentManager) GetStudent(ctx context.Context, request *pb.StudentRequest, response *pb.Student) error {
	studentMap := map[string]pb.Student{
		"jack": {Name: "jack", Classes: "信息工程", Grade: 90},
		"tony": {Name: "tony", Classes: "工程设计", Grade: 80},
		"nick": {Name: "nick", Classes: "电子工程", Grade: 85},
	}

	if request.Name == "" {
		return errors.New("请求参数错误，请重新请求。")
	}
	student := studentMap[request.Name]

	if student.Name != "" {
		fmt.Println(student.Name, student.Classes, student.Grade)
		*response = student
		return nil
	}
	return errors.New("未查询到学生信息。")
}

//服务端调用服务
func GetStudent(service micro.Service) {
	time.Sleep(time.Second * 5)
	student, err := pb.NewStudentService("rpc_server", service.Client()).GetStudent(context.TODO(), &pb.StudentRequest{Name: "jack"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("student:%+v\n", student)
}

func main() {
	service := micro.NewService(
		micro.Name("rpc_server"),
		micro.Registry(etcd.NewRegistry(registry.Addrs(ConsulClusterServer...))),
	)

	service.Init()

	pb.RegisterStudentServiceHandler(service.Server(), new(StudentManager))

	go GetStudent(service)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
