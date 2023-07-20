package api

import (
	"github.com/gin-gonic/gin"
	"horse/app/controller/api"
)

func VisitorRouter(Router *gin.RouterGroup) {
	{
		indexRouter := Router.Group("")
		indexRouter.POST("user/wallet-login", api.NewUserInfo().WalletLogin)
		indexRouter.GET("user/login-msg", api.NewUserInfo().WalletLoginNonce)
	}
	{
		matchRouter := Router.Group("match")
		matchRouter.GET("game/process", api.NewMatchApi().Process)
		matchRouter.GET("game/end", api.NewMatchApi().ProcessEnd)
	}
	{
		horseRouter := Router.Group("horse")
		horseRouter.GET("record", api.NewHorseApi().Record)
	}
}
