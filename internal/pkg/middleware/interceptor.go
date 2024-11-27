package middleware

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

type MetricsInterface interface {
	IncreaseHits(method string)
	IncreaseErr(method string)
	AddDurationToHistogram(method string, duration time.Duration)
}

type GrpcMiddleware struct {
	metrics MetricsInterface
}

func NewGrpcMiddleware(metrics MetricsInterface) *GrpcMiddleware {
	return &GrpcMiddleware{
		metrics: metrics,
	}
}
func (m *GrpcMiddleware) ServerMetricsInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()

	resp, err := handler(ctx, req)

	duration := time.Since(start)

	m.metrics.IncreaseHits(info.FullMethod)

	if err != nil {
		m.metrics.IncreaseErr(info.FullMethod)
	}
	m.metrics.AddDurationToHistogram(info.FullMethod, duration)

	return resp, err
}
