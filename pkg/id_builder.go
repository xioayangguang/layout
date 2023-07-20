package utils

import (
	"horse/global"
)

func IdBuilder(key string, initCallback func() int) int {
	timer := global.Redis.Incr(key)
	count := timer.Val()
	if count == 1 {
		count = int64(initCallback())
		count++
		global.Redis.Set(key, count, 0)
	}
	return int(count)
}
