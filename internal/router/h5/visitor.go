package h5

import (
	"github.com/gin-gonic/gin"
	"horse/app/controller/api"
	"horse/app/controller/h5"
)

func VisitorRouter(Router *gin.RouterGroup) {
	{
		indexRouter := Router.Group("")
		indexRouter.POST("user/wallet-login", h5.NewUserInfo().WalletLogin)
		indexRouter.GET("user/login-msg", h5.NewUserInfo().WalletLoginNonce)
	}
	{
		matchRouter := Router.Group("match")
		matchRouter.GET("game/process", h5.NewMatchApi().Process)
		matchRouter.GET("game/end", h5.NewMatchApi().ProcessEnd)
	}
	{
		horseRouter := Router.Group("horse")
		horseRouter.GET("record", api.NewHorseApi().Record)
	}
}
