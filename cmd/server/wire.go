//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"layout/internal/handler"
	"layout/internal/middleware"
	"layout/internal/repository"
	"layout/internal/server"
	"layout/internal/service"
	"layout/pkg/helper/sid"
	"layout/pkg/log"
)

var ServerSet = wire.NewSet(server.NewServerHTTP)

var SidSet = wire.NewSet(sid.NewSid)

var JwtSet = wire.NewSet(middleware.NewJwt)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var RepositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewUserRepository,
)

func newApp(*viper.Viper, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,
		RepositorySet,
		ServiceSet,
		HandlerSet,
		SidSet,
		JwtSet,
	))
}
