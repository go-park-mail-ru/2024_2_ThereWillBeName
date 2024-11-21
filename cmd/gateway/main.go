package main

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/attractions/delivery/grpc/gen"
	httpPlaces "2024_2_ThereWillBeName/internal/pkg/attractions/delivery/http"
	httpCategories "2024_2_ThereWillBeName/internal/pkg/categories/delivery/http"
	httpCities "2024_2_ThereWillBeName/internal/pkg/cities/delivery/http"
	"2024_2_ThereWillBeName/internal/pkg/httpresponses"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"2024_2_ThereWillBeName/internal/pkg/jwt"
	"2024_2_ThereWillBeName/internal/pkg/logger"
	"2024_2_ThereWillBeName/internal/pkg/middleware"
	httpReviews "2024_2_ThereWillBeName/internal/pkg/reviews/delivery/http"
	httpTrips "2024_2_ThereWillBeName/internal/pkg/trips/delivery/http"
	httpUsers "2024_2_ThereWillBeName/internal/pkg/user/delivery/http"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	var cfg models.Config
	flag.IntVar(&cfg.Port, "port", 8080, "API server port")
	flag.StringVar(&cfg.Env, "env", "production", "Environment")
	flag.StringVar(&cfg.AllowedOrigin, "allowed-origin", "*", "Allowed origin")
	flag.Parse()

	logger := setupLogger()

	jwtSecret := os.Getenv("JWT_SECRET")
	jwtHandler := jwt.NewJWT(jwtSecret, logger)

	attractionsConn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to attractions service: %v", err)
	}
	defer attractionsConn.Close()
	attractionsClient := gen.NewAttractionsClient(attractionsConn)

	usersConn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to users service: %v", err)
	}
	defer usersConn.Close()
	usersClient := gen.NewUsersClient(usersConn)

	tripsConn, err := grpc.NewClient("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to trips service: %v", err)
	}
	defer tripsConn.Close()
	tripsClient := gen.NewTripsClient(tripsConn)

	// Инициализация HTTP сервера
	corsMiddleware := middleware.NewCORSMiddleware([]string{cfg.AllowedOrigin})
	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	r.Use(corsMiddleware.CorsMiddleware)

	// Обработка ненайденных маршрутов
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := httpresponses.ErrorResponse{
			Message: "Not found",
		}
		httpresponses.SendJSONResponse(w, response, http.StatusNotFound, logger)
	})

	// Маршрут для healthcheck
	r.HandleFunc("/healthcheck", healthcheckHandler).Methods(http.MethodGet)

	// Маршруты для attractions
	placesHandler := httpPlaces.NewPlacesHandler(attractionsClient, logger)
	places := r.PathPrefix("/places").Subrouter()
	places.HandleFunc("", placesHandler.GetPlacesHandler).Methods(http.MethodGet)
	places.HandleFunc("/search/{placeName}", placesHandler.SearchPlacesHandler).Methods(http.MethodGet)
	places.HandleFunc("/{id}", placesHandler.GetPlaceHandler).Methods(http.MethodGet)
	places.HandleFunc("/category/{categoryName}", placesHandler.GetPlacesByCategoryHandler).Methods(http.MethodGet)

	categoriesHandler := httpCategories.NewCategoriesHandler(attractionsClient, logger)
	categories := r.PathPrefix("/categories").Subrouter()
	categories.HandleFunc("", categoriesHandler.GetCategoriesHandler).Methods(http.MethodGet)

	reviewsHandler := httpReviews.NewReviewHandler(attractionsClient, logger)
	reviews := places.PathPrefix("/{placeID}/reviews").Subrouter()
	reviews.Handle("", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(reviewsHandler.CreateReviewHandler), logger)).Methods(http.MethodPost)
	reviews.Handle("/reviewID", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(reviewsHandler.UpdateReviewHandler), logger)).Methods(http.MethodPut)
	reviews.Handle("/reviewID", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(reviewsHandler.DeleteReviewHandler), logger)).Methods(http.MethodDelete)
	reviews.HandleFunc("/{reviewID}", reviewsHandler.GetReviewHandler).Methods(http.MethodGet)
	reviews.HandleFunc("", reviewsHandler.GetReviewsByPlaceIDHandler).Methods(http.MethodGet)

	citiesHandler := httpCities.NewCitiesHandler(attractionsClient, logger)
	cities := r.PathPrefix("/cities").Subrouter()
	cities.HandleFunc("/search", citiesHandler.SearchCitiesByNameHandler).Methods(http.MethodGet)
	cities.HandleFunc("/{id}", citiesHandler.SearchCityByIDHandler).Methods(http.MethodGet)

	usersHandler := httpUsers.NewUsersHandler(usersClient, logger)
	users := r.PathPrefix("/users").Subrouter()
	users.HandleFunc("", usersHandler.GetUsersHandler).Methods(http.MethodGet)
	users.HandleFunc("/{id}", usersHandler.GetUserHandler).Methods(http.MethodGet)

	tripsHandler := httpTrips.NewTripsHandler(tripsClient, logger)
	trips := r.PathPrefix("/trips").Subrouter()
	trips.HandleFunc("", tripsHandler.GetTripsHandler).Methods(http.MethodGet)
	trips.HandleFunc("/{id}", tripsHandler.GetTripHandler).Methods(http.MethodGet)

	httpSrv := &http.Server{Handler: r, Addr: fmt.Sprintf(":%d", cfg.Port)}
	go func() {
		logger.Info("HTTP server listening on :%d", cfg.Port)
		if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("failed to serve HTTP: %v", err)
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	logger.Info("Shutting down HTTP server...")
	if err := httpSrv.Shutdown(context.Background()); err != nil {
		logger.Error("HTTP server shutdown failed: %v", err)
		os.Exit(1)
	}
	logger.Info("HTTP server gracefully stopped")
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	logger := setupLogger()

	_, err := fmt.Fprintf(w, "STATUS: OK")
	if err != nil {
		logger.Error("Failed to write healthcheck response", slog.Any("error", err))
		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, logger)
	}
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
