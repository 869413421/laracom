package laracom_service_user

import (
	"github.com/jinzhu/gorm"
	"time"
)

import (
	"github.com/satori/go.uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now().Format(time.RFC3339))
	uuid := uuid.NewV4()
	return scope.SetColumn("Id", uuid.String())
}

func (model *User) BeforeSave(scope *gorm.Scope) error {
	return scope.SetColumn("UpdatedAt",time.Now().Format(time.RFC3339))
}

func (model *PasswordReset) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("CreatedAt", time.Now().Format(time.RFC3339))
	return nil
}
