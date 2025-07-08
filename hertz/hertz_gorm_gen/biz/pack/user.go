package pack

import (
	"github.com/njupt-sakura/hertz/hertz_gorm_gen/biz/model/hertz/user"
	"github.com/njupt-sakura/hertz/hertz_gorm_gen/biz/model/model"
)

// Users Convert model.User list to user.User list
func Users(models []*model.User) []*user.User {
	users := make([]*user.User, 0, len(models))
	for _, m := range models {
		if u := User(m); u != nil {
			users = append(users, u)
		}
	}
	return users
}

// User Convert model.User to user.User
func User(model *model.User) *user.User {
	if model == nil {
		return nil
	}
	return &user.User{
		UserId:    model.ID,
		Name:      model.Name,
		Gender:    user.Gender(model.Gender),
		Age:       int64(model.Age),
		Introduce: model.Introduce,
	}
}
