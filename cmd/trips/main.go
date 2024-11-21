package main

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/trips/delivery/grpc/gen"
	tripRepo "2024_2_ThereWillBeName/internal/pkg/trips/repo"
	tripUsecase "2024_2_ThereWillBeName/internal/pkg/trips/usecase"
	"database/sql"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var cfg models.ConfigGrpc
	flag.IntVar(&cfg.Port, "grpc-port", 50053, "gRPC server port")
	flag.StringVar(&cfg.ConnStr, "connStr", "host=tripdb port=5432 user=service password=test dbname=trip sslmode=disable", "PostgreSQL connection string")
	flag.Parse()

	db, err := sql.Open("postgres", cfg.ConnStr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	tripRepo := tripRepo.NewTripRepository(db)
	tripUsecase := tripUsecase.NewTripUsecase(tripRepo)

	grpcTripsServer := grpc.NewServer()
	tripsHandler := grpc.NewGrpcTripsHandler(tripUsecase)
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
