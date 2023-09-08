package server

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/validate"

	v1 "kubecit-service/api/helloworld/v1"
	"kubecit-service/gin"
	"kubecit-service/internal/conf"
	"kubecit-service/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.KubecitService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			Auth(),
			validate.Validator(),
		),
	}

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	openAPIHandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIHandler)

	logService := gin.NewGinService()
	srv.HandlePrefix("/web/", logService)
	v1.RegisterKubecitHTTPServer(srv, greeter)
	srv.WalkRoute(func(info http.RouteInfo) error {
		fmt.Printf("%-50s \t %s\n", info.Path, info.Method)
		return nil
	})
	return srv
}
