package repo

import (
	. "github.com/869413421/laracom/user-service/model"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	Create(user *User) error
	Get(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
	GetAll() ([]*User, error)
}

type UserRepository struct {
	Db *gorm.DB
}

// Create 创建用户
func (repo *UserRepository) Create(user *User) error {
	if err := repo.Db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// Update 更新用户信息
func (repo *UserRepository) Update(user *User) error {
	if err := repo.Db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

// Get 根据主键获取
func (repo *UserRepository) Get(id uint) (*User, error) {
	user := &User{}
	user.ID = id
	if err := repo.Db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetByEmail 根据邮箱获取
func (repo *UserRepository) GetByEmail(email string) (*User, error) {
	user := &User{}
	if err := repo.Db.Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetAll 获取所有用户
func (repo *UserRepository) GetAll() ([]*User, error) {
	var users []*User
	if err := repo.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
