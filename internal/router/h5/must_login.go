package h5

import (
	"github.com/gin-gonic/gin"
	"layout/internal/handler/h5"
)

func MustLoginRouter(Router *gin.RouterGroup, router *h5.Router) {
	{
		indexRouter := Router.Group("user")
		_ = indexRouter
	}
}
