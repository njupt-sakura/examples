package mysql

import (
	"github.com/njupt-sakura/hertz/tiktok_demo/pkg/constants"
	"github.com/njupt-sakura/hertz/tiktok_demo/pkg/errno"
)

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
		err := errno.Err_UserNotExist
		return nil, err
	}

	return &user, nil
}

func VerifyUser(username, password string) (int64, error) {
	var user User
	if err := DB.Where("username = ? and password = ?", username, password).Find(&user).Error; err != nil {
		return 0, err
	}

	if user.ID == 0 {
		err := errno.Err_NameOrPwdNotVerified
		return user.ID, err
	}

	return user.ID, nil
}

func CheckUserExistById(userId int64) (bool, error) {
	var user User
	if err := DB.Where("id = ?", userId).Find(&user).Error; err != nil {
		return false, err
	}
	if user == (User{}) {
		return false, nil
	}
	return true, nil
}
