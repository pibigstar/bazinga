package server

import (
	"context"
	"github.com/pibigstar/bazinga/internal/consts"
	"github.com/pibigstar/bazinga/utils/ctxkit"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
	"strings"
	"time"
)

var (
	args      = []string{"service", "method", "endpoint", "err"}
	subsystem = "grpc"
	grpcCnt   = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: subsystem,
			Name:      "execute_count_total",
			Help:      "grpc execute count total",
		}, args)

	grpcDur = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Subsystem: subsystem,
			Name:      "request_duration_seconds",
			Help:      "grpc execute duration seconds",
		}, args)

	collectors = []prometheus.Collector{grpcCnt, grpcDur}
)

func init() {
	for _, collector := range collectors {
		err := prometheus.Register(collector)
		if err != nil {
			log.Println("prometheus register error", err)
		}
	}
}

func UnaryServerMetrics() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		metrics := &grpcMetrics{
			ctx:        ctx,
			fullMethod: info.FullMethod,
		}
		metrics.Before()
		resp, err = handler(ctx, req)
		metrics.err = err
		metrics.After()
		return resp, err
	}
}

type grpcMetrics struct {
	ctx        context.Context
	err        error
	startTime  time.Time
	fullMethod string
}

func (m *grpcMetrics) Before() {
	m.startTime = time.Now()
}

func (m *grpcMetrics) After() {
	service, method := splitMethodName(m.fullMethod)

	var err string
	if m.err != nil {
		s, _ := status.FromError(m.err)
		err = s.Code().String()
	}
	endpoint := ctxkit.GetEndpoint(m.ctx)

	labels := []string{service, method, endpoint, err}
	grpcCnt.WithLabelValues(labels...).Inc()
	elapsed := time.Since(m.startTime).Seconds()
	grpcDur.WithLabelValues(labels...).Observe(elapsed)
}

func splitMethodName(fullMethodName string) (string, string) {
	fullMethodName = strings.TrimPrefix(fullMethodName, "/")
	if i := strings.Index(fullMethodName, "/"); i >= 0 {
		return fullMethodName[:i], fullMethodName[i+1:]
	}
	return consts.UNKNOWN, consts.UNKNOWN
}
