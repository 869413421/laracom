package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/869413421/laracom/user-service/model"
	pb "github.com/869413421/laracom/user-service/proto/user"
	"github.com/869413421/laracom/user-service/repo"
	"github.com/869413421/laracom/user-service/service"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/broker"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
)

const topic = "password.reset"

type UserService struct {
	Repo      repo.UserRepository
	ResetRepo repo.PasswordResetInterface
	Token     service.Authable
	PubSub    broker.Broker
}

func (srv *UserService) Get(ctx context.Context, request *pb.User, response *pb.Response) error {
	var user *model.User
	var err error
	fmt.Println(request)
	if request.Id != "" {

		id, _ := strconv.ParseUint(request.Id, 10, 64)
		user, err = srv.Repo.Get(uint(id))
	}
	if request.Email != "" {
		user, err = srv.Repo.GetByEmail(request.Email)
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if user != nil {
		response.User, _ = user.ToProtobuf()
	}
	return nil
}

func (srv *UserService) GetAll(ctx context.Context, request *pb.Request, response *pb.Response) error {
	users, err := srv.Repo.GetAll()
	if err != nil {
		return err
	}
	userItems := make([]*pb.User, len(users))
	for index, user := range users {
		userItem, _ := user.ToProtobuf()
		userItems[index] = userItem
	}
	response.Users = userItems
	return nil
}

func (srv *UserService) Create(ctx context.Context, request *pb.User, response *pb.Response) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	request.Password = string(hashPassword)
	userModel := &model.User{}
	user, _ := userModel.ToORM(request)
	if err := srv.Repo.Create(user); err != nil {
		return err
	}
	response.User, _ = user.ToProtobuf()
	return nil
}

func (srv *UserService) Update(ctx context.Context, request *pb.User, response *pb.Response) error {
	if request.Password != "" {
		// 如果密码字段不为空的话对密码进行哈希加密
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		request.Password = string(hashedPass)
	}
	if request.Id == "" {
		return errors.New("用户 ID 不能为空")
	}
	id, _ := strconv.ParseUint(request.Id, 10, 64)
	user, _ := srv.Repo.Get(uint(id))
	if user == nil {
		return errors.New("用户不存在")
	}
	user, _ = user.ToORM(request)
	if err := srv.Repo.Update(user); err != nil {
		return err
	}
	response.User, _ = user.ToProtobuf()
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
	pbUser, _ := user.ToProtobuf()
	token, err := srv.Token.Encode(pbUser)
	if err != nil {
		return err
	}

	response.Token = token
	return nil
}

func (srv *UserService) ValidateToken(ctx context.Context, request *pb.Token, response *pb.Token) error {
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

// CreatePasswordReset 创建重置密码记录
func (srv *UserService) CreatePasswordReset(tx context.Context, request *pb.PasswordReset, response *pb.PasswordResetResponse) error {
	//1.检验邮箱
	if request.Email == "" {
		return errors.New("邮箱不允许为空")
	}

	//2.创建重置记录
	resetModel := &model.PasswordReset{}
	reset, _ := resetModel.ToORM(request)
	if err := srv.ResetRepo.Create(reset); err != nil {
		return err
	}

	//3.发布消息到消息系统
	response.PasswordReset, _ = reset.ToProtobuf()
	if response.PasswordReset != nil {
		err := srv.publishEvent(response.PasswordReset)
		if err != nil {
			return err
		}
	}
	return nil
}

// ValidatePasswordResetToken 验证重置密码token
func (srv *UserService) ValidatePasswordResetToken(tx context.Context, request *pb.Token, response *pb.Token) error {
	if request.Token == "" {
		return errors.New("token 不允许为空")
	}

	_, err := srv.ResetRepo.GetByToken(request.Token)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if err == gorm.ErrRecordNotFound {
		response.Valid = false
	} else {
		response.Valid = true
	}

	return nil
}

// DeletePasswordReset 删除重置密码记录
func (srv *UserService) DeletePasswordReset(tx context.Context, request *pb.PasswordReset, response *pb.PasswordResetResponse) error {
	if request.Email == "" {
		return errors.New("邮箱 不允许为空")
	}

	reset, err := srv.ResetRepo.GetByEmail(request.Email)
	if err != nil {
		return err
	}

	if err = srv.ResetRepo.Delete(reset); err != nil {
		return err
	}

	response.PasswordReset = nil

	return nil
}

// publishEvent 发布重置密码事件
func (srv *UserService) publishEvent(reset *pb.PasswordReset) error {
	//1.JSON 编码
	body, err := json.Marshal(reset)
	if err != nil {
		return err
	}

	//2.构建broker消息
	msg := &broker.Message{
		Header: map[string]string{
			"email": reset.Email,
		},
		Body: body,
	}

	//3.通过broker 发布消息到消息系统
	err = srv.PubSub.Publish(topic, msg)
	if err != nil {
		log.Printf("[pub] failed: %v", err)
		return err
	}

	return nil
}
