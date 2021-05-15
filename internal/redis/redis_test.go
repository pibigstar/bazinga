package redis_test

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/pibigstar/bazinga/internal/redis"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	span := opentracing.StartSpan("test")
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	redis.GetClient(ctx).Set("test", "Hello2", time.Minute*10)
}
