package router

import (
	"github.com/gin-gonic/gin"
	"layout/internal/handler/app"
	"layout/internal/handler/h5"
	"layout/internal/middleware"
)

// InitApiRouter 不登陆也不验证签名的路由，通常是一些回调路由
func InitApiRouter(Router *gin.Engine, approuter *app.Router, h5router *h5.Router) {
	PublicApiGroup := Router.Group("common")
	PublicApiGroup.Use(middleware.RequestLog())
	PublicApiGroup.Use(middleware.SpeedLimit())
	PublicApiGroup.Use(middleware.Recover())
	{
		indexRouter := PublicApiGroup.Group("horses")
		_ = indexRouter
	}
}
