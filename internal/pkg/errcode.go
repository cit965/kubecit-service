package pkg

import (
	"github.com/go-kratos/kratos/v2/errors"
	v1 "kubecit-service/api/helloworld/v1"
)

const (
	ErrRespCode = 450
)

var (
	// 在proto中自定义的错误码
	// BadRequest这个方法固定返回 http的400 Code！
	// Notice 但是这些code都http错误码的含义，如果业务中有更多错误情况的话，建议自定义一个专门返回业务错误的错误码！
	ErrUserNotFound = errors.BadRequest(v1.ErrorReason_UserModel_UserNotFound.String(), "没找到用户!")
	ErrUserHasBaned = errors.BadRequest(v1.ErrorReason_UserModel_UserHasBeenBaned.String(), "用户被封禁了!")

	// Notice 如果想要用自己的错误码（0 < code <= 600 之间, 超出范围将抛出异常），可以这样写～
	// 这样，客户端只要收到了450错误（正常返回200错误），就可以处理业务逻辑的错误了！
	ErrUserInBlankList = errors.New(ErrRespCode, v1.ErrorReason_UserModel_UserInBlankList.String(), "用户在黑名单中了!")
)
