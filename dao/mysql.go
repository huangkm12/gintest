package dao

import (
	"bubble/modles"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitMySQL() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test02?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	err = DB.DB().Ping()
	//
	DB.AutoMigrate(&modles.Todo{})
	return
}