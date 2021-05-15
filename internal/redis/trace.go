package redis

import (
	"github.com/go-redis/redis"
	"github.com/gogf/gf/util/gconv"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/pibigstar/bazinga/utils/trace"
	"strings"
)

const (
	defaultComponentName = "redis"
)

type redisTrace struct {
}

func init() {
	AddCallback(&redisTrace{})
}

func (*redisTrace) Before(c *callbackInfo) {
	parentSpan := opentracing.SpanFromContext(c.ctx)
	if parentSpan == nil {
		return
	}
	switch c.attachment.(type) {
	case redis.Cmder:
		cmder := c.attachment.(redis.Cmder)
		span := buildSpan(parentSpan, strings.ToUpper(cmder.Name()), []redis.Cmder{cmder})
		c.span = span

	case []redis.Cmder:
		cmders := c.attachment.([]redis.Cmder)
		span := buildSpan(parentSpan, "pipeline", cmders)
		c.span = span
	}
}

func (*redisTrace) After(c *callbackInfo) {
	if c.span == nil {
		return
	}
	c.span.Finish()
}

func buildSpan(parentSpan opentracing.Span, operationName string, cmders []redis.Cmder) opentracing.Span {
	tracerClient := trace.GetTracer()
	span := tracerClient.StartSpan(operationName, opentracing.ChildOf(parentSpan.Context()))

	ext.Component.Set(span, defaultComponentName)
	ext.PeerAddress.Set(span, config.Addr)

	method, key := buildMethodAndKeys(cmders)
	span.SetTag("method", method)
	span.SetTag("keys", key)
	return span
}

func buildMethodAndKeys(cmders []redis.Cmder) (string, string) {
	var methods, keys []string
	for _, c := range cmders {
		methods = append(methods, c.Name())
		if len(c.Args()) >= 2 {
			keys = append(keys, gconv.String(c.Args()[1]))
		}
	}
	return strings.Join(methods, "->"), strings.Join(keys, ";")
}
