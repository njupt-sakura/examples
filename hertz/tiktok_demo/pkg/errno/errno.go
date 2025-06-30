package errno

import (
	"errors"
	"fmt"
)

const (
	Code_Success = 0

	ErrCode_ServiceError = iota + 1000
	ErrCode_ParameterError
	ErrCode_AuthorizationFailed

	ErrCode_UserAlreadyExist
	ErrCode_UserNotExist

	ErrCode_FollowRelationAlreadyExist
	ErrCode_FollowRelationNotExist

	ErrCode_FavoriteRelationAlreadyExist
	ErrCode_FavoriteRelationNotExist
	ErrCode_FavoriteActionError

	ErrCode_MessageAddFailed
	ErrCode_FriendListNoPermission

	ErrCode_VideoNotExist
	ErrCode_CommentNotExist
)

const (
	Msg_Success = "Success"

	ErrMsg_ServerError         = "Server error"
	ErrMsg_ParameterError      = "Parameter error"
	ErrMsg_AuthorizationFailed = "Authorization failed"

	ErrMsg_UserAlreadyExist = "User already exist"
	ErrMsg_UserNotExist     = "User is not exist"

	ErrMsg_FollowRelationAlreadyExist = "Follow relation already exist"
	ErrMsg_FollowRelationNotExist     = "Follow relation not exist"

	ErrMsg_FavoriteRelationAlreadyExist = "Favorite relation already exist"
	ErrMsg_FavoriteRelationNotExist     = "Favorite relation not exist"
	ErrMsg_FavoriteActionError          = "Favorite action error"

	ErrMsg_MessageAddFailed       = "Message add failed"
	ErrMsg_FriendListNoPermission = "Friend list no permission"

	ErrMsg_VideoNotExist   = "Video is not exist"
	ErrMsg_CommentNotExist = "Comment is not exist"

	ErrMsg_NameOrPwdNotVerified = "Username or password not verified"
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("ErrCode=%d, ErrMsg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e *ErrNo) WithMsg(msg string) ErrNo {
	e.ErrMsg = msg
	return *e
}

var _ error = (*ErrNo)(nil)

var (
	OK_Success = NewErrNo(Code_Success, Msg_Success)

	Err_ServiceError        = NewErrNo(ErrCode_ServiceError, ErrMsg_ServerError)
	Err_ParameterError      = NewErrNo(ErrCode_ParameterError, ErrMsg_ParameterError)
	Err_AuthorizationFailed = NewErrNo(ErrCode_AuthorizationFailed, ErrMsg_AuthorizationFailed)

	Err_UserAlreadyExist = NewErrNo(ErrCode_UserAlreadyExist, ErrMsg_UserAlreadyExist)
	Err_UserNotExist     = NewErrNo(ErrCode_UserNotExist, ErrMsg_UserNotExist)

	Err_FollowRelationAlreadyExist = NewErrNo(ErrCode_FollowRelationAlreadyExist, ErrMsg_FollowRelationAlreadyExist)
	Err_FollowRelationNotExist     = NewErrNo(ErrCode_FollowRelationNotExist, ErrMsg_FollowRelationNotExist)

	Err_FavoriteRelationAlreadyExist = NewErrNo(ErrCode_FavoriteRelationAlreadyExist, ErrMsg_FollowRelationAlreadyExist)
	Err_FavoriteRelationNotExist     = NewErrNo(ErrCode_FavoriteRelationNotExist, ErrMsg_FavoriteRelationNotExist)
	Err_FavoriteActionError          = NewErrNo(ErrCode_FavoriteActionError, ErrMsg_FavoriteActionError)

	Err_MessageAddFailed       = NewErrNo(ErrCode_MessageAddFailed, ErrMsg_MessageAddFailed)
	Err_FriendListNoPermission = NewErrNo(ErrCode_FriendListNoPermission, ErrMsg_FriendListNoPermission)

	Err_VideoNotExist   = NewErrNo(ErrCode_VideoNotExist, ErrMsg_VideoNotExist)
	Err_CommentNotExist = NewErrNo(ErrCode_CommentNotExist, ErrMsg_CommentNotExist)

	Err_NameOrPwdNotVerified = NewErrNo(ErrCode_AuthorizationFailed, ErrMsg_NameOrPwdNotVerified)
)

func ConvertErr(err error) ErrNo {
	if errors.As(err, &ErrNo{}) {
		return *(err.(*ErrNo))
	}

	errno := Err_ServiceError
	errno.ErrMsg = err.Error()
	return errno
}
