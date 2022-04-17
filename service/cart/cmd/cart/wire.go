// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"cart/internal/biz"
	"cart/internal/conf"
	"cart/internal/data"
	"cart/internal/server"
	"cart/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
