package user

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzUtils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/sessions"
	"github.com/njupt-sakura/hertz/hertz_session/biz/dal/mysql"
	user "github.com/njupt-sakura/hertz/hertz_session/biz/model/user"
	"github.com/njupt-sakura/hertz/hertz_session/pkg/consts"
	"github.com/njupt-sakura/hertz/hertz_session/pkg/utils"
)

func toRegister(c *app.RequestContext, message, token string) {
	c.HTML(http.StatusOK, "register.html", hertzUtils.H{
		"message": utils.BuildMsg(message),
		"token":   utils.BuildMsg(token),
	})
}

func toLogin(c *app.RequestContext, message, token string) {
	c.HTML(http.StatusOK, "login.html", hertzUtils.H{
		"message": utils.BuildMsg(message),
		"token":   utils.BuildMsg(token),
	})
}

// Register .
// @router /register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.RegisterRequest
	token := ""

	err = c.BindAndValidate(&req)
	if err != nil {
		toRegister(c, err.Error(), token)
		return
	}

	users, err := mysql.FindUserByNameOrEmail(req.Username, req.Email)
	if err != nil {
		toRegister(c, err.Error(), token)
		return
	}
	if len(users) != 0 {
		toRegister(c, consts.RegisterErr, token)
		return
	}

	if err = mysql.CreateUsers([]*mysql.User{
		{
			Username: req.Username,
			Password: utils.MD5(req.Password),
			Email:    req.Email,
		},
	}); err != nil {
		toRegister(c, err.Error(), token)
		return
	}

	// toLogin(c, consts.Success, token)
	c.Redirect(http.StatusMovedPermanently, []byte("/login.html"))
}

// Login .
// @router /login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.LoginRequest
	token := ""

	err = c.BindAndValidate(&req)
	if err != nil {
		toLogin(c, err.Error(), token)
		return
	}

	users, err := mysql.CheckUser(req.Username, utils.MD5(req.Password))
	if err != nil {
		toLogin(c, err.Error(), token)
		return
	}

	if len(users) == 0 {
		toLogin(c, consts.LoginErr, token)
		return
	}

	session := sessions.Default(c)
	session.Set(consts.Username, req.Username)
	err = session.Save()
	if err != nil {
		hlog.CtxErrorf(ctx, "Save session error")
	}
	c.Redirect(http.StatusMovedPermanently, []byte("/index.html"))
}

func Logout(ctx context.Context, c *app.RequestContext) {
	session := sessions.Default(c)
	session.Delete(consts.Username)
	err := session.Save()
	if err != nil {
		hlog.CtxErrorf(ctx, "Save session error")
	}
	c.Redirect(http.StatusMovedPermanently, []byte("/login.html"))
}
