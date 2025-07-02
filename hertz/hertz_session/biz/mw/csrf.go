package mw

import (
	"context"
	"errors"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/csrf"
	"github.com/njupt-sakura/hertz/hertz_session/pkg/consts"
	"github.com/njupt-sakura/hertz/hertz_session/pkg/utils"
)

func InitCsrf(h *server.Hertz) {
	h.Use(csrf.New(
		csrf.WithSecret(consts.CsrfSecretKey),
		csrf.WithKeyLookUp(consts.CsrfKeyLookUp),
		csrf.WithNext(utils.IsLogout),
		csrf.WithErrorFunc(func(ctx context.Context, c *app.RequestContext) {
			c.String(http.StatusBadRequest, errors.New(consts.CsrfErr).Error())
			c.Abort()
		}),
	))
}
