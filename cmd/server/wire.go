//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"layout/internal/handler"
	"layout/internal/repository"
	"layout/internal/router"
	"layout/internal/service"
	_ "layout/pkg/config"
	_ "layout/pkg/redis"
)

var ServerSet = wire.NewSet(router.NewServerHTTP)

var HandlerSet = wire.NewSet(
	handler.NewUserHandler,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var RepositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRepository,
	repository.NewUserRepository,
)

func newApp() (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,
		RepositorySet,
		ServiceSet,
		HandlerSet,
	))
}
