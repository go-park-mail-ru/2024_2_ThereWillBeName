package main

import (
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
	"2024_2_ThereWillBeName/internal/pkg/config"
	"2024_2_ThereWillBeName/internal/pkg/dblogger"
	"2024_2_ThereWillBeName/internal/pkg/logger"
	grpcReviews "2024_2_ThereWillBeName/internal/pkg/reviews/delivery/grpc"
	genReviews "2024_2_ThereWillBeName/internal/pkg/reviews/delivery/grpc/gen"
	reviewRepo "2024_2_ThereWillBeName/internal/pkg/reviews/repo"
	reviewUsecase "2024_2_ThereWillBeName/internal/pkg/reviews/usecase"
	grpcSearch "2024_2_ThereWillBeName/internal/pkg/search/delivery/grpc"
	genSearch "2024_2_ThereWillBeName/internal/pkg/search/delivery/grpc/gen"
	searchRepo "2024_2_ThereWillBeName/internal/pkg/search/repo"
	searchUsecase "2024_2_ThereWillBeName/internal/pkg/search/usecase"
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

	reviewsRepo := reviewRepo.NewReviewRepository(wrappedDB)
	reviewUsecase := reviewUsecase.NewReviewsUsecase(reviewsRepo)
	placeRepo := placeRepo.NewPLaceRepository(wrappedDB)
	placeUsecase := placeUsecase.NewPlaceUsecase(placeRepo)
	citiesRepo := citiesRepo.NewCitiesRepository(wrappedDB)
	citiesUsecase := citiesUsecase.NewCitiesUsecase(citiesRepo)
	categoriesRepo := categoriesRepo.NewCategoriesRepo(wrappedDB)
	categoriesUsecase := categoriesUsecase.NewCategoriesUsecase(categoriesRepo)
	searchRepo := searchRepo.NewSearchRepository(wrappedDB)
	searchUsecase := searchUsecase.NewSearchUsecase(searchRepo)

	grpcAttractionsServer := grpc.NewServer()

	attractionsHandler := grpcAttractions.NewGrpcAttractionsHandler(placeUsecase)
	genPlaces.RegisterAttractionsServer(grpcAttractionsServer, attractionsHandler)

	citiesHandler := grpcCities.NewGrpcCitiesHandler(citiesUsecase)
	genCities.RegisterCitiesServer(grpcAttractionsServer, citiesHandler)

	reviewsHandler := grpcReviews.NewGrpcReviewsHandler(reviewUsecase)
	genReviews.RegisterReviewsServer(grpcAttractionsServer, reviewsHandler)

	categoriesHandler := grpcCategories.NewGrpcCategoriesHandler(categoriesUsecase)
	genCategories.RegisterCategoriesServer(grpcAttractionsServer, categoriesHandler)

	searchHandler := grpcSearch.NewGrpcSearchHandler(searchUsecase, logger)
	genSearch.RegisterSearchServer(grpcAttractionsServer, searchHandler)

	reflection.Register(grpcAttractionsServer)

	go func() {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Grpc.AttractionPort))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Printf("gRPC server listening on :%d", cfg.Grpc.AttractionPort)
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
