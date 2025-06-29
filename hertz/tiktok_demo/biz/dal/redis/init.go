package redis

import (
	"github.com/njupt-sakura/hertz/tiktok_demo/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MysqlDefaultDsn))
}
