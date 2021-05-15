package middleware

import (
	"context"
	"github.com/gogf/gf/net/ghttp"
	"github.com/pibigstar/bazinga/internal/consts"
	"github.com/pibigstar/bazinga/utils/timex"
)

// 接口耗时
func Timestamp() func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		ctx := context.WithValue(r.Context(), consts.Timestamp, timex.UnixMill())
		r.SetCtx(ctx)
		r.Middleware.Next()
	}
}
