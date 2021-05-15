package middleware

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"strconv"
	"time"
)

var (
	args      = []string{"code", "status", "method", "host", "url"}
	subsystem = "gf"
	reqCnt    = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: subsystem,
			Name:      "requests_total",
			Help:      "How many HTTP requests processed",
		}, args)

	reqDur = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Subsystem: subsystem,
			Name:      "request_duration_seconds",
			Help:      "HTTP request latencies in seconds",
		}, args)

	collectors = []prometheus.Collector{reqCnt, reqDur}
)

func init() {
	for _, collector := range collectors {
		err := prometheus.Register(collector)
		if err != nil {
			log.Println("prometheus register error", err)
		}
	}
}

func Metrics() func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		if r.Request.RequestURI == "/metrics" {
			r.Middleware.Next()
			return
		}

		start := time.Now()

		r.Middleware.Next()

		codeString := "ok"
		if s := r.Get("code"); s != nil {
			codeString = fmt.Sprintf("%v", s)
		}

		httpStatus := strconv.Itoa(r.Response.Writer.Status)
		url := fmt.Sprintf("[%s]%s", r.Request.Method, r.Request.RequestURI)
		labels := []string{codeString, httpStatus, r.Request.Method, r.Host, url}

		reqCnt.WithLabelValues(labels...).Inc()

		elapsed := float64(time.Since(start)) / float64(time.Second)
		reqDur.WithLabelValues(labels...).Observe(elapsed)
	}
}
