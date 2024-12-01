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
	metricsMw "2024_2_ThereWillBeName/internal/pkg/metrics/middleware"
	grpcReviews "2024_2_ThereWillBeName/internal/pkg/reviews/delivery/grpc"
	genReviews "2024_2_ThereWillBeName/internal/pkg/reviews/delivery/grpc/gen"
	reviewRepo "2024_2_ThereWillBeName/internal/pkg/reviews/repo"
	reviewUsecase "2024_2_ThereWillBeName/internal/pkg/reviews/usecase"
	grpcSearch "2024_2_ThereWillBeName/internal/pkg/search/delivery/grpc"
	genSearch "2024_2_ThereWillBeName/internal/pkg/search/delivery/grpc/gen"
	searchRepo "2024_2_ThereWillBeName/internal/pkg/search/repo"
	searchUsecase "2024_2_ThereWillBeName/internal/pkg/search/usecase"
	"database/sql"
	"fmt"
	"errors"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
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

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	cfg := config.Load()

	logger := setupLogger()

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Database.DbHost, cfg.Database.DbPort, cfg.Database.DbUser, cfg.Database.DbPass, cfg.Database.DbName))
	metricMw := metricsMw.Create()

	db, err := sql.Open("postgres", cfg.ConnStr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	wrappedDB := dblogger.NewDB(db, logger)

	reviewsRepo := reviewRepo.NewReviewRepository(wrappedDB)
	r := mux.NewRouter()
	r.Handle("/metrics", promhttp.Handler())
	httpSrv := &http.Server{
		Addr:              ":8091",
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	go func() {
		logger.Info("Starting HTTP server for metrics on :8091")
		if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error(fmt.Sprintf("HTTP server listen: %s\n", err))
		}
	}()

	reviewsRepo := reviewRepo.NewReviewRepository(db)
	reviewUsecase := reviewUsecase.NewReviewsUsecase(reviewsRepo)
	placeRepo := placeRepo.NewPLaceRepository(wrappedDB)
	placeUsecase := placeUsecase.NewPlaceUsecase(placeRepo)
	citiesRepo := citiesRepo.NewCitiesRepository(wrappedDB)
	citiesUsecase := citiesUsecase.NewCitiesUsecase(citiesRepo)
	categoriesRepo := categoriesRepo.NewCategoriesRepo(wrappedDB)
	categoriesUsecase := categoriesUsecase.NewCategoriesUsecase(categoriesRepo)
	searchRepo := searchRepo.NewSearchRepository(wrappedDB)
	searchUsecase := searchUsecase.NewSearchUsecase(searchRepo)

	attractionsHandler := grpcAttractions.NewGrpcAttractionsHandler(placeUsecase)
	citiesHandler := grpcCities.NewGrpcCitiesHandler(citiesUsecase)
	reviewsHandler := grpcReviews.NewGrpcReviewsHandler(reviewUsecase)
	categoriesHandler := grpcCategories.NewGrpcCategoriesHandler(categoriesUsecase)
	searchHandler := grpcSearch.NewGrpcSearchHandler(searchUsecase, logger)

	grpcAttractionsServer := grpc.NewServer(grpc.UnaryInterceptor(metricMw.ServerMetricsInterceptor))

	genPlaces.RegisterAttractionsServer(grpcAttractionsServer, attractionsHandler)
	genCities.RegisterCitiesServer(grpcAttractionsServer, citiesHandler)
	genReviews.RegisterReviewsServer(grpcAttractionsServer, reviewsHandler)
	genCategories.RegisterCategoriesServer(grpcAttractionsServer, categoriesHandler)
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

	go func() {
		ticker := time.NewTicker(15 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				metricMw.TrackSystemMetrics("attractions")
			case <-stop:
				return
			}
		}
	}()

	stop1 := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-stop1

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
