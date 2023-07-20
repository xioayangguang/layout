package utils

import (
	red "github.com/go-redis/redis"
	"math/rand"
	"strconv"
	"sync/atomic"
	"time"
)

const (
	letters     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lockCommand = `if redis.call("GET", KEYS[1]) == ARGV[1] then
    redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])
    return "OK"
else
    return redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2])
end`
	delCommand = `if redis.call("GET", KEYS[1]) == ARGV[1] then
    return redis.call("DEL", KEYS[1])
else
    return 0
end`
	randomLen = 16
	// 默认超时时间，防止死锁
	tolerance = 500
)

type RedisLock struct {
	// redis客户端
	store *red.Client
	// 超时时间
	seconds uint32
	// 锁key
	key string
	// 锁value，防止锁被别人获取到
	id string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewRedisLock(store *red.Client, key string) *RedisLock {
	return &RedisLock{
		store: store,
		key:   "Lock:" + key,
		id:    randomStr(randomLen),
	}
}

//加锁
func (rl *RedisLock) Lock(retryCount uint32) bool {
	// 获取过期时间
	seconds := atomic.LoadUint32(&rl.seconds)
	for {
		if retryCount == 0 {
			return false
		}
		retryCount--
		resp, err := rl.store.Eval(lockCommand, []string{rl.key}, []string{
			rl.id, strconv.Itoa(int(seconds)*1000 + tolerance),
		}).Result()
		if err == red.Nil || err != nil || resp == nil {
			time.Sleep(time.Duration(100) * time.Millisecond)
			continue
		}
		reply, ok := resp.(string)
		if ok && reply == "OK" {
			return true
		}
	}
}

// ReleaseLock 释放锁
func (rl *RedisLock) ReleaseLock() bool {
	resp, err := rl.store.Eval(delCommand, []string{rl.key}, []string{rl.id}).Result()
	if err != nil {
		return false
	}
	reply, ok := resp.(int64)
	if !ok {
		return false
	}
	return reply == 1
}

// SetExpire 需要注意的是需要在Lock()之前调用 不然默认为500ms自动释放
func (rl *RedisLock) SetExpire(seconds int) *RedisLock {
	atomic.StoreUint32(&rl.seconds, uint32(seconds))
	return rl
}

func randomStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
