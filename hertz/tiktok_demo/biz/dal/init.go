package dal

import (
	"github.com/njupt-sakura/hertz/tiktok_demo/biz/dal/mysql"
	"github.com/njupt-sakura/hertz/tiktok_demo/biz/dal/redis"
)

func Init() {
	mysql.Init() // mysql
	redis.Init()
}
