package models

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}
