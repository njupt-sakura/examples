package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../model/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	dsn := "whoami:pass@(127.0.0.1:3307)/db0?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	g.UseDB(db)

	// Generate struct `User` based on table `users`
	genModel := g.GenerateModel("users")

	// Generate basic type-safe DAO API for generated struct `model.User` following conventions
	g.ApplyBasic(genModel)

	// Generate the code
	g.Execute()
}
