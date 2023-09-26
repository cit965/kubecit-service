package server

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"

	pb "kubecit-service/api/helloworld/v1"
	"kubecit-service/internal/biz"
)

func Privilege() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tmp := ctx.Value("role_id"); tmp != nil {
				roleID := tmp.(uint8)
				log.Debugf("current role_id: %v\n", roleID)
				switch roleID {
				case biz.UserRoleGuest, biz.UserRoleRegisterUser:
					if tr, ok := transport.FromServerContext(ctx); ok {
						if _, exist := RolePrivileges[roleID][tr.Operation()]; exist {
							return handler(ctx, req)
						}
					}
				case biz.UserRoleSuperAdmin:
					return handler(ctx, req)
				}

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

var RolePrivileges map[uint8]map[string]struct{}

func init() {
	RolePrivileges = make(map[uint8]map[string]struct{})
	// Role Guest
	RolePrivileges[biz.UserRoleGuest] = map[string]struct{}{
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
	RolePrivileges[biz.UserRoleRegisterUser] = map[string]struct{}{
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
	RolePrivileges[biz.UserRoleSuperAdmin] = map[string]struct{}{}
}
