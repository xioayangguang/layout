package middleware

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"layout/global"
	"layout/internal/response"
	"layout/pkg/contextValue"
	"time"
)

func MustTokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Server", "Tomcat8.0")
		apiAuth := c.Request.Header.Get("ApiAuth")
		if apiAuth == "" {
			response.FailWithCode(c, response.TokenError)
			c.Abort()
			return
		}
		if jsonStr, err := global.Redis.Get(context.Background(), apiAuth).Result(); err != nil {
			response.FailWithCode(c, response.TokenError)
			c.Abort()
		} else {
			var userInfo contextValue.LoginUserInfo
			err = json.Unmarshal([]byte(jsonStr), &userInfo)
			if err != nil {
				response.FailWithCode(c, response.Error)
				c.Abort()
			} else {
				c.Set("u_id", userInfo.Id)
				c.Set("u_info", userInfo)
				c.Next()
			}
		}
	}
}

func ShouldTokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Server", "Tomcat8.0")
		apiAuth := c.Request.Header.Get("ApiAuth")
		if apiAuth == "" {
			c.Next()
		} else {
			if jsonStr, err := global.Redis.Get(context.Background(), apiAuth).Result(); err != nil {
				response.FailWithCode(c, response.TokenError)
				c.Abort()
			} else {
				var userInfo contextValue.LoginUserInfo
				err = json.Unmarshal([]byte(jsonStr), &userInfo)
				if err != nil {
					response.FailWithCode(c, response.TokenError)
					c.Abort()
				} else {
					c.Set("u_id", userInfo.Id)
					c.Set("u_info", userInfo)
					global.Redis.Set(context.Background(), apiAuth, jsonStr, time.Duration(86400*2)*time.Second)
					c.Next()
				}
			}
		}
	}
}
