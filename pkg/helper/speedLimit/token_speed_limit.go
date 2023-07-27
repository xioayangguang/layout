package speedLimit

import (
	"golang.org/x/time/rate"
	"sync"
	"time"
)

// Limiter 单一限速器
type Limiter struct {
	limiter  rate.Limiter
	lastTime time.Time
	key      string
}

func (l *Limiter) Allow() bool {
	return l.limiter.Allow()
}

// NewLimiters 限速器集合
func NewLimiters() *Limiters {
	ls := &Limiters{}
	go ls.cleanLimiter()
	return ls
}

type Limiters struct {
	limiters sync.Map
}

// GetLimiter
// 第一个参数r Limit：产生令牌的速率，也就是每秒往桶中放入多少个令牌。
// 第二个参数b int：令牌桶的大小。
func (ls *Limiters) GetLimiter(key string, r rate.Limit, b int) *Limiter {
	if v, ok := ls.limiters.Load(key); ok {
		l := v.(*Limiter)
		l.lastTime = time.Now()
		return l
	}
	l := &Limiter{
		limiter:  *rate.NewLimiter(r, b),
		lastTime: time.Now(),
		key:      key,
	}
	ls.limiters.Store(key, l)
	return l
}

func (ls *Limiters) cleanLimiter() {
	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()
	for {
		<-ticker.C
		ls.limiters.Range(func(key, value interface{}) bool {
			l := value.(*Limiter)
			if time.Now().Sub(l.lastTime) > time.Minute {
				ls.limiters.Delete(key)
			}
			return true
		})
	}
}
