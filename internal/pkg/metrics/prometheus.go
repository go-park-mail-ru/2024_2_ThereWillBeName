package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusMetrics struct {
	hits      *prometheus.CounterVec
	errors    *prometheus.CounterVec
	durations *prometheus.HistogramVec
}

func NewPrometheusMetrics() *PrometheusMetrics {
	hits := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_method_hits_total",
			Help: "Total number of gRPC method calls.",
		},
		[]string{"method"},
	)
	errors := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_method_errors_total",
			Help: "Total number of gRPC method errors.",
		},
		[]string{"method"},
	)
	durations := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "grpc_method_duration_seconds",
			Help:    "Histogram of gRPC method call durations.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)

	prometheus.MustRegister(hits, errors, durations)

	return &PrometheusMetrics{
		hits:      hits,
		errors:    errors,
		durations: durations,
	}
}

func (m *PrometheusMetrics) IncreaseHits(method string) {
	m.hits.WithLabelValues(method).Inc()
}

func (m *PrometheusMetrics) IncreaseErr(method string) {
	m.errors.WithLabelValues(method).Inc()
}

func (m *PrometheusMetrics) AddDurationToHistogram(method string, duration time.Duration) {
	m.durations.WithLabelValues(method).Observe(duration.Seconds())
}
