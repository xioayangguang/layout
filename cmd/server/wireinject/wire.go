//go:build wireinject
// +build wireinject

package wireinject

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"layout/internal/handler"
	"layout/internal/repository"
	"layout/internal/router"
	"layout/internal/service"
)

var ServerSet = wire.NewSet(router.NewServerHTTP, handler.StructProvider)

func NewApp() (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,
		repository.ProviderSet,
		service.ProviderSet,
		handler.ProviderSet,
	))
}
