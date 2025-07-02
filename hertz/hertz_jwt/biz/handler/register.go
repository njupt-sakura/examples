package handler

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/njupt-sakura/hertz/hertz_jwt/biz/dal/mysql"
	"github.com/njupt-sakura/hertz/hertz_jwt/biz/model"

	bizUtils "github.com/njupt-sakura/hertz/hertz_jwt/biz/utils"
)

func Register(ctx context.Context, c *app.RequestContext) {
	var registerStruct struct {
		Username string `form:"username" json:"username" query:"username" vd:"(len($) > 0 && len($) < 128); msg:'Illegal username format'"`
		Email    string `form:"email" json:"email" query:"email" vd:"(len($) > 0 && len($) < 128) && email($); msg:'Illegal email format'"`
		Password string `form:"password" json:"password" query:"password" vd:"(len($) > 0 && len($) < 128); msg: 'Illegal password format'"`
	}

	if err := c.BindAndValidate(&registerStruct); err != nil {
		c.JSON(http.StatusOK, utils.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
		return
	}

	users, err := mysql.FindUserByNameOrEmail(registerStruct.Username, registerStruct.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
		return
	}

	if len(users) != 0 {
		c.JSON(http.StatusOK, utils.H{
			"message": "user already exists",
			"code":    http.StatusBadRequest,
		})
		return
	}

	if err = mysql.CreateUsers([]*model.User{
		{
			Username: registerStruct.Username,
			Email:    registerStruct.Email,
			Password: bizUtils.MD5(registerStruct.Password),
		},
	}); err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"code":    http.StatusOK,
	})
}
