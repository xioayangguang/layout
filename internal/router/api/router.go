package api

import (
	"github.com/gin-gonic/gin"
	"layout/internal/handler"
	"layout/internal/middleware"
)

const SignSalt = "T^N5kJDOJ7seK3Z$"

func InitApiRouter(Router *gin.Engine, userHandler handler.UserHandler) {
	ApiRouter := Router.Group("api")
	ApiRouter.Use(middleware.RequestLog())
	ApiRouter.Use(middleware.Sign(SignSalt))
	ApiRouter.Use(middleware.SpeedLimit())
	ApiRouter.Use(middleware.Recover())
	//必须登录的路由
	PrivateApiGroup := ApiRouter.Group("")
	PrivateApiGroup.Use(middleware.MustTokenAuth())
	PrivateApiGroup.Use(middleware.AccessRecords())
	MustLoginRouter(PrivateApiGroup, userHandler)

	//可以登录也可以不登录的路由
	ShouldLoginApiGroup := ApiRouter.Group("")
	ShouldLoginApiGroup.Use(middleware.ShouldTokenAuth())
	ShouldLoginApiGroup.Use(middleware.AccessRecords())
	ShouldLoginRouter(ShouldLoginApiGroup, userHandler)

	//可以不登陆的路由
	PublicApiGroup := ApiRouter.Group("")
	VisitorRouter(PublicApiGroup, userHandler)
}
