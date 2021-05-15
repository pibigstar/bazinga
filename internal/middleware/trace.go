package middleware

import (
	"bytes"
	"github.com/gogf/gf/net/ghttp"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/pibigstar/bazinga/internal/consts"
	"github.com/pibigstar/bazinga/utils/trace"
	"github.com/uber/jaeger-client-go"
	"net/http"
)

// gin 集成 trace
func Trace() func(r *ghttp.Request) {
	return func(r *ghttp.Request) {

		tracer := trace.GetTracer()
		if tracer == nil {
			r.Middleware.Next()
			return
		}
		var span opentracing.Span
		extract, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Request.Header))
		if err == nil {
			// 继承父span
			span = opentracing.StartSpan(r.Request.URL.Path, opentracing.ChildOf(extract))
		} else {
			// 初始化新的span
			span = opentracing.StartSpan(r.Request.URL.Path)
		}

		// 这里设置一个root span，本次请求的其他span挂在root span下面
		r.SetCtxVar(consts.RootSpan, span)
		r.SetCtx(opentracing.ContextWithSpan(r.Context(), span))
		r.SetCtxVar(consts.Endpoint, r.RequestURI)

		ext.HTTPMethod.Set(span, r.Request.Method)
		ext.HTTPUrl.Set(span, r.Request.URL.Path)

		defer func() {
			ext.HTTPStatusCode.Set(span, uint16(r.Response.Status))
			if r.Response.Status >= http.StatusInternalServerError {
				ext.Error.Set(span, true)
			}
			span.Finish()
		}()

		// 记录body信息
		if b := r.Request.GetBody; b != nil {
			if body, err := b(); err == nil {
				var buff bytes.Buffer
				if _, err = buff.ReadFrom(body); err == nil {
					span.LogKV("form-body", buff.String())
				}
			}
		}
		// 设置context
		if sp, ok := span.Context().(jaeger.SpanContext); ok {
			r.Response.Writer.Header().Set("x-request-id", sp.TraceID().String())
			r.Response.Header().Set("x-trace-id", sp.TraceID().String())
			r.Response.Header().Set("X-Span-id", sp.SpanID().String())
			r.SetCtxVar(consts.TraceId, sp.TraceID().String())
		}
		r.Middleware.Next()
	}
}
