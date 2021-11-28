package main

import (
	"fmt"
	"github.com/869413421/laracom/product-service/db"
	"github.com/869413421/laracom/product-service/handler"
	"github.com/869413421/laracom/product-service/model"
	pb "github.com/869413421/laracom/product-service/proto/product"
	"github.com/869413421/laracom/product-service/repo"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
)

func main() {
	//1.初始化数据库，执行数据迁移
	database, err := db.CreateConnection()
	defer database.Close()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	database.AutoMigrate(&model.Product{})

	// 2初始化 Repo 实例用于后续数据库操作
	productRepo := &repo.ProductRepository{Db: database}

	//3.创建微服务
	srv := micro.NewService(micro.Name("laracom.service.product"), micro.Version("latest"))
	srv.Init()

	//4.注册服务处理器
	pb.RegisterProductServiceHandler(srv.Server(), &handler.ProductService{ProductRepo: productRepo})

	//5.启动服务
	err = srv.Run()
	if err != nil {
		fmt.Println(err)
	}
}
