package dal

import "github.com/njupt-sakura/hertz/hertz_gorm/biz/dal/mysql"

func Init() {
	mysql.Init()
}
