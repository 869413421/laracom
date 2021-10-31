package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

// CreateConnection 创建gorm链接
func CreateConnection() (*gorm.DB, error) {
	//1.从系统变量获取数据库配置
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	database := os.Getenv("DB_DATABASE")
	password := os.Getenv("DB_PASSWORD")

	return gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user, password, host, database,
		),
	)
}
