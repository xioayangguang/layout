package h5

import (
	"github.com/gin-gonic/gin"
	"horse/app/controller/api"
	"horse/app/controller/h5"
)

func MustLoginRouter(Router *gin.RouterGroup) {

	{
		indexRouter := Router.Group("user")
		indexRouter.POST("edit-user", h5.NewUserInfo().EditUserInfo)
		indexRouter.GET("info", h5.NewUserInfo().GetUserInfo)
	}

	{
		horseRouter := Router.Group("horse")
		horseRouter.GET("detail", h5.NewHorseApi().GetHorseDetail)
		horseRouter.POST("update", h5.NewHorseApi().ChangeName)
		horseRouter.POST("list", h5.NewHorseApi().ListHorse)
		horseRouter.GET("listFilter", h5.NewHorseApi().ListHorseFilter)
		horseRouter.GET("unratedList", h5.NewHorseApi().ListUnratedHorse)
		horseRouter.GET("select", h5.NewHorseApi().GetSelectHorseList)
		horseRouter.GET("checkHorseStatus", h5.NewHorseApi().CheckHorseStatus)
	}

	{
		matchRouter := Router.Group("match")
		matchRouter.GET("list/playing", h5.NewMatchApi().ListPlaying)
		matchRouter.GET("game/sign", h5.NewMatchApi().Sign)
		matchRouter.GET("game/signature", api.NewMatchApi().Signature)
		matchRouter.POST("game/bonusList", h5.NewMatchApi().BonusList)
	}

	{
		matchRouter := Router.Group("box")
		matchRouter.GET("list", h5.NewBox().BoxList)
	}
}
