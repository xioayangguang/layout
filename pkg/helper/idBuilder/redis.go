package idBuilder

import (
	"context"
	"layout/global"
)

func Generate(key string, initCallback func() int) int {
	timer := global.Redis.Incr(context.Background(), key)
	count := timer.Val()
	if count == 1 {
		count = int64(initCallback())
		count++
		global.Redis.Set(context.Background(), key, count, 0)
	}
	return int(count)
}
