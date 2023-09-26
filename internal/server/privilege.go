package server

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	pb "kubecit-service/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/middleware"
)

type UserRole string

func Privilege() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tmp := ctx.Value("role_id"); tmp != nil {
				roleID := int(tmp.(uint8))
				log.Debugf("current role_id: %v\n", roleID)
				if tr, ok := transport.FromServerContext(ctx); ok {
					if _, exist := RolePrivileges[roleID][tr.Operation()]; exist {
						return handler(ctx, req)
					}
				}
			}
			return "not enough privileges", errors.New("not enough privileges")
		}
	}
}

var RolePrivileges map[int]map[string]struct{}

func init() {
	RolePrivileges = make(map[int]map[string]struct{})
	// Role Guest
	RolePrivileges[1] = map[string]struct{}{
		pb.OperationKubecitSystemSettings:        {},
		pb.OperationKubecitSearchCourse:          {},
		pb.OperationKubecitRegisterUsername:      {},
		pb.OperationKubecitMostNew:               {},
		pb.OperationKubecitLoginByJson:           {},
		pb.OperationKubecitGetSlider:             {},
		pb.OperationKubecitGetTeacher:            {},
		pb.OperationKubecitListAllTeacher:        {},
		pb.OperationKubecitListSlidersByPriority: {},
		pb.OperationKubecitListCategory:          {},
		pb.OperationKubecitListCategoryV2:        {},
		pb.OperationKubecitListLessons:           {},
		pb.OperationKubecitGetCourse:             {},
	}
	// Role RegisterUser
	RolePrivileges[2] = map[string]struct{}{
		pb.OperationKubecitSystemSettings:   {},
		pb.OperationKubecitSearchCourse:     {},
		pb.OperationKubecitRegisterUsername: {},
		pb.OperationKubecitMostNew:          {},
		pb.OperationKubecitLoginByJson:      {},
		pb.OperationKubecitGetSlider:        {},
		pb.OperationKubecitGetTeacher:       {},
		pb.OperationKubecitListAllTeacher:   {},
		pb.OperationKubecitListCategory:     {},
		pb.OperationKubecitListCategoryV2:   {},
		pb.OperationKubecitListLessons:      {},
		pb.OperationKubecitGetCourse:        {},
		pb.OperationKubecitGetInfo:          {},
		pb.OperationKubecitMyOrderList:      {},
		pb.OperationKubecitCreateOrder:      {},
		pb.OperationKubecitWalletBalance:    {},
	}
	// Role Lecturer
	RolePrivileges[3] = map[string]struct{}{
		pb.OperationKubecitSystemSettings:   {},
		pb.OperationKubecitSearchCourse:     {},
		pb.OperationKubecitRegisterUsername: {},
		pb.OperationKubecitMostNew:          {},
		pb.OperationKubecitLoginByJson:      {},
		pb.OperationKubecitGetSlider:        {},
		pb.OperationKubecitGetTeacher:       {},
		pb.OperationKubecitListAllTeacher:   {},
		pb.OperationKubecitListCategory:     {},
		pb.OperationKubecitListCategoryV2:   {},
		pb.OperationKubecitListLessons:      {},
		pb.OperationKubecitGetCourse:        {},
		pb.OperationKubecitGetInfo:          {},
		pb.OperationKubecitMyOrderList:      {},
		pb.OperationKubecitCreateOrder:      {},
		pb.OperationKubecitWalletBalance:    {},
	}
	// Role Admin
	RolePrivileges[4] = map[string]struct{}{
		pb.OperationKubecitCreateCategory:        {},
		pb.OperationKubecitCreateChapter:         {},
		pb.OperationKubecitCreateCourse:          {},
		pb.OperationKubecitCreateLesson:          {},
		pb.OperationKubecitCreateOrder:           {},
		pb.OperationKubecitCreateSlider:          {},
		pb.OperationKubecitCreateTeacher:         {},
		pb.OperationKubecitDeleteCategory:        {},
		pb.OperationKubecitDeleteChapter:         {},
		pb.OperationKubecitDeleteCourse:          {},
		pb.OperationKubecitDeleteLesson:          {},
		pb.OperationKubecitDeleteSlider:          {},
		pb.OperationKubecitGetCourse:             {},
		pb.OperationKubecitGetInfo:               {},
		pb.OperationKubecitGetSlider:             {},
		pb.OperationKubecitGetTeacher:            {},
		pb.OperationKubecitListAllTeacher:        {},
		pb.OperationKubecitListCategory:          {},
		pb.OperationKubecitListCategoryV2:        {},
		pb.OperationKubecitListLessons:           {},
		pb.OperationKubecitListSlidersByPriority: {},
		pb.OperationKubecitLoginByJson:           {},
		pb.OperationKubecitMostNew:               {},
		pb.OperationKubecitMyOrderList:           {},
		pb.OperationKubecitRechargeWallet:        {},
		pb.OperationKubecitRegisterUsername:      {},
		pb.OperationKubecitReviewCourse:          {},
		pb.OperationKubecitSearchCourse:          {},
		pb.OperationKubecitSystemSettings:        {},
		pb.OperationKubecitUpdateCategory:        {},
		pb.OperationKubecitUpdateChapter:         {},
		pb.OperationKubecitUpdateCourse:          {},
		pb.OperationKubecitUpdateLesson:          {},
		pb.OperationKubecitUpdateSlider:          {},
		pb.OperationKubecitWalletBalance:         {},
	}

}
