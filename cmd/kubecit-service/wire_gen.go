// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"kubecit-service/internal/biz"
	"kubecit-service/internal/conf"
	"kubecit-service/internal/data"
	"kubecit-service/internal/server"
	"kubecit-service/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, gin *conf.Gin, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	categoryRepo := data.NewCategoryRepo(dataData, logger)
	courseRepo := data.NewCourseRepo(dataData, logger)
	courseUsecase := biz.NewCourseUsecase(categoryRepo, courseRepo, logger)
	sliderRepo := data.NewSliderRepo(dataData, logger)
	systemUsecase := biz.NewSystemUsecase(sliderRepo, logger)
	accountRepo := data.NewAccountRepo(dataData, logger)
	userRepo := data.NewUserRepo(dataData, logger)
	userUsecase := biz.NewUserUsecase(accountRepo, userRepo, logger)
	orderRepo := data.NewOrderRepo(dataData, logger)
	orderUseCase := biz.NewOrderUseCase(orderRepo, logger)
	teacherRepo := data.NewTeacherRepo(dataData, logger)
	teacherCase := biz.NewTeacherCase(teacherRepo, logger)
	walletRepo := data.NewWalletRepo(dataData, logger)
	walletUseCase := biz.NewWalletUseCase(walletRepo, logger)
	kubecitService := service.NewKubecitService(courseUsecase, systemUsecase, userUsecase, orderUseCase, teacherCase, walletUseCase)
	grpcServer := server.NewGRPCServer(confServer, kubecitService, logger)
	httpServer := server.NewHTTPServer(confServer, gin, confData, kubecitService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
