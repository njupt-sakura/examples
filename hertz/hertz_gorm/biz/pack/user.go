package pack

import (
	"github.com/njupt-sakura/hertz/hertz_gorm/biz/hertz_gen/user_gorm"
	"github.com/njupt-sakura/hertz/hertz_gorm/biz/model"
)

func Users(models []*model.User) []*user_gorm.User {
	users := make([]*user_gorm.User, 0, len(models) /* cap */)

	for _, m := range models {
		if u := User(m); u != nil {
			users = append(users, u)
		}
	}
	return users
}

func User(model *model.User) *user_gorm.User {
	if model == nil {
		return nil
	}

	return &user_gorm.User{
		UserID:    int64(model.ID),
		Name:      model.Name,
		Gender:    user_gorm.Gender(model.Gender),
		Age:       model.Age,
		Introduce: model.Introduce,
	}
}
