package main

import (
	"2024_2_ThereWillBeName/internal/pkg/config"
	"2024_2_ThereWillBeName/internal/pkg/dblogger"
	"2024_2_ThereWillBeName/internal/pkg/logger"
	metricsMw "2024_2_ThereWillBeName/internal/pkg/metrics/middleware"
	grpcUsers "2024_2_ThereWillBeName/internal/pkg/user/delivery/grpc"
	"2024_2_ThereWillBeName/internal/pkg/user/delivery/grpc/gen"
	userRepo "2024_2_ThereWillBeName/internal/pkg/user/repo"
	userUsecase "2024_2_ThereWillBeName/internal/pkg/user/usecase"
	"database/sql"
	"errors"
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
	cfg := config.Load()

	logger := setupLogger()

	metricMw := metricsMw.Create()
	metricMw.RegisterMetrics()

	storagePath := os.Getenv("AVATAR_STORAGE_PATH")

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Database.DbHost, cfg.Database.DbPort, cfg.Database.DbUser, cfg.Database.DbPass, cfg.Database.DbName))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()
	log.Printf("DB_HOST: %s, DB_PORT: %d, DB_USER: %s, DB_PASS: %s, DB_NAME: %s", cfg.Database.DbHost, cfg.Database.DbPort, cfg.Database.DbUser, cfg.Database.DbPass, cfg.Database.DbName)
	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	wrappedDB := dblogger.NewDB(db, logger)

	r := mux.NewRouter()
	r.Handle("/metrics", promhttp.Handler())
	httpSrv := &http.Server{
		Addr:              ":8093",
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	go func() {

		logger.Info("Starting HTTP server for metrics on :8093")
		if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error(fmt.Sprintf("HTTP server listen: %s\n", err))
		}
	}()

	userRepo := userRepo.NewAuthRepository(wrappedDB)
	userUsecase := userUsecase.NewUserUsecase(userRepo, storagePath)

	grpcUsersServer := grpc.NewServer(grpc.UnaryInterceptor(metricMw.ServerMetricsInterceptor))
	usersHandler := grpcUsers.NewGrpcUserHandler(userUsecase, logger)
	gen.RegisterUserServiceServer(grpcUsersServer, usersHandler)
	reflection.Register(grpcUsersServer)

	go func() {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Grpc.UserPort))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Printf("gRPC server listening on :%d", cfg.Grpc.UserPort)
		if err := grpcUsersServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	go func() {
		ticker := time.NewTicker(15 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				metricMw.TrackSystemMetrics("users")
			case <-stop:
				return
			}
		}
	}()

	log.Println("Shutting down gRPC server...")
	grpcUsersServer.GracefulStop()
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
