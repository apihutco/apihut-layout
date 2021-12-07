// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"apihut-layout/internal/biz"
	"apihut-layout/internal/conf"
	"apihut-layout/internal/data"
	"apihut-layout/internal/server"
	"apihut-layout/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	client := data.NewDB(confData, logger)
	dataData, cleanup, err := data.NewData(client, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUseCase := biz.NewGreeterUseCase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, greeterService, logger)
	grpcServer := server.NewGRPCServer(confServer, greeterService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}