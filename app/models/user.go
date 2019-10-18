package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string
	Username string `gorm:"unique_index"`
	Email    string `gorm:"unique_index"`
	Password string
}

func NewUser() *User {
	return &User{}
}
