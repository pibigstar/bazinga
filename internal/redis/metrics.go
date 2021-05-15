package redis

import (
	"github.com/go-redis/redis"
	"github.com/gogf/gf/util/gconv"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"strconv"
	"strings"
	"time"
)

var (
	args      = []string{"command", "host", "db", "err"}
	subsystem = "redis"
	redisCnt  = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: subsystem,
			Name:      "execute_count_total",
			Help:      "redis execute count total",
		}, args)

	redisDur = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Subsystem: subsystem,
			Name:      "request_duration_seconds",
			Help:      "redis execute duration seconds",
		}, args)

	collectors = []prometheus.Collector{redisCnt, redisDur}
)

func init() {
	for _, collector := range collectors {
		err := prometheus.Register(collector)
		if err != nil {
			log.Println("prometheus register error", err)
		}
	}
	AddCallback(&redisMetrics{})
}

type redisMetrics struct {
}

func (m *redisMetrics) Before(c *callbackInfo) {
	c.startTime = time.Now()
}

func (m *redisMetrics) After(c *callbackInfo) {
	switch c.attachment.(type) {
	case redis.Cmder:
		cmder := c.attachment.(redis.Cmder)
		m.reportMetrics(cmder, c)

	case []redis.Cmder:
		cmders := c.attachment.([]redis.Cmder)
		for _, cmder := range cmders {
			m.reportMetrics(cmder, c)
		}
	}
}

func (m *redisMetrics) reportMetrics(cmd redis.Cmder, c *callbackInfo) {
	method := strings.ToUpper(cmd.Name())
	err := cmd.Err() != nil && cmd.Err() != redis.Nil

	labels := []string{method, config.Addr, gconv.String(config.DB), strconv.FormatBool(err)}
	redisCnt.WithLabelValues(labels...).Inc()

	elapsed := time.Since(c.startTime).Seconds()
	redisDur.WithLabelValues(labels...).Observe(elapsed)
}
