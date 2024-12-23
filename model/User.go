package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `gorm:"not null;unique" json:"email"`
	Password     string `json:"-"`
	RefreshToken string `json:"refresh-token"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	RefreshToken string `json:"refresh-token"`
	AccessToken  string `json:"access-token"`
}
