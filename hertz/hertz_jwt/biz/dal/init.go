package dal

import "github.com/njupt-sakura/hertz/hertz_jwt/biz/dal/mysql"

func Init() {
	mysql.Init()
}
