package api

import (
	"github.com/gin-gonic/gin"
	"layout/internal/handler"
)

func VisitorRouter(Router *gin.RouterGroup, router *handler.Router) {
	{
		indexRouter := Router.Group("")
		_ = indexRouter
	}
}
