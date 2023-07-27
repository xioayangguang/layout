package speedLimit

import (
	"context"
	"github.com/redis/go-redis/v9"
	"layout/global"
	"strconv"
	"time"
)

func SpeedLimit(c context.Context, key string, period, maxCount int) bool {
	key = "SpeedLimit:" + key
	msecTime := int(time.Now().UnixNano() / 1e6)
	pipe := global.Redis.Pipeline()
	pipe.ZRemRangeByRank(c, key, 0, -(int64(maxCount) + 1))
	count := pipe.ZCount(c, key, strconv.Itoa(msecTime-period*1000), strconv.Itoa(msecTime))
	_, _ = pipe.Exec(c)
	if count.Val() >= int64(maxCount) {
		return true
	} else {
		pipe := global.Redis.Pipeline()
		members := []redis.Z{
			redis.Z{Score: float64(msecTime), Member: msecTime},
		}
		pipe.ZAdd(c, key, members...)
		pipe.Expire(c, key, time.Duration(period)*time.Second)
		_, _ = pipe.Exec(c)
		return false
	}
}
