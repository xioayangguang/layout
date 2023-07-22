package speedLimit

import (
	"context"
	"github.com/redis/go-redis/v9"
	"layout/global"
	"strconv"
	"time"
)

func SpeedLimit(key string, period, maxCount int) bool {
	key = "SpeedLimit:" + key
	msecTime := int(time.Now().UnixNano() / 1e6)
	pipe := global.Redis.Pipeline()
	pipe.ZRemRangeByRank(context.Background(), key, 0, -(int64(maxCount) + 1))
	count := pipe.ZCount(context.Background(), key, strconv.Itoa(msecTime-period*1000), strconv.Itoa(msecTime))
	_, _ = pipe.Exec(context.Background())
	if count.Val() >= int64(maxCount) {
		return true
	} else {
		pipe := global.Redis.Pipeline()
		members := []redis.Z{
			redis.Z{Score: float64(msecTime), Member: msecTime},
		}
		pipe.ZAdd(context.Background(), key, members...)
		pipe.Expire(context.Background(), key, time.Duration(period)*time.Second)
		_, _ = pipe.Exec(context.Background())
		return false
	}
}
