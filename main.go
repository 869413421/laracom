package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"log"
	pb "service/proto/demo"
)

const (
	address  = "localhost:9091"
	grpcPort = ":9999"
	httpPort = ":8089"
	appName  = "Service"
)

type DemoServiceHandler struct {
}

func (s *DemoServiceHandler) SayHello(ctx context.Context, req *pb.DemoRequest, rsp *pb.DemoResponse) error {
	rsp.Text = "你好, " + req.Name
	return nil
}

func main() {
	//1.创建一个服务，服务名必须何proto中的pack声明一直
	service := micro.NewService(micro.Name("laravel.demo.service"))
	service.Init()

	//2.注册一个服务
	pb.RegisterDemoServiceHandler(service.Server(), &DemoServiceHandler{})
	if err := service.Run(); err != nil {
		log.Fatalf("服务启动失败：%v", err)
	}
}
