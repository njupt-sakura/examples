package main

import (
	"context"

	user_gorm "github.com/njupt-sakura/kitex/kitex_gorm_gen/kitex_gen/user_gorm"
	"github.com/njupt-sakura/kitex/kitex_gorm_gen/model/model"
	"github.com/njupt-sakura/kitex/kitex_gorm_gen/model/query"
	"github.com/njupt-sakura/kitex/kitex_gorm_gen/pack"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *user_gorm.UpdateUserRequest) (resp *user_gorm.UpdateUserResponse, err error) {
	resp = new(user_gorm.UpdateUserResponse)
	u := &model.User{
		Name:      req.Name,
		Gender:    int32(req.Gender),
		Age:       int32(req.Age),
		Introduce: req.Introduce,
	}

	if _, err = query.User.WithContext(ctx).Updates(u); err != nil {
		resp.Code = user_gorm.Code_DBErr
		resp.Msg = err.Error()
		return
	}

	resp.Code = user_gorm.Code_Success
	return
}

// DeleteUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteUser(ctx context.Context, req *user_gorm.DeleteUserRequest) (resp *user_gorm.DeleteUserResponse, err error) {
	resp = new(user_gorm.DeleteUserResponse)
	if _, err = query.User.WithContext(ctx).Where(query.User.ID.Eq((int64(req.UserId)))).Delete(); err != nil {
		resp.Code = user_gorm.Code_DBErr
		resp.Msg = err.Error()
		return
	}

	resp.Code = user_gorm.Code_Success
	return
}

// QueryUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUser(ctx context.Context, req *user_gorm.QueryUserRequest) (resp *user_gorm.QueryUserResponse, err error) {
	resp = new(user_gorm.QueryUserResponse)
	u, exec := query.User, query.User.WithContext(ctx)
	if *req.Keyword != "" {
		exec = exec.Where(u.Introduce.Like("%" + *req.Keyword + "%"))
	}

	var total int64
	total, err = exec.Count()
	if err != nil {
		resp.Code = user_gorm.Code_DBErr
		resp.Msg = err.Error()
		return
	}

	var users []*model.User
	if total > 0 {
		users, err = exec.Limit(int(req.PageSize)).Offset(int(req.PageSize * (req.Page - 1))).Find()
		if err != nil {
			resp.Code = user_gorm.Code_DBErr
			resp.Msg = err.Error()
			return
		}
	}

	resp.Code = user_gorm.Code_Success
	resp.Total = total
	resp.Users = pack.Users(users)

	return
}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user_gorm.CreateUserRequest) (resp *user_gorm.CreateUserResponse, err error) {
	resp = new(user_gorm.CreateUserResponse)

	err = query.User.WithContext(ctx).Create(&model.User{
		Name:      req.Name,
		Gender:    int32(req.Gender),
		Age:       int32(req.Age),
		Introduce: req.Introduce,
	})

	if err != nil {
		resp.Code = user_gorm.Code_DBErr
		resp.Msg = err.Error()
		return
	}

	resp.Code = user_gorm.Code_Success
	return
}
