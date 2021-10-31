package handler

import (
	"context"
	pb "github.com/869413421/laracom/user-service/proto/user"
	"github.com/869413421/laracom/user-service/repo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repo.UserRepository
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
	hashPassword,err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
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
