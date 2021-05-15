package mysql

import (
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"strings"
	"time"
)

var (
	args      = []string{"command", "err"}
	subsystem = "mysql"
	mysqlCnt  = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: subsystem,
			Name:      "execute_count_total",
			Help:      "mysql execute count total",
		}, args)

	mysqlDur = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Subsystem: subsystem,
			Name:      "request_duration_seconds",
			Help:      "mysql execute duration seconds",
		}, args)

	collectors = []prometheus.Collector{mysqlCnt, mysqlDur}
)

func init() {
	for _, collector := range collectors {
		err := prometheus.Register(collector)
		if err != nil {
			log.Println("prometheus register error", err)
		}
	}
	AddCallback(&mysqlMetrics{})
}

type mysqlMetrics struct {
}

func (m *mysqlMetrics) Before(c *callbackInfo) {
	c.startTime = time.Now()
}

func (m *mysqlMetrics) After(c *callbackInfo) {
	command := "unknown"
	if s := strings.Split(c.sql, " "); len(s) > 0 {
		command = s[0]
	}
	// todo： table
	err := ""
	if c.err != nil {
		err = c.err.Error()
	}
	labels := []string{command, err}
	mysqlCnt.WithLabelValues(labels...).Inc()
	elapsed := time.Since(c.startTime).Seconds()
	mysqlDur.WithLabelValues(labels...).Observe(elapsed)
}
