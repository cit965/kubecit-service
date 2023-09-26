package server

import (
	"context"

	"kubecit-service/internal/pkg/jwt"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

func Auth() middleware.Middleware {

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if header, ok := transport.FromServerContext(ctx); ok {
				token := header.RequestHeader().Get("Authorization")
				if token != "" {
					if claims, err := jwt.VerifyToken(token); err == nil {
						ctx = context.WithValue(ctx, "user_id", claims.UserID)
						ctx = context.WithValue(ctx, "role_id", claims.RoleID)
					} else {
						return "token is invalid", err
					}
				} else {
					// 如果token为空则设置角色id为uint8(1),表示未登录访客
					ctx = context.WithValue(ctx, "role_id", uint8(1))
				}
			}
			return handler(ctx, req)
		}
	}
}
