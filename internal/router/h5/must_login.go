package h5

import (
	"github.com/gin-gonic/gin"
	"layout/internal/handler"
)

func MustLoginRouter(Router *gin.RouterGroup, userHandler handler.UserHandler) {
	{
		indexRouter := Router.Group("user")
		_ = indexRouter
	}
}
