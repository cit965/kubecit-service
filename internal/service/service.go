package service

import (
	"github.com/google/wire"
	v1 "kubecit-service/api/helloworld/v1"
	"kubecit-service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewKubecitService)

// KubecitService is a greeter service.
type KubecitService struct {
	v1.UnimplementedKubecitServer

	cc          *biz.CourseUsecase
	su          *biz.SystemUsecase
	userUseCase *biz.UserUsecase
	orderCase   *biz.OrderUseCase
	walletCase  *biz.WalletUseCase
	teacherCase *biz.TeacherCase
}

// NewGreeterService new a greeter service.
func NewKubecitService(cc *biz.CourseUsecase, su *biz.SystemUsecase, userUseCase *biz.UserUsecase, orderCase *biz.OrderUseCase, teacherCase *biz.TeacherCase,walletCase *biz.WalletUseCase) *KubecitService {
	return &KubecitService{cc: cc, su: su, userUseCase: userUseCase, orderCase: orderCase, teacherCase: teacherCase,walletCase: walletCase}
}
