package api

import (
	"github.com/gin-gonic/gin"
	"horse/middleware"
)

const SignSalt = "T^N5kJDOJ7seK3Z$"

func InitApiRouter(Router *gin.Engine) {
	ApiRouter := Router.Group("api")
	ApiRouter.Use(middleware.RequestLog())
	ApiRouter.Use(middleware.Sign(SignSalt))
	ApiRouter.Use(middleware.SpeedLimit())
	//if !global.Config.Debug {
	// debug模式下 开启控制台报错 日志错误不全
	ApiRouter.Use(middleware.Recover())
	//}

	//必须登录的路由
	PrivateApiGroup := ApiRouter.Group("")
	PrivateApiGroup.Use(middleware.MustTokenAuth())
	PrivateApiGroup.Use(middleware.AccessRecords())
	MustLoginRouter(PrivateApiGroup)

	//可以登录也可以不登录的路由
	ShouldLoginApiGroup := ApiRouter.Group("")
	ShouldLoginApiGroup.Use(middleware.ShouldTokenAuth())
	ShouldLoginApiGroup.Use(middleware.AccessRecords())
	ShouldLoginRouter(ShouldLoginApiGroup)

	//可以不登陆的路由
	PublicApiGroup := ApiRouter.Group("")
	VisitorRouter(PublicApiGroup)
}
