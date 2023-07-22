package api

import (
	"github.com/gin-gonic/gin"
	"layout/internal/handler"
)

func VisitorRouter(Router *gin.RouterGroup, userHandler handler.UserHandler) {
	{
		indexRouter := Router.Group("")
		_ = indexRouter
	}
}
