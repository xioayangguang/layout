package core

import "github.com/gin-gonic/gin"

type BusinessHandleFunc func(c *BusinessContext)

func WrapContext(handle BusinessHandleFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("trace", "假设这是一个调用链追踪sdk")
		businessCtx := BusinessContext{c}
		handle(&businessCtx)
	}
}
