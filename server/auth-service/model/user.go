package model

import "gorm.io/gorm"

type UserDTO struct {
	Id       int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"userpassword"`
}

type User struct {
	gorm.Model
	Username     string `gorm:"unique"`
	Userpassword []byte
}
