package main

import (
	"fmt"
	db "github.com/869413421/laracom/user-service/db"
	"github.com/869413421/laracom/user-service/handler"
	pb "github.com/869413421/laracom/user-service/proto/user"
	repo "github.com/869413421/laracom/user-service/repo"
	"github.com/869413421/laracom/user-service/service"
	"github.com/micro/go-micro/v2"
	"log"
)

func main() {
	//1.创建数据库链接
	db, err := db.CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB %v", err)
	}

	//2.执行数据库迁移
	db.AutoMigrate(&pb.User{})
	repo := repo.UserRepository{Db: db}
	token := &service.TokenService{Repo: &repo}
	srvHandler := &handler.UserService{Repo: repo, Token: token}

	//3.创建微服务
	srv := micro.NewService(micro.Name("laracom.service.user"), micro.Version("latest"))
	srv.Init()

	//4.注册服务
	pb.RegisterUserServiceHandler(srv.Server(), srvHandler)

	//5.启动服务
	err = srv.Run()
	if err != nil {
		fmt.Println(err)
	}
}
