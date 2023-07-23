package api

import (
	"github.com/gin-gonic/gin"
	"layout/internal/handler"
)

func ShouldLoginRouter(Router *gin.RouterGroup, router *handler.Router) {
	{
		indexRouter := Router.Group("user")
		_ = indexRouter
	}
}
