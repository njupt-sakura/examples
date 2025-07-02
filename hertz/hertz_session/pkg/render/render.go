package render

import (
	"context"
	"net/http"
	"text/template"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	hertzUtils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/csrf"
	"github.com/hertz-contrib/sessions"
	"github.com/njupt-sakura/hertz/hertz_session/pkg/consts"
	"github.com/njupt-sakura/hertz/hertz_session/pkg/utils"
)

func InitHtml(h *server.Hertz) {
	h.Delims("{[{", "}]}")
	h.SetFuncMap(template.FuncMap{
		"BuildMsg": utils.BuildMsg,
	})

	// load templates
	h.LoadHTMLGlob("static/html/*")
	h.Static("/", "./static")
	token := ""
	h.GET("/register.html", func(ctx context.Context, c *app.RequestContext) {
		if !utils.IsLogout(ctx, c) {
			token = csrf.GetToken(c)
		}
		c.HTML(http.StatusOK, "register.html", hertzUtils.H{
			"message": utils.BuildMsg("Register a new membership"),
			"token":   utils.BuildMsg(token),
		})
	})

	// login.html
	h.GET("/login.html", func(ctx context.Context, c *app.RequestContext) {
		if !utils.IsLogout(ctx, c) {
			token = csrf.GetToken(c)
		}
		c.HTML(http.StatusOK, "login.html", hertzUtils.H{
			"message": utils.BuildMsg("Login to start your session"),
			"token":   utils.BuildMsg(token),
		})
	})

	// index.html
	h.GET("/index.html", func(ctx context.Context, c *app.RequestContext) {
		if !utils.IsLogout(ctx, c) {
			token = csrf.GetToken(c)
		}
		session := sessions.Default(c)
		username := session.Get(consts.Username)
		if username == nil {
			c.HTML(http.StatusOK, "index.html", hertzUtils.H{
				"message": utils.BuildMsg(consts.PageErr),
				"token":   utils.BuildMsg(token),
			})
			c.Redirect(http.StatusMovedPermanently, []byte("/login.html"))
			return
		}

		c.HTML(http.StatusOK, "index.html", hertzUtils.H{
			"message": utils.BuildMsg(username.(string)),
			"token":   utils.BuildMsg(token),
		})
	})
}
