// Code generated by hertz generator. DO NOT EDIT.

package user

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	user "github.com/njupt-sakura/hertz/hertz_gorm_gen/biz/handler/user"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_v1 := root.Group("/v1", _v1Mw()...)
		{
			_user := _v1.Group("/user", _userMw()...)
			_user.POST("/create", append(_createuserresponseMw(), user.CreateUserResponse)...)
			_user.POST("/query", append(_queryuserresponseMw(), user.QueryUserResponse)...)
			{
				_delete := _user.Group("/delete", _deleteMw()...)
				_delete.POST("/:user_id", append(_deleteuserresponseMw(), user.DeleteUserResponse)...)
			}
			{
				_update := _user.Group("/update", _updateMw()...)
				_update.POST("/:user_id", append(_updateuserresponseMw(), user.UpdateUserResponse)...)
			}
		}
	}
}
