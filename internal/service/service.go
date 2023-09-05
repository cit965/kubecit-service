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
	v1.UnimplementedGreeterServer

	cc *biz.CourseUsecase
}

// NewGreeterService new a greeter service.
func NewKubecitService(cc *biz.CourseUsecase) *KubecitService {
	return &KubecitService{cc: cc}
}
