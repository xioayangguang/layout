package middleware

import (
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Timeout(t time.Duration) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(t),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(func(c *gin.Context) {
			c.JSON(http.StatusRequestTimeout, gin.H{
				"code": 1,
				"data": nil,
				"msg":  "RequestTimeout",
			})
		}),
	)
}
