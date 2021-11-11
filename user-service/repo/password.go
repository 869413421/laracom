package repo

import (
	pb "github.com/869413421/laracom/user-service/proto/user"
	"github.com/jinzhu/gorm"
)

type PasswordResetInterface interface {
	Create(reset *pb.PasswordReset) error
	GetByToken(token string) (*pb.PasswordReset, error)
}

type PasswordResetRepository struct {
	Db *gorm.DB
}

// Create 创建获取记录
func (repo *PasswordResetRepository) Create(reset *pb.PasswordReset) error {
	if err := repo.Db.Create(reset).Error; err != nil {
		return err
	}
	return nil
}

// GetByToken 根据token获取重置记录
func (repo *PasswordResetRepository) GetByToken(token string) (*pb.PasswordReset, error) {
	reset := &pb.PasswordReset{}
	if err := repo.Db.Where("token = ?", token).First(reset).Error; err != nil {
		return nil, err
	}

	return reset, nil
}
