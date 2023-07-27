//go:build wireinject
// +build wireinject

package wireinject

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"layout/internal/handler"
	"layout/internal/handler/app"
	"layout/internal/handler/h5"
	"layout/internal/repository"
	"layout/internal/router"
	"layout/internal/service"
	_ "layout/pkg/monitor"
)

var HandlerSet = wire.NewSet(handler.ProviderSet,
	app.ProviderSet, app.StructProvider,
	h5.ProviderSet, h5.StructProvider,
)

func NewApp() (*gin.Engine, func(), error) {
	panic(wire.Build(
		router.NewServerHTTP,
		repository.ProviderSet,
		service.ProviderSet,
		HandlerSet,
	))
}
