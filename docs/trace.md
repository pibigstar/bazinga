# 链路追踪

## 1. 安装jaeger
```bash
docker run -d --restart always -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:latest
```

浏览器访问：http://127.0.0.1:16686  即可看到页面


## 2. opentracing-go使用
> 库地址：github.com/opentracing/opentracing-go

```go
var span opentracing.Span
extract, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Request.Header))
if err == nil {
    // 继承父span
    span = opentracing.StartSpan(r.Request.URL.Path, opentracing.ChildOf(extract))
} else {
    // 初始化新的span
    span = opentracing.StartSpan(r.Request.URL.Path)
}

// 设置root span，本次请求的其他span挂在root span下面
r.SetCtxVar(consts.RootSpan, span)
// 将 span 设置到context中
r.SetCtx(opentracing.ContextWithSpan(r.Context(), span))

ext.HTTPMethod.Set(span, r.Request.Method)
ext.HTTPUrl.Set(span, r.Request.URL.Path)

defer func() {
    ext.HTTPStatusCode.Set(span, uint16(r.Response.Status))
    if r.Response.Status >= http.StatusInternalServerError {
        ext.Error.Set(span, true)
    }
    span.Finish()
}()

// 设置context
if sp, ok := span.Context().(jaeger.SpanContext); ok {
    r.Response.Header().Set("x-trace-id", sp.TraceID().String())
    r.SetCtxVar(consts.TraceId, sp.TraceID().String())
}
```
