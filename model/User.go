package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `gorm:"not null;unique" json:"email"`
	Password     string `json:"-"`
	RefreshToken string `json:"Refresh-Token"`
}
