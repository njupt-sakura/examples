// Code generated by Kitex v0.13.1. DO NOT EDIT.
package userservice

import (
	server "github.com/cloudwego/kitex/server"
	user_gorm "github.com/njupt-sakura/kitex/kitex_gorm/kitex_gen/user_gorm"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler user_gorm.UserService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)
	options = append(options, server.WithCompatibleMiddlewareForUnary())

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler user_gorm.UserService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
