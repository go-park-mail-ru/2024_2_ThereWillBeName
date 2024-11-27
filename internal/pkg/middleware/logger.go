package middleware

import (
	log "2024_2_ThereWillBeName/internal/pkg/logger"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
)

const (
	RequestIDKey contextKey = "request_id"
	MethodKey    contextKey = "method"
	PathKey      contextKey = "path"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func RequestLoggerMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := uuid.New().String()

			logCtx := r.Context()
			logCtx = log.AppendCtx(logCtx, slog.String("request_id", requestID))
			logCtx = log.AppendCtx(logCtx, slog.String("method", r.Method))
			logCtx = log.AppendCtx(logCtx, slog.String("path", r.URL.Path))
			r = r.WithContext(logCtx)

			startTime := time.Now()
			logger.Info("Request started",
				slog.String("request_id", requestID),
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Time("start_time", startTime),
			)

			rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

			next.ServeHTTP(rw, r)

			duration := time.Since(startTime)

			logger.Info("Request completed",
				slog.String("request_id", requestID),
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Int("status_code", rw.statusCode),
				slog.Duration("duration", duration),
			)
		})
	}
}