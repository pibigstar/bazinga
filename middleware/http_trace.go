package middleware

import (
	"bytes"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/zipkin"
)

// gin 集成 trace
func Trace() func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		filters := g.Cfg().GetArray("trace.filter")
		for _, filter := range filters {
			if f, ok := filter.(string); ok {
				if gstr.ContainsI(r.URL.Path, f) {
					r.Middleware.Next()
					return
				}
			}
		}

		tracer := NewTracer()
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
		defer span.Finish()

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
			// 这里设置一个root span，本次请求的其他span挂在root span下面

			r.SetCtxVar("root_span", span)
			r.Response.Writer.Header().Set("x-request-id", sp.TraceID().String())
			r.Response.Header().Set("x-trace-id", sp.TraceID().String())
			r.Response.Header().Set("X-Span-id", sp.SpanID().String())
		}
		r.Middleware.Next()
	}
}

func NewTracer() opentracing.Tracer {
	// 判断是否已注册了，单例模式
	if opentracing.IsGlobalTracerRegistered() {
		return opentracing.GlobalTracer()
	}
	cfg := &config.Configuration{
		ServiceName: "pibigstar",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LocalAgentHostPort: g.Cfg().GetString("trace.addr"),
			LogSpans:           true,
		},
		Headers: &jaeger.HeadersConfig{
			JaegerDebugHeader:        "x-debug-id",
			JaegerBaggageHeader:      "x-baggage",
			TraceContextHeaderName:   "x-trace-id",
			TraceBaggageHeaderPrefix: "x-ctx",
		},
	}

	propagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
	tracer, _, err := cfg.NewTracer(
		config.Logger(jaeger.StdLogger),
		config.Injector(opentracing.HTTPHeaders, propagator),
		config.Extractor(opentracing.HTTPHeaders, propagator),
		config.ZipkinSharedRPCSpan(true),
		config.MaxTagValueLength(256),
		config.PoolSpans(true),
	)
	if err != nil {
		panic(fmt.Sprintf("Init failed: %v\n", err))
	}
	// 设置到全局里
	opentracing.SetGlobalTracer(tracer)

	return tracer
}
