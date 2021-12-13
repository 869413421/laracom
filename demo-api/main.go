package main

import (
	"context"
	"github.com/869413421/laracom/service/proto/demo"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/web"
	"log"
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
	service := web.NewService(
		web.Name("laracom.api.demo"),
	)
	service.Init()

	cli = demo.NewDemoService("laracom.service.demo",client.DefaultClient)

	// Create RESTful handler (using Gin)
	say := new(Say)
	router := gin.Default()
	router.GET("/hello", say.Anything)
	router.GET("/hello/:name", say.Hello)

	// Register Handler
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
