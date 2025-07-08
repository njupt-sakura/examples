package utils

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/njupt-sakura/hertz/hertz_session/pkg/consts"
)

func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func BuildMsg(msg string) string {
	return fmt.Sprintf("%v", msg)
}

func IsLogout(_ context.Context, c *app.RequestContext) bool {
	return string(c.Cookie(consts.HertzSession)) == ""
}
