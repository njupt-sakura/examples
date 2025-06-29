package dal

import (
	"github.com/njupt-sakura/hertz/hertz_gorm_gen/biz/dal/mysql"
	"github.com/njupt-sakura/hertz/hertz_gorm_gen/biz/model/query"
)

func Init() {
	mysql.Init()
	query.SetDefault(mysql.DB)
}
