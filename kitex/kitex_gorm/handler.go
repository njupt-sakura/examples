package main

import (
	"context"

	"github.com/njupt-sakura/kitex/kitex_gorm/dal/mysql"
	user_gorm "github.com/njupt-sakura/kitex/kitex_gorm/kitex_gen/user_gorm"
	"github.com/njupt-sakura/kitex/kitex_gorm/model"
	"github.com/njupt-sakura/kitex/kitex_gorm/pack"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *user_gorm.UpdateUserRequest) (resp *user_gorm.UpdateUserResponse, err error) {
	resp = new(user_gorm.UpdateUserResponse)

	u := &model.User{
		Name:      req.Name,
		Gender:    int64(req.Gender),
		Age:       req.Age,
		Introduce: req.Introduce,
	}

	if err = mysql.UpdateUser(u); err != nil {
		resp.Msg = err.Error()
		resp.Code = user_gorm.Code_DBErr
		return
	}

	resp.Code = user_gorm.Code_Success
	return
}

// DeleteUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteUser(ctx context.Context, req *user_gorm.DeleteUserRequest) (resp *user_gorm.DeleteUserResponse, err error) {
	resp = new(user_gorm.DeleteUserResponse)

	if err = mysql.DeleteUser(req.UserId); err != nil {
		resp.Msg = err.Error()
		resp.Code = user_gorm.Code_DBErr
		return
	}

	resp.Code = user_gorm.Code_Success
	return
}

// QueryUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUser(ctx context.Context, req *user_gorm.QueryUserRequest) (resp *user_gorm.QueryUserResponse, err error) {
	resp = new(user_gorm.QueryUserResponse)

	users, total, err := mysql.QueryUser(req.Keyword, req.Page, req.PageSize)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = user_gorm.Code_DBErr
		return
	}

	resp.Total = total
	resp.Users = pack.Users(users)
	resp.Code = user_gorm.Code_Success

	return
}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user_gorm.CreateUserRequest) (resp *user_gorm.CreateUserResponse, err error) {
	resp = new(user_gorm.CreateUserResponse)

	if err = mysql.CreateUser([]*model.User{{
		Name:      req.Name,
		Gender:    int64(req.Gender),
		Age:       req.Age,
		Introduce: req.Introduce,
	}}); err != nil {
		resp.Msg = err.Error()
		resp.Code = user_gorm.Code_DBErr
		return
	}

	resp.Code = user_gorm.Code_Success
	return
}
