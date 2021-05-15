package redis

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"time"
)

type callback interface {
	Before(*callbackInfo)
	After(*callbackInfo)
}

type callbackInfo struct {
	ctx        context.Context
	err        error
	span       opentracing.Span
	startTime  time.Time
	attachment interface{}
}

var callbacks []callback

func AddCallback(cb callback) {
	callbacks = append(callbacks, cb)
}

func Before(c *callbackInfo) {
	for _, cb := range callbacks {
		cb.Before(c)
	}
}

func After(c *callbackInfo) {
	for _, cb := range callbacks {
		cb.After(c)
	}
}
