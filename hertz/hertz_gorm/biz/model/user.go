package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `json:"name" column:"name"`
	Gender    int64  `json:"gender" column:"gender"`
	Age       int64  `json:"age" column:"age"`
	Introduce string `json:"introduce" column:"introduce"`
}

func (u *User) TableName() string {
	return "users"
}
