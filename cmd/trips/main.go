package main

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/logger"
	metricsMw "2024_2_ThereWillBeName/internal/pkg/metrics/middleware"
	grpcTrips "2024_2_ThereWillBeName/internal/pkg/trips/delivery/grpc"
	"2024_2_ThereWillBeName/internal/pkg/trips/delivery/grpc/gen"
	tripRepo "2024_2_ThereWillBeName/internal/pkg/trips/repo"
	tripUsecase "2024_2_ThereWillBeName/internal/pkg/trips/usecase"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var cfg models.ConfigGrpc
	flag.IntVar(&cfg.Port, "grpc-port", 50053, "gRPC server port")
	flag.StringVar(&cfg.ConnStr, "connStr", "host=tripdb port=5432 user=service password=test dbname=trip sslmode=disable", "PostgreSQL connection string")
	flag.Parse()

	logger := setupLogger()

	db, err := sql.Open("postgres", cfg.ConnStr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	r.PathPrefix("/metrics").Handler(promhttp.Handler())
	httpSrv := &http.Server{
		Addr:              ":8092",
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	go func() {

		logger.Info("Starting HTTP server for metrics on :8092")
		if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error(fmt.Sprintf("HTTP server listen: %s\n", err))
		}
	}()

	tripRepo := tripRepo.NewTripRepository(db)
	tripUsecase := tripUsecase.NewTripsUsecase(tripRepo)

	metricMw := metricsMw.Create()

	grpcTripsServer := grpc.NewServer(grpc.UnaryInterceptor(metricMw.ServerMetricsInterceptor))
	tripsHandler := grpcTrips.NewGrpcTripHandler(tripUsecase, logger)
	gen.RegisterTripsServer(grpcTripsServer, tripsHandler)
	reflection.Register(grpcTripsServer)

	go func() {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Printf("gRPC server listening on :%d", cfg.Port)
		if err := grpcTripsServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down gRPC server...")
	grpcTripsServer.GracefulStop()
	log.Println("gRPC server gracefully stopped")
}
func setupLogger() *slog.Logger {

	levelEnv := os.Getenv("LOG_LEVEL")
	logLevel := slog.LevelDebug
	if level, err := strconv.Atoi(levelEnv); err == nil {
		logLevel = slog.Level(level)
	}

	opts := logger.PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: logLevel,
		},
	}

	handler := logger.NewPrettyHandler(os.Stdout, opts)

	return slog.New(handler)
}
