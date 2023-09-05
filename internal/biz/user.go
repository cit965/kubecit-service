package biz

import "github.com/go-kratos/kratos/v2/log"

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo CategoryRepo
	log  *log.Helper
}

// NewCourseUsecase new a Category usecase.
func NewUserUsecase(logger log.Logger) *UserUsecase {
	return &UserUsecase{log: log.NewHelper(logger)}
}
