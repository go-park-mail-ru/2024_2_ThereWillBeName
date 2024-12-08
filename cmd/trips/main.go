package main

import (
	"2024_2_ThereWillBeName/internal/pkg/config"
	"2024_2_ThereWillBeName/internal/pkg/dblogger"
	"2024_2_ThereWillBeName/internal/pkg/logger"
	grpcTrips "2024_2_ThereWillBeName/internal/pkg/trips/delivery/grpc"
	"2024_2_ThereWillBeName/internal/pkg/trips/delivery/grpc/gen"
	tripRepo "2024_2_ThereWillBeName/internal/pkg/trips/repo"
	tripUsecase "2024_2_ThereWillBeName/internal/pkg/trips/usecase"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	_ "github.com/lib/pq"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	logger := setupLogger()

	db, err := dblogger.SetupDBPool(cfg, logger)
	if err != nil {
		log.Fatalf("failed to initialize connection pool: %v", err)
	}
	defer db.Close()

	wrappedDB := dblogger.NewDB(db, logger)

	tripRepo := tripRepo.NewTripRepository(wrappedDB)
	tripUsecase := tripUsecase.NewTripsUsecase(tripRepo)

	grpcTripsServer := grpc.NewServer()
	tripsHandler := grpcTrips.NewGrpcTripHandler(tripUsecase, logger)
	gen.RegisterTripsServer(grpcTripsServer, tripsHandler)
	reflection.Register(grpcTripsServer)

	go func() {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Grpc.TripPort))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Printf("gRPC server listening on :%d", cfg.Grpc.TripPort)
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
