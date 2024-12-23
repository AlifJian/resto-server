package database

import (
	"fmt"
	"os"

	"gihtub.com/AlifJian/resto-server/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db gorm.DB

func InitDb() {

	errLoadEnv := godotenv.Load()

	if errLoadEnv != nil {
		panic("Database > " + errLoadEnv.Error())
	}

	USER := os.Getenv("USER")
	PASSWORD := os.Getenv("PASSWORD")
	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")
	DATABASE_NAME := os.Getenv("DATABASE_NAME")
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`, USER, PASSWORD, HOST, PORT, DATABASE_NAME)

	db, errOpenDatabase := gorm.Open(mysql.Open(dsn), &gorm.Config{TranslateError: true})
	if errOpenDatabase != nil {
		panic("Database > " + errOpenDatabase.Error())
	}

	db.AutoMigrate(model.User{})

	Db = *db
}
