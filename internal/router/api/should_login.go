package api

import (
	"github.com/gin-gonic/gin"
	"horse/app/controller/api"
)

func ShouldLoginRouter(Router *gin.RouterGroup) {
	{
		indexRouter := Router.Group("user")
		_ = indexRouter
	}
	{
		matchRouter := Router.Group("match")
		matchRouter.POST("game/list", api.NewMatchApi().GameList)
		matchRouter.POST("game/listFilter", api.NewMatchApi().GameListFilter)
		matchRouter.GET("game/outputList", api.NewMatchApi().OutputGameList)
		matchRouter.GET("game/checkStatus", api.NewMatchApi().CheckGameStatus)
		matchRouter.POST("game/resultList", api.NewMatchApi().GetResultList)
		matchRouter.GET("game/resultFilter", api.NewMatchApi().GetResultFilter)
		matchRouter.GET("game/detail", api.NewMatchApi().GameDetail)
	}
}
