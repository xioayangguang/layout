package app

import (
	"github.com/gin-gonic/gin"
	"layout/internal/handler/app"
	"layout/internal/middleware"
	"time"
)

const SignSalt = "T^N5kJDOJ7seK3Z$"

const Timeout = 500 * time.Millisecond

func InitAppRouter(Router *gin.Engine, router *app.Router) {
	ApiRouter := Router.Group("api")
	ApiRouter.Use(middleware.RequestLog())
	ApiRouter.Use(middleware.Timeout(Timeout))
	ApiRouter.Use(middleware.Sign(SignSalt))
	//全局分布式限速
	//ApiRouter.Use(middleware.SpeedLimit())
	//单机限速（如果网关做了ip哈希的优先试用单机限速提高性能）
	ApiRouter.Use(middleware.TokenLimit())
	ApiRouter.Use(middleware.Recover())
	//必须登录的路由
	PrivateApiGroup := ApiRouter.Group("")
	PrivateApiGroup.Use(middleware.MustTokenAuth())
	PrivateApiGroup.Use(middleware.AccessRecords())
	MustLoginRouter(PrivateApiGroup, router)

	//可以登录也可以不登录的路由
	ShouldLoginApiGroup := ApiRouter.Group("")
	ShouldLoginApiGroup.Use(middleware.ShouldTokenAuth())
	ShouldLoginApiGroup.Use(middleware.AccessRecords())
	ShouldLoginRouter(ShouldLoginApiGroup, router)

	//可以不登陆的路由
	PublicApiGroup := ApiRouter.Group("")
	VisitorRouter(PublicApiGroup, router)
}
