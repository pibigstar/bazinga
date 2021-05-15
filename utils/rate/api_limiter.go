package rate

import (
	"github.com/gogf/gf/frame/g"
	"go.uber.org/ratelimit"
)

var (
	// 每秒可以接受100个请求
	apiLimiter ratelimit.Limiter
)

func GetApiLimiter() ratelimit.Limiter {
	if apiLimiter == nil {
		apiLimiter = ratelimit.New(g.Cfg().GetInt("rate.apiLimit"))
	}
	return apiLimiter
}
