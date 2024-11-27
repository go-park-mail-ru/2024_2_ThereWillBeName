package metrics

import (
	"2024_2_ThereWillBeName/internal/pkg/metrics"
	"context"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc"
)

type SystemMetrics struct {
	cpuUsage    *prometheus.GaugeVec
	memoryUsage *prometheus.GaugeVec
	diskUsage   *prometheus.GaugeVec
}

type GrpcMiddleware struct {
	hits         *prometheus.CounterVec
	errors       *prometheus.CounterVec
	durations    *prometheus.HistogramVec
	systemMetric *SystemMetrics
	mu           sync.Mutex
}

func Create() metrics.MetricsHTTP {
	middleware := &GrpcMiddleware{
		hits: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "grpc_method_hits_total",
			Help: "Total number of gRPC method calls across all services",
		}, []string{"method", "path", "service"}),

		errors: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "grpc_method_errors_total",
			Help: "Total number of gRPC method errors across all services",
		}, []string{"method", "path", "service"}),

		durations: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "grpc_method_duration_seconds",
			Help:    "Histogram of gRPC method call durations across services",
			Buckets: prometheus.DefBuckets,
		}, []string{"method", "service"}),

		systemMetric: &SystemMetrics{
			cpuUsage: promauto.NewGaugeVec(prometheus.GaugeOpts{
				Name: "service_cpu_usage_percent",
				Help: "CPU usage percentage per service",
			}, []string{"service", "hostname"}),

			memoryUsage: promauto.NewGaugeVec(prometheus.GaugeOpts{
				Name: "service_memory_usage_bytes",
				Help: "Memory usage in bytes per service",
			}, []string{"service", "hostname"}),

			diskUsage: promauto.NewGaugeVec(prometheus.GaugeOpts{
				Name: "service_disk_usage_percent",
				Help: "Disk usage percentage per service",
			}, []string{"service", "mount_point", "hostname"}),
		},
	}

	return middleware
}

func (m *GrpcMiddleware) IncreaseHits(method, path, service string) {
	m.hits.WithLabelValues(method, path, service).Inc()
}

func (m *GrpcMiddleware) IncreaseErr(method, path, service string) {
	m.errors.WithLabelValues(method, path, service).Inc()
}

func (m *GrpcMiddleware) AddDurationToHistogram(method, service string, duration time.Duration) {
	m.durations.WithLabelValues(method, service).Observe(duration.Seconds())
}

func (m *GrpcMiddleware) TrackSystemMetrics(serviceName string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	hostname, _ := os.Hostname()

	// CPU Usage
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err == nil && len(cpuPercent) > 0 {
		m.systemMetric.cpuUsage.WithLabelValues(serviceName, hostname).Set(cpuPercent[0])
	}

	// Memory Usage
	vmStat, err := mem.VirtualMemory()
	if err == nil {
		m.systemMetric.memoryUsage.WithLabelValues(serviceName, hostname).Set(float64(vmStat.Used))
	}

	// Disk Usage
	diskStat, err := disk.Usage("/")
	if err == nil {
		m.systemMetric.diskUsage.WithLabelValues(serviceName, "/", hostname).Set(diskStat.UsedPercent)
	}
}

func (m *GrpcMiddleware) ServerMetricsInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	serviceName := extractServiceName(info.FullMethod)
	path := extractPath(info.FullMethod)

	h, err := handler(ctx, req)
	duration := time.Since(start)

	if err != nil {
		m.IncreaseErr(info.FullMethod, path, serviceName)
	}
	m.IncreaseHits(info.FullMethod, path, serviceName)
	m.AddDurationToHistogram(info.FullMethod, serviceName, duration)

	return h, err
}

func extractServiceName(fullMethod string) string {
	parts := strings.Split(fullMethod, "/")
	if len(parts) >= 2 {
		return parts[1]
	}
	return "unknown"
}

func extractPath(fullMethod string) string {
	parts := strings.Split(fullMethod, "/")
	if len(parts) >= 3 {
		// Возвращаем путь (третий элемент в массиве)
		return parts[2]
	}
	return "unknown"
}
