package api

import (
	"github.com/gin-gonic/gin"
	"layout/internal/handler"
)

func MustLoginRouter(Router *gin.RouterGroup, router *handler.Router) {
	//用户信息
	{
		indexRouter := Router.Group("user")
		_ = indexRouter
	}
}
