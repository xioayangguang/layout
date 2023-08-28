package h5

import (
	"github.com/gin-gonic/gin"
	"layout/internal/handler/h5"
	"layout/internal/middleware"
	"time"
)

const SignSalt = "bWAOoXvIqxeiqk6*"

const Timeout = 500 * time.Millisecond

func InitH5Router(Router *gin.Engine, router *h5.Router) {
	H5Router := Router.Group("h5")
	H5Router.Use(middleware.RequestLog())
	H5Router.Use(middleware.Timeout(Timeout))
	H5Router.Use(middleware.CORSMiddleware())
	H5Router.Use(middleware.Sign(SignSalt))
	H5Router.Use(middleware.SpeedLimit())
	H5Router.Use(middleware.Recover())
	//必须登录的路由
	PrivateApiGroup := H5Router.Group("")
	PrivateApiGroup.Use(middleware.MustTokenAuth())
	PrivateApiGroup.Use(middleware.AccessRecords())
	MustLoginRouter(PrivateApiGroup, router)
	//可以登录也可以不登录的路由
	ShouldLoginApiGroup := H5Router.Group("")
	ShouldLoginApiGroup.Use(middleware.ShouldTokenAuth())
	ShouldLoginApiGroup.Use(middleware.AccessRecords())
	ShouldLoginRouter(ShouldLoginApiGroup, router)
	//可以不登陆的路由
	PublicApiGroup := H5Router.Group("")
	VisitorRouter(PublicApiGroup, router)
}
