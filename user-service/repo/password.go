package repo

import (
	. "github.com/869413421/laracom/user-service/model"
	"github.com/jinzhu/gorm"
)

type PasswordResetInterface interface {
	Create(reset *PasswordReset) error
	Delete(reset *PasswordReset) error
	GetByToken(token string) (*PasswordReset, error)
	GetByEmail(email string) (*PasswordReset, error)
}

type PasswordResetRepository struct {
	Db *gorm.DB
}

// Create 创建获取记录
func (repo *PasswordResetRepository) Create(reset *PasswordReset) error {
	if err := repo.Db.Create(reset).Error; err != nil {
		return err
	}
	return nil
}

// GetByToken 根据token获取重置记录
func (repo *PasswordResetRepository) GetByToken(token string) (*PasswordReset, error) {
	reset := &PasswordReset{}
	if err := repo.Db.Where("token = ?", token).First(reset).Error; err != nil {
		return nil, err
	}

	return reset, nil
}

// GetByEmail 根据emil获取重置记录
func (repo *PasswordResetRepository) GetByEmail(email string) (*PasswordReset, error) {
	reset := &PasswordReset{}
	if err := repo.Db.Where("email = ?", email).First(reset).Error; err != nil {
		return nil, err
	}

	return reset, nil
}

// Delete 删除重置记录
func (repo *PasswordResetRepository) Delete(reset *PasswordReset) error {
	err := repo.Db.Delete(&reset).Error
	return err
}
