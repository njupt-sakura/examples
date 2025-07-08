package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" column:"username"`
	Email    string `json:"email" column:"email"`
	Password string `json:"password" column:"password"`
}

func (u *User) TableName() string {
	return "users"
}
