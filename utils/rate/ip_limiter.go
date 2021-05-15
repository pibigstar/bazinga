package rate

import (
	"github.com/gogf/gf/frame/g"
	"golang.org/x/time/rate"
	"sync"
)

var (
	// IP 限流
	ipLimiter *IPRateLimit
)

func GetIPLimiter() *IPRateLimit {
	if ipLimiter == nil {
		ipLimiter = NewIPRateLimiter(3, g.Cfg().GetInt("rate.ipLimit"))
	}
	return ipLimiter
}

type IPRateLimit struct {
	mu      *sync.RWMutex
	limiter map[string]*rate.Limiter
	r       rate.Limit
	b       int
}

func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimit {
	return &IPRateLimit{
		limiter: make(map[string]*rate.Limiter),
		mu:      &sync.RWMutex{},
		r:       r, // 1s创建多少个令牌
		b:       b, // 最大存储多少个令牌
	}
}

func (i *IPRateLimit) AddIp(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()
	limiter := rate.NewLimiter(i.r, i.b)
	i.limiter[ip] = limiter
	return limiter
}

func (i *IPRateLimit) GetLimiter(ip string) *rate.Limiter {
	if limiter, ok := i.limiter[ip]; ok {
		return limiter
	}
	return i.AddIp(ip)
}
