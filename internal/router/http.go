package router

import (
	"github.com/gin-gonic/gin"
	"layout/internal/handler"
	"layout/internal/router/api"
	"layout/internal/router/h5"
)

func NewServerHTTP(
	userHandler handler.UserHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	api.InitApiRouter(r, userHandler)
	h5.InitApiRouter(r, userHandler)
	InitApiRouter(r, userHandler)
	InitExtraRouter(r)
	return r
}
