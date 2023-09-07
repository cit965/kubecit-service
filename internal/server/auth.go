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
				token := header.RequestHeader().Get("kubecit-usertoken")
				if token != "" {
					if claims, err := jwt.VerifyToken(token); err == nil {
						ctx = context.WithValue(ctx, "user_id", claims.UserID)
						ctx = context.WithValue(ctx, "role_id", claims.RoleID)
					}

				}
			}
			return handler(ctx, req)
		}
	}
}
