package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"layout/global"
	"layout/pkg/contextValue"
	"time"
)

// 用户访问记录
func AccessRecords() gin.HandlerFunc {
	return func(c *gin.Context) {
		if uInfo, ok := c.Get("u_info"); ok {
			uInfo := uInfo.(contextValue.LoginUserInfo)
			_, _ = global.Redis.SetBit(context.Background(), time.Now().Format("2006-01-02"), int64(uInfo.Serial-1), 1).Result()
		}
		c.Next()
	}
}
