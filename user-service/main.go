package main

import (
	"fmt"
	db "github.com/869413421/laracom/user-service/db"
	"github.com/869413421/laracom/user-service/handler"
	"github.com/869413421/laracom/user-service/model"
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
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.PasswordReset{})
	userRepo := repo.UserRepository{Db: db}
	token := &service.TokenService{Repo: &userRepo}
	resetRepo := &repo.PasswordResetRepository{Db: db}
	srvHandler := &handler.UserService{Repo: userRepo, Token: token,ResetRepo: resetRepo}

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
