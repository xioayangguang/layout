package router

import (
	"github.com/gin-gonic/gin"
	"layout/global"
	"layout/internal/handler"
	"layout/internal/router/api"
	"layout/internal/router/h5"
	"layout/pkg/helper/rotatelogs"
)

func NewServerHTTP(
	userHandler handler.UserHandler,
) *gin.Engine {
	var r *gin.Engine
	if !global.Config.Debug {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		r.Use(gin.LoggerWithWriter(rotatelogs.GetRotateLogs("output")), gin.RecoveryWithWriter(rotatelogs.GetRotateLogs("recovery")))
	} else {
		r = gin.Default()
	}

	api.InitApiRouter(r, userHandler)
	h5.InitApiRouter(r, userHandler)
	InitApiRouter(r, userHandler)
	InitExtraRouter(r)
	return r
}
