package h5

import (
	"github.com/gin-gonic/gin"
	"layout/internal/handler"
)

func MustLoginRouter(Router *gin.RouterGroup, router *handler.Router) {
	{
		indexRouter := Router.Group("user")
		_ = indexRouter
	}
}
