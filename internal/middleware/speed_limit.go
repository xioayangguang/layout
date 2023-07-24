package middleware

import (
	"github.com/gin-gonic/gin"
	"layout/internal/response"
	"layout/pkg/helper/speedLimit"
)

// SpeedLimit 分布式限速，要是网关做了ip哈希可以直接本地限速提高效率，减少网络io
func SpeedLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ApiAuth := c.Request.Header.Get("ApiAuth")
		if ApiAuth != "" && speedLimit.SpeedLimit(ApiAuth, 1, 10) {
			response.FailWithCode(c, response.RateIsTooHigh)
			c.Abort()
			return
		}
		ip := c.ClientIP()
		if speedLimit.SpeedLimit(ip, 1, 10) {
			response.FailWithCode(c, response.RateIsTooHigh)
			c.Abort()
			return
		}
		c.Next()
	}
}
