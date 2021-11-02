package handler

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/869413421/laracom/user-service/proto/user"
	"github.com/869413421/laracom/user-service/repo"
	"github.com/869413421/laracom/user-service/service"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserService struct {
	Repo  repo.UserRepository
	Token service.Authable
}

func (srv *UserService) Get(ctx context.Context, request *pb.User, response *pb.Response) error {
	user, err := srv.Repo.Get(request.Id)
	if err != nil {
		return err
	}
	response.User = user
	return nil
}

func (srv *UserService) GetAll(ctx context.Context, request *pb.Request, response *pb.Response) error {
	users, err := srv.Repo.GetAll()
	if err != nil {
		return err
	}
	response.Users = users
	return nil
}

func (srv *UserService) Create(ctx context.Context, request *pb.User, response *pb.Response) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	request.Password = string(hashPassword)
	if err := srv.Repo.Create(request); err != nil {
		return err
	}
	response.User = request
	return nil
}

func (srv *UserService) Auth(ctx context.Context, request *pb.User, response *pb.Token) error {
	//1.根据邮件获取用户
	log.Println("Logging in with:", request.Email, request.Password)
	user, err := srv.Repo.GetByEmail(request.Email)
	if err != nil {
		fmt.Println(user)
		fmt.Println(err)
		return err
	}

	//2.校验用户密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return err
	}

	//3.生成token
	token, err := srv.Token.Encode(user)
	if err != nil {
		return err
	}

	response.Token = token
	return nil
}

func (srv UserService) ValidateToken(ctx context.Context, request *pb.Token, response *pb.Token) error {
	claims, err := srv.Token.Decode(request.Token)
	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errors.New("无效的用户")
	}

	response.Valid = true

	return nil
}
