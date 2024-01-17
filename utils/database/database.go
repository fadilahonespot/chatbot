package database

import (
	"fmt"
	"os"

	"github.com/fadilahonespot/chatbot/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if os.Getenv("DB_DEBUG") == "true" {
		DB = DB.Debug()
	}

	DB.AutoMigrate(&entity.User{})
	DB.AutoMigrate(&entity.Chat{})

	return DB
}
