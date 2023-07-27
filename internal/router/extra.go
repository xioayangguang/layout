package router

import (
	ginpprof "github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/locxiang/gindebugcharts"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "layout/docs"
	"layout/global"
	"layout/internal/response"
	"net/http"
)

func InitExtraRouter(r *gin.Engine) {
	if global.Config.Debug {
		gindebugcharts.Wrapper(r)
		ginpprof.Register(r)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	r.GET("/version", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"go_version": global.GoVersion,
			"build_time": global.BuildTime,
			"git_hash":   global.GitHash,
		})
	})
	r.NoRoute(func(c *gin.Context) {
		c.Header("Server", "Tomcat8.0")
		c.JSON(http.StatusNotFound, response.Response{
			Code: 404,
			Msg:  "路径错误",
		})
	})
}
