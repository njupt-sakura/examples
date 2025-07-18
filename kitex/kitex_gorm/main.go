package main

import (
	"log"

	"github.com/njupt-sakura/kitex/kitex_gorm/dal"
	user_gorm "github.com/njupt-sakura/kitex/kitex_gorm/kitex_gen/user_gorm/userservice"
)

func main() {
	dal.Init()
	svr := user_gorm.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
