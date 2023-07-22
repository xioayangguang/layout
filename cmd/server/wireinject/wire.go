//go:build wireinject
// +build wireinject

package wireinject

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"layout/internal/handler"
	"layout/internal/repository"
	"layout/internal/router"
	"layout/internal/service"
	"layout/pkg/redis"
)

var _ = redis.InitRedis

var ServerSet = wire.NewSet(router.NewServerHTTP)

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
	repository.NewRepository,
	repository.NewUserRepository,
)

func NewApp(*viper.Viper) (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,
		RepositorySet,
		ServiceSet,
		HandlerSet,
	))
}
