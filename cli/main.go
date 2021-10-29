package main

import (
	"context"
	pb "github.com/869413421/laracom/service/proto/demo"
	"github.com/micro/go-micro/v2"
	"log"
)

func main() {
	service := micro.NewService(micro.Name("laravel.demo.cli"))
	service.Init()

	client := pb.NewDemoService("laravel.demo.service", service.Client())
	rsp, err := client.SayHello(context.TODO(), &pb.DemoRequest{Name: "学院君"})
	if err != nil {
		log.Fatalf("服务调用失败：%v", err)
		return
	}
	log.Println(rsp.Text)
}