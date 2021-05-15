package resp

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"time"
)

type Resp struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Timestamp float64     `json:"timestamp"`
	TraceId   string      `json:"traceId"`
}

func Success(r *ghttp.Request, data interface{}) *Resp {
	r.Context()
	return &Resp{
		Code:      200,
		Msg:       "OK",
		Data:      data,
		Timestamp: getTimestamp(r),
		TraceId:   getTraceId(r),
	}
}

func Error(r *ghttp.Request, msg string, codes ...int) *Resp {
	code := 500
	if len(codes) > 0 {
		code = codes[0]
	}
	return &Resp{
		Code:      code,
		Msg:       msg,
		Data:      nil,
		Timestamp: getTimestamp(r),
		TraceId:   getTraceId(r),
	}
}

func getTraceId(r *ghttp.Request) string {
	return gconv.String(r.Context().Value("traceId"))
}

func getTimestamp(r *ghttp.Request) float64 {
	if t := r.Context().Value("timestamp"); t != nil {
		f := time.Now().Unix() - gconv.Int64(t)
		return float64(f) / 60
	}
	return 0
}
