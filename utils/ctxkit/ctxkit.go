package ctxkit

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/pibigstar/bazinga/internal/consts"
	"github.com/pibigstar/bazinga/utils/timex"
)

// 从context获取最初的端点调用
func GetEndpoint(ctx context.Context) string {
	if ctx == nil {
		return consts.UNKNOWN
	}
	e := ctx.Value(consts.Endpoint)
	if e == nil {
		return consts.UNKNOWN
	}
	return gconv.String(e)
}

func GetTraceId(ctx context.Context) string {
	return gconv.String(ctx.Value(consts.TraceId))
}

func GetTimestamp(ctx context.Context) float64 {
	if ctx == nil {
		return 0
	}
	if t := ctx.Value(consts.Timestamp); t != nil {
		f := timex.UnixMill() - gconv.Int64(t)
		return float64(f) / 1000
	}
	return 0
}
