// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/apihutco/apihut-layout/internal/biz"
	"github.com/apihutco/apihut-layout/internal/conf"
	"github.com/apihutco/apihut-layout/internal/data"
	"github.com/apihutco/apihut-layout/internal/server"
	"github.com/apihutco/apihut-layout/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
