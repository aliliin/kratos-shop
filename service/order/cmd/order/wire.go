//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"order/internal/conf"
	"order/internal/repo"
	"order/internal/server"
	"order/internal/service"
	"order/internal/usecase"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Registry, *conf.Data, *conf.Auth, *conf.Service, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, repo.ProviderSet, usecase.ProviderSet, service.ProviderSet, newApp))
}
