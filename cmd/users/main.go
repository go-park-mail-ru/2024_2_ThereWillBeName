package main

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/dblogger"
	"2024_2_ThereWillBeName/internal/pkg/logger"
	grpcUsers "2024_2_ThereWillBeName/internal/pkg/user/delivery/grpc"
	"2024_2_ThereWillBeName/internal/pkg/user/delivery/grpc/gen"
	userRepo "2024_2_ThereWillBeName/internal/pkg/user/repo"
	userUsecase "2024_2_ThereWillBeName/internal/pkg/user/usecase"
	"database/sql"
	"flag"
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
	var cfg models.ConfigGrpc
	flag.IntVar(&cfg.Port, "grpc-port", 50052, "gRPC server port")
	flag.StringVar(&cfg.ConnStr, "connStr", "host=tripdb port=5432 user=service password=test dbname=trip sslmode=disable", "PostgreSQL connection string")
	flag.Parse()

	logger := setupLogger()

	storagePath := os.Getenv("AVATAR_STORAGE_PATH")

	db, err := sql.Open("postgres", cfg.ConnStr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	wrappedDB := dblogger.NewDB(db, logger)

	userRepo := userRepo.NewAuthRepository(wrappedDB)
	userUsecase := userUsecase.NewUserUsecase(userRepo, storagePath)

	grpcUsersServer := grpc.NewServer()
	usersHandler := grpcUsers.NewGrpcUserHandler(userUsecase, logger)
	gen.RegisterUserServiceServer(grpcUsersServer, usersHandler)
	reflection.Register(grpcUsersServer)

	go func() {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Printf("gRPC server listening on :%d", cfg.Port)
		if err := grpcUsersServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-stop

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
