package metrics

import (
	"context"
	"google.golang.org/grpc"
	"time"
)

type MetricsHTTP interface {
	IncreaseHits(method, path, service string)
	IncreaseErr(method, path, service string)
	AddDurationToHistogram(method, service string, duration time.Duration)
	TrackSystemMetrics(serviceName string)
	ServerMetricsInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
	RegisterMetrics()
}
