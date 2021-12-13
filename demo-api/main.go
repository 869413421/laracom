package main

import (
	"context"
	"github.com/869413421/laracom/common/tracer"
	"github.com/869413421/laracom/common/wrapper/tracer/opentracing/gin2micro"
	"github.com/869413421/laracom/service/proto/demo"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/web"
	"github.com/opentracing/opentracing-go"
	"log"
	"os"
)

var cli demo.DemoService
type Say struct {

}

func (s *Say) Anything(c *gin.Context) {
	log.Print("Received Say.Anything API request")
	c.JSON(200, map[string]string{
		"message": "你好，学院君",
	})
}

func (s *Say) Hello(c *gin.Context) {
	log.Print("Received Say.Hello API request")

	name := c.Param("name")

	response, err := cli.SayHello(context.TODO(), &demo.DemoRequest{
		Name: name,
	})

	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, response)
}

func main() {

	// 初始化追踪器
	gin2micro.SetSamplingFrequency(50)
	t, io, err := tracer.NewTracer(name, os.Getenv("MICRO_TRACE_SERVER"))
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	service := web.NewService(
		web.Name("laracom.api.demo"),
	)
	service.Init()

	cli = demo.NewDemoService("laracom.service.service",client.DefaultClient)

	// Create RESTful handler (using Gin)
	say := new(Say)
	router := gin.Default()
	router.GET("/hello", say.Anything)
	router.GET("/hello/:name", say.Hello)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(200, "111111111111111111111")
	})

	// Register Handler
	service.Handle("/", router)


	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
