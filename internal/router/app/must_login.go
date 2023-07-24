package app

import (
	"github.com/gin-gonic/gin"
	"layout/internal/handler/app"
)

func MustLoginRouter(Router *gin.RouterGroup, router *app.Router) {
	//用户信息
	{
		indexRouter := Router.Group("user")
		indexRouter.GET("info", router.AppUser.GetProfile)
		indexRouter.POST("update", router.AppUser.UpdateProfile)
	}
}
