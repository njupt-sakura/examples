package mysql

import "github.com/njupt-sakura/hertz/tiktok_demo/pkg/constants"

type User struct {
	ID              int64  `json:"id"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
}

func (User) TableName() string {
	return constants.UsersTableName
}

func CreateUser(user *User) (int64, error) {
	err := DB.Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, err
}

func QueryUserById(userId int64) (*User, error) {
	var user User
	if err := DB.Where("id = ?", userId).Find(&user).Error; err != nil {
		return nil, err
	}

	if user == (User{}) {
	}
}
