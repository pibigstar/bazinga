package resp

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/golang/glog"
	"github.com/pibigstar/bazinga/internal/code"
	"github.com/pibigstar/bazinga/utils/ctxkit"
	"github.com/pibigstar/bazinga/utils/errx"
)

type Resp struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Timestamp float64     `json:"timestamp"`
	TraceId   string      `json:"traceId"`
}

func Success(r *ghttp.Request, data interface{}) {
	r.Context()
	resp := &Resp{
		Code:      200,
		Msg:       "OK",
		Data:      data,
		Timestamp: ctxkit.GetTimestamp(r.Context()),
		TraceId:   ctxkit.GetTraceId(r.Context()),
	}
	if err := r.Response.WriteJson(resp); err != nil {
		glog.Errorln(err)
	}
}

func Error(r *ghttp.Request, err error) {
	errCode := code.Error_Internal.Code()
	if e, ok := err.(errx.ErrX); ok {
		errCode = e.Code()
	}
	resp := &Resp{
		Code:      errCode,
		Msg:       err.Error(),
		Timestamp: ctxkit.GetTimestamp(r.Context()),
		TraceId:   ctxkit.GetTraceId(r.Context()),
	}
	if err := r.Response.WriteJson(resp); err != nil {
		glog.Errorln(err)
	}
}
