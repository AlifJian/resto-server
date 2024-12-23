package model

import "gorm.io/gorm"

type User struct {
	name         string
	email        string
	password     string
	refreshToken string
	gorm.Model
}
