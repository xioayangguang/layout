package middleware

import (
	"github.com/gin-gonic/gin"
	response2 "layout/internal/response"
	"layout/pkg/helper/speedLimit"
)

func SpeedLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ApiAuth := c.Request.Header.Get("ApiAuth")
		if ApiAuth != "" && speedLimit.SpeedLimit(ApiAuth, 1, 10) {
			response2.FailWithCode(c, response2.RateIsTooHigh)
			c.Abort()
			return
		}
		ip := c.ClientIP()
		if speedLimit.SpeedLimit(ip, 1, 10) {
			response2.FailWithCode(c, response2.RateIsTooHigh)
			c.Abort()
			return
		}
		c.Next()
	}
}
