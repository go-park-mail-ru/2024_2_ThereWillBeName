package middleware

import (
	"2024_2_ThereWillBeName/internal/pkg/metrics"
	"context"
	"google.golang.org/grpc"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type GrpcMiddleware struct {
	hits      *prometheus.CounterVec
	errors    *prometheus.CounterVec
	durations *prometheus.HistogramVec
}

func Create() metrics.MetricsHTTP {
	hits := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_method_hits_total",
			Help: "Total number of gRPC method calls.",
		},
		[]string{"method", "path"},
	)
	errors := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_method_errors_total",
			Help: "Total number of gRPC method errors.",
		},
		[]string{"method", "path"},
	)
	durations := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "grpc_method_duration_seconds",
			Help:    "Histogram of gRPC method call durations.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	prometheus.MustRegister(hits, errors, durations)

	return &GrpcMiddleware{
		hits:      hits,
		errors:    errors,
		durations: durations,
	}
}

func (m *GrpcMiddleware) IncreaseHits(method, path string) {
	m.hits.WithLabelValues(method, path).Inc()
}

func (m *GrpcMiddleware) IncreaseErr(method, path string) {
	m.errors.WithLabelValues(method, path).Inc()
}

func (m *GrpcMiddleware) AddDurationToHistogram(method string, duration time.Duration) {
	m.durations.WithLabelValues(method).Observe(duration.Seconds())
}

func (m *GrpcMiddleware) ServerMetricsInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	start := time.Now()
	h, err := handler(ctx, req)
	duration := time.Since(start)

	if err != nil {
		m.IncreaseErr(info.FullMethod, "path")
	}
	m.IncreaseHits(info.FullMethod, "path")
	m.AddDurationToHistogram(info.FullMethod, duration)
	return h, err
}
