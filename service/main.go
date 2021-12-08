package main

import (
	"context"
	pb "github.com/869413421/laracom/service/proto/demo"
	"github.com/869413421/laracom/service/trace"
	"github.com/micro/go-micro/v2"
	traceplugin "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"log"
	"os"
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

	// 初始化全局服务追踪
	t, io, err := trace.NewTracer("laracom.service.service", os.Getenv("MICRO_TRACE_SERVER"))
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	service := micro.NewService(
		micro.Name("laravel.service.demo"),
		micro.WrapHandler(traceplugin.NewHandlerWrapper(opentracing.GlobalTracer())), // 基于 jaeger 采集追踪数据
	)
	service.Init()

	//2.注册一个服务
	pb.RegisterDemoServiceHandler(service.Server(), &DemoServiceHandler{})
	if err := service.Run(); err != nil {
		log.Fatalf("服务启动失败：%v", err)
	}
}
