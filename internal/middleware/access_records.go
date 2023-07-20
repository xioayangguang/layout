package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"time"
)

// 用户访问记录
func AccessRecords(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		if uInfo, ok := c.Get("u_info"); ok {
			uInfo := uInfo.(dto.LoginUserInfo)
			_, _ = rdb.SetBit(time.Now().Format("2006-01-02"), int64(uInfo.Serial-1), 1).Result()
		}
		c.Next()
	}
}
