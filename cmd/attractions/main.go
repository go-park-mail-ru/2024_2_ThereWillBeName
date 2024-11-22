package main

import (
	"2024_2_ThereWillBeName/internal/models"
	grpcAttractions "2024_2_ThereWillBeName/internal/pkg/attractions/delivery/grpc"
	genPlaces "2024_2_ThereWillBeName/internal/pkg/attractions/delivery/grpc/gen"
	placeRepo "2024_2_ThereWillBeName/internal/pkg/attractions/repo"
	placeUsecase "2024_2_ThereWillBeName/internal/pkg/attractions/usecase"
	grpcCategories "2024_2_ThereWillBeName/internal/pkg/categories/delivery/grpc"
	genCategories "2024_2_ThereWillBeName/internal/pkg/categories/delivery/grpc/gen"
	categoriesRepo "2024_2_ThereWillBeName/internal/pkg/categories/repo"
	categoriesUsecase "2024_2_ThereWillBeName/internal/pkg/categories/usecase"
	grpcCities "2024_2_ThereWillBeName/internal/pkg/cities/delivery/grpc"
	genCities "2024_2_ThereWillBeName/internal/pkg/cities/delivery/grpc/gen"
	citiesRepo "2024_2_ThereWillBeName/internal/pkg/cities/repo"
	citiesUsecase "2024_2_ThereWillBeName/internal/pkg/cities/usecase"
	grpcReviews "2024_2_ThereWillBeName/internal/pkg/reviews/delivery/grpc"
	genReviews "2024_2_ThereWillBeName/internal/pkg/reviews/delivery/grpc/gen"
	reviewRepo "2024_2_ThereWillBeName/internal/pkg/reviews/repo"
	reviewUsecase "2024_2_ThereWillBeName/internal/pkg/reviews/usecase"
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
	flag.IntVar(&cfg.Port, "grpc-port", 50051, "gRPC server port")
	flag.StringVar(&cfg.ConnStr, "connStr", "host=tripdb port=5432 user=service password=test dbname=trip sslmode=disable", "PostgreSQL connection string")
	flag.Parse()

	db, err := sql.Open("postgres", cfg.ConnStr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	reviewsRepo := reviewRepo.NewReviewRepository(db)
	reviewUsecase := reviewUsecase.NewReviewsUsecase(reviewsRepo)
	placeRepo := placeRepo.NewPLaceRepository(db)
	placeUsecase := placeUsecase.NewPlaceUsecase(placeRepo)
	citiesRepo := citiesRepo.NewCitiesRepository(db)
	citiesUsecase := citiesUsecase.NewCitiesUsecase(citiesRepo)
	categoriesRepo := categoriesRepo.NewCategoriesRepo(db)
	categoriesUsecase := categoriesUsecase.NewCategoriesUsecase(categoriesRepo)

	grpcAttractionsServer := grpc.NewServer()

	attractionsHandler := grpcAttractions.NewGrpcAttractionsHandler(placeUsecase)
	genPlaces.RegisterAttractionsServer(grpcAttractionsServer, attractionsHandler)

	citiesHandler := grpcCities.NewGrpcCitiesHandler(citiesUsecase)
	genCities.RegisterCitiesServer(grpcAttractionsServer, citiesHandler)

	reviewsHandler := grpcReviews.NewGrpcReviewsHandler(reviewUsecase)
	genReviews.RegisterReviewsServer(grpcAttractionsServer, reviewsHandler)

	categoriesHandler := grpcCategories.NewGrpcCategoriesHandler(categoriesUsecase)
	genCategories.RegisterCategoriesServer(grpcAttractionsServer, categoriesHandler)

	reflection.Register(grpcAttractionsServer)

	go func() {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Printf("gRPC server listening on :%d", cfg.Port)
		if err := grpcAttractionsServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down gRPC server...")
	grpcAttractionsServer.GracefulStop()
	log.Println("gRPC server gracefully stopped")
}
