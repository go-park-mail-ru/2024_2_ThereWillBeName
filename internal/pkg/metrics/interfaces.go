package metrics

import (
	"context"
	"google.golang.org/grpc"
	"time"
)

type MetricsHTTP interface {
	IncreaseHits(string, string)
	IncreaseErr(string, string)
	AddDurationToHistogram(string, time.Duration)
	ServerMetricsInterceptor(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) (interface{}, error)
}
