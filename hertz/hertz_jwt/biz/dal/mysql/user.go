package mysql

import "github.com/njupt-sakura/hertz/hertz_jwt/biz/model"

func CreateUsers(users []*model.User) error {
	return DB.Create(users).Error
}

func FindUserByNameOrEmail(username, email string) ([]*model.User, error) {
	res := make([]*model.User, 0)
	if err := DB.Where(DB.Or("username = ?", username).Or("email = ?", email)).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func CheckUser(account, password string) ([]*model.User, error) {
	res := make([]*model.User, 0)
	if err := DB.Where(DB.Or("username = ?", account).
		Or("email = ?", account)).
		Where("password = ?", password).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
