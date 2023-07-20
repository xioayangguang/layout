package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/locxiang/gindebugcharts"
	"horse/global"
	"horse/response"
	"net/http"
)

func InitExtraRouter(r *gin.Engine) {

	if true {
		gindebugcharts.Wrapper(r)
		pprof.Register(r)
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
