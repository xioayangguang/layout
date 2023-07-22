package api

import (
	"github.com/gin-gonic/gin"
	"layout/internal/handler"
)

func MustLoginRouter(Router *gin.RouterGroup, userHandler handler.UserHandler) {
	//用户信息
	{
		indexRouter := Router.Group("user")
		_ = indexRouter
	}
}
