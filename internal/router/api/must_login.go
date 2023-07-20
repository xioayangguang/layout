package api

import (
	"github.com/gin-gonic/gin"
	"horse/app/controller/api"
)

func MustLoginRouter(Router *gin.RouterGroup) {
	//用户信息
	{
		indexRouter := Router.Group("user")
		indexRouter.POST("edit-user", api.NewUserInfo().EditUserInfo)
		indexRouter.GET("info", api.NewUserInfo().GetUserInfo)
	}
	// 马匹相关
	{
		horseRouter := Router.Group("horse")
		horseRouter.GET("detail", api.NewHorseApi().GetHorseDetail)
		horseRouter.POST("update", api.NewHorseApi().ChangeName)
		horseRouter.POST("list", api.NewHorseApi().ListHorse)
		horseRouter.GET("listFilter", api.NewHorseApi().ListHorseFilter)
		horseRouter.GET("unratedList", api.NewHorseApi().ListUnratedHorse)
		horseRouter.GET("select", api.NewHorseApi().GetSelectHorseList)
		horseRouter.GET("checkHorseStatus", api.NewHorseApi().CheckHorseStatus)
	}
	// 比赛相关
	{
		matchRouter := Router.Group("match")
		matchRouter.GET("list/playing", api.NewMatchApi().ListPlaying)
		matchRouter.GET("game/sign", api.NewMatchApi().Sign)
		matchRouter.GET("game/signature", api.NewMatchApi().Signature)
		matchRouter.POST("game/bonusList", api.NewMatchApi().BonusList)
	}
	// 盲盒
	{
		matchRouter := Router.Group("box")
		matchRouter.GET("list", api.NewBox().BoxList)
	}
}
