package api

import (
	"github.com/gin-gonic/gin"
	"layout/internal/handler"
)

func ShouldLoginRouter(Router *gin.RouterGroup, userHandler handler.UserHandler) {
	{
		indexRouter := Router.Group("user")
		_ = indexRouter
	}
}
