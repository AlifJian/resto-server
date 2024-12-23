package database

import (
	"fmt"
	"os"

	"gihtub.com/AlifJian/resto-server/util"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db gorm.DB

func InitDb() {
	defer util.Catch()

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

	db, errOpenDatabase := gorm.Open(mysql.Open(dsn))
	if errOpenDatabase != nil {
		panic("Database > " + errOpenDatabase.Error())
	}

	Db = *db

}
