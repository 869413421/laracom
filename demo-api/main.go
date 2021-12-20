package main

import (
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
		"text": "你好，学院君",
	})
}

func (s *Say) Hello(c *gin.Context) {
	log.Println("Received Say.Hello API request")

	name := c.Param("name")
	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		log.Println("get context err")
	}
	response, err := cli.SayHello(ctx, &demo.DemoRequest{
		Name: name,
	})

	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, response)
}

func main() {
	var name = "laracom.api.demo"
	// 初始化追踪器
	gin2micro.SetSamplingFrequency(50)
	t, io, err := tracer.NewTracer(name, os.Getenv("MICRO_TRACE_SERVER"))
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// Create service
	service := web.NewService(
		web.Name(name),
	)

	service.Init()

	// setup Demo Server Client
	cli = demo.NewDemoService("laracom.service.service", client.DefaultClient)

	// Create RESTful handler (using Gin)
	say := new(Say)
	router := gin.Default()
	r := router.Group("/demo")
	r.Use(gin2micro.TracerWrapper)
	r.GET("/hello", say.Anything)
	r.GET("/hello/:name", say.Hello)

	// Register Handler
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
