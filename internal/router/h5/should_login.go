package h5

import (
	"github.com/gin-gonic/gin"
	"horse/app/controller/h5"
)

func ShouldLoginRouter(Router *gin.RouterGroup) {
	{
		matchRouter := Router.Group("match")
		matchRouter.POST("game/list", h5.NewMatchApi().GameList)
		matchRouter.POST("game/listFilter", h5.NewMatchApi().GameListFilter)
		matchRouter.GET("game/outputList", h5.NewMatchApi().OutputGameList)
		matchRouter.GET("game/checkStatus", h5.NewMatchApi().CheckGameStatus)
		matchRouter.POST("game/resultList", h5.NewMatchApi().GetResultList)
		matchRouter.GET("game/resultFilter", h5.NewMatchApi().GetResultFilter)
		matchRouter.GET("game/detail", h5.NewMatchApi().GameDetail)
	}
}
