package mysql

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/pibigstar/bazinga/internal/consts"
	"github.com/pibigstar/bazinga/utils/trace"
	"strings"
)

type mysqlTrace struct {
}

func init() {
	AddCallback(&mysqlTrace{})
}

func (m *mysqlTrace) Before(c *callbackInfo) {
	parentSpan := opentracing.SpanFromContext(c.ctx)
	if parentSpan == nil {
		return
	}
	var options []opentracing.StartSpanOption
	options = append(options, opentracing.ChildOf(parentSpan.Context()))
	operation := strings.ToUpper(strings.Split(c.sql, " ")[0])

	span := trace.GetTracer().StartSpan(operation, options...)
	ext.Component.Set(span, "MySQL")

	span.SetTag("db.method", operation)
	span.SetTag("db.sql", c.sql)
	span.SetTag("db.traceId", c.ctx.Value(consts.TraceId))
	c.span = span
}

func (m *mysqlTrace) After(c *callbackInfo) {
	if c.span == nil {
		return
	}
	hasErr := c.err != nil
	ext.Error.Set(c.span, hasErr)
	if hasErr {
		c.span.SetTag("db.err", c.err)
	}
	c.span.Finish()
}
