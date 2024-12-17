package metrics

import (
	"2024_2_ThereWillBeName/internal/pkg/metrics"
	"context"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
)

type SystemMetrics struct {
	cpuUsage    *prometheus.GaugeVec
	memoryUsage *prometheus.GaugeVec
	diskUsage   *prometheus.GaugeVec
}

type GrpcMiddleware struct {
	hits      *prometheus.CounterVec
	errors    *prometheus.CounterVec
	durations *prometheus.HistogramVec
	// statusCodes  *prometheus.CounterVec // Добавляем метрику для статусов кодов
	systemMetric *SystemMetrics
	mu           sync.Mutex
}

func Create() metrics.MetricsHTTP {
	middleware := &GrpcMiddleware{
		hits: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "http_method_hits_total",
			Help: "Total number of http method calls across all services",
		}, []string{"method", "path", "status_code"}),

		errors: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "http_method_errors_total",
			Help: "Total number of http method errors across all services",
		}, []string{"method", "path", "service"}),

		durations: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "http_method_duration_seconds",
			Help:    "Histogram of http method call durations across services",
			Buckets: prometheus.DefBuckets,
		}, []string{"method", "service"}),

		systemMetric: &SystemMetrics{
			cpuUsage: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name: "service_cpu_usage_percent",
				Help: "CPU usage percentage per service",
			}, []string{"service", "hostname"}),

			memoryUsage: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name: "service_memory_usage_bytes",
				Help: "Memory usage in bytes per service",
			}, []string{"service", "hostname"}),

			diskUsage: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name: "service_disk_usage_percent",
				Help: "Disk usage percentage per service",
			}, []string{"service", "hostname"}),
		},
	}

	return middleware
}

func (m *GrpcMiddleware) RegisterMetrics() {
	prometheus.MustRegister(m.hits)
	prometheus.MustRegister(m.errors)
	prometheus.MustRegister(m.durations)
	prometheus.MustRegister(m.systemMetric.cpuUsage)
	prometheus.MustRegister(m.systemMetric.memoryUsage)
	prometheus.MustRegister(m.systemMetric.diskUsage)
}

func (m *GrpcMiddleware) IncreaseHits(method, path, statusCode string) {
	m.hits.WithLabelValues(method, path, statusCode).Inc()
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
		m.systemMetric.diskUsage.WithLabelValues(serviceName, hostname).Set(diskStat.UsedPercent)
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
	log.Println(err)
	duration := time.Since(start)

	if err != nil {
		m.IncreaseErr(info.FullMethod, path, serviceName)
	}
	m.AddDurationToHistogram(info.FullMethod, serviceName, duration)

	return h, err
}

func (m *GrpcMiddleware) MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := NewResponseWriter(w)
		next.ServeHTTP(rw, r)
		normalizedPath := normalizePath(r.URL.Path)

		m.IncreaseHits(r.Method, normalizedPath, strconv.Itoa(rw.statusCode))
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func extractServiceName(fullMethod string) string {
	parts := strings.Split(fullMethod, "/")
	if len(parts) >= 2 {
		if len(parts[1]) >= 2 {
			return strings.Split(parts[1], ".")[0]
		}
	}
	return "unknown"
}

func extractPath(fullMethod string) string {
	parts := strings.Split(fullMethod, "/")
	if len(parts) >= 3 {
		log.Println(fullMethod)
		log.Println(parts[2])
		return parts[2]
	}
	return "unknown"
}

func normalizePath(path string) string {
	re := regexp.MustCompile(`/\d+`)
	return re.ReplaceAllString(path, "/:id")
}
