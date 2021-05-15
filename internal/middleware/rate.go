package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/pibigstar/bazinga/internal/code"
	"github.com/pibigstar/bazinga/utils/errx"
	"github.com/pibigstar/bazinga/utils/rate"
	"github.com/pibigstar/bazinga/utils/resp"
)

// IP限流
func IpLate() func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		limiter := rate.GetIPLimiter().GetLimiter(r.GetRemoteIp())
		if !limiter.Allow() {
			resp.Error(r, errx.NewWithCode(code.Error_Internal))
		}
		r.Middleware.Next()
	}
}

// 接口限流
func ApiLate() func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		rate.GetApiLimiter().Take()
		r.Middleware.Next()
	}
}
