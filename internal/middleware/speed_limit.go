package middleware

import (
	"github.com/gin-gonic/gin"
)

func SpeedLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		//ApiAuth := c.Request.Header.Get("ApiAuth")
		//if ApiAuth != "" && utils.SpeedLimit(ApiAuth, 1, 10) {
		//	response.FailWithCode(response.RateIsTooHigh, c)
		//	c.Abort()
		//	return
		//}
		//ip := c.ClientIP()
		//if utils.SpeedLimit(ip, 1, 10) {
		//	response.FailWithCode(response.RateIsTooHigh, c)
		//	c.Abort()
		//	return
		//}
		c.Next()
	}
}
