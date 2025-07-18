package mysql

import (
	"github.com/njupt-sakura/hertz/hertz_session/pkg/consts"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(consts.MysqlDefaultDsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	m := DB.Migrator()
	if m.HasTable(&User{}) {
		return
	}

	if err = m.CreateTable(&User{}); err != nil {
		panic(err)
	}
}
