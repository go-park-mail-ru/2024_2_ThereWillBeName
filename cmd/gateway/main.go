package main

import (
	"2024_2_ThereWillBeName/internal/models"
	genAttractions "2024_2_ThereWillBeName/internal/pkg/attractions/delivery/grpc/gen"
	httpPlaces "2024_2_ThereWillBeName/internal/pkg/attractions/delivery/http"
	genCategories "2024_2_ThereWillBeName/internal/pkg/categories/delivery/grpc/gen"
	httpCategories "2024_2_ThereWillBeName/internal/pkg/categories/delivery/http"
	genCities "2024_2_ThereWillBeName/internal/pkg/cities/delivery/grpc/gen"
	httpCities "2024_2_ThereWillBeName/internal/pkg/cities/delivery/http"
	"2024_2_ThereWillBeName/internal/pkg/httpresponses"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"2024_2_ThereWillBeName/internal/pkg/jwt"
	"2024_2_ThereWillBeName/internal/pkg/logger"
	"2024_2_ThereWillBeName/internal/pkg/middleware"
	genReviews "2024_2_ThereWillBeName/internal/pkg/reviews/delivery/grpc/gen"
	httpReviews "2024_2_ThereWillBeName/internal/pkg/reviews/delivery/http"
	genSearch "2024_2_ThereWillBeName/internal/pkg/search/delivery/grpc/gen"
	httpSearch "2024_2_ThereWillBeName/internal/pkg/search/delivery/http"
	genSurvey "2024_2_ThereWillBeName/internal/pkg/survey/delivery/grpc/gen"
	httpSurvey "2024_2_ThereWillBeName/internal/pkg/survey/delivery/http"
	genTrips "2024_2_ThereWillBeName/internal/pkg/trips/delivery/grpc/gen"
	httpTrips "2024_2_ThereWillBeName/internal/pkg/trips/delivery/http"
	genUsers "2024_2_ThereWillBeName/internal/pkg/user/delivery/grpc/gen"
	httpUsers "2024_2_ThereWillBeName/internal/pkg/user/delivery/http"
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	attractionsConn, err := grpc.Dial("attractions:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to attractions service: %v", err)
	}
	defer attractionsConn.Close()

	attractionsClient := genAttractions.NewAttractionsClient(attractionsConn)
	categoriesClient := genCategories.NewCategoriesClient(attractionsConn)
	citiesClient := genCities.NewCitiesClient(attractionsConn)
	reviewsClient := genReviews.NewReviewsClient(attractionsConn)
	searchClient := genSearch.NewSearchClient(attractionsConn)

	usersConn, err := grpc.NewClient("users:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to users service: %v", err)
	}
	defer usersConn.Close()
	usersClient := genUsers.NewUserServiceClient(usersConn)

	tripsConn, err := grpc.NewClient("trips:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to trips service: %v", err)
	}
	defer tripsConn.Close()
	tripsClient := genTrips.NewTripsClient(tripsConn)

	surveyConn, err := grpc.NewClient("survey:50054", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to survey service: %v", err)
	}
	defer surveyConn.Close()
	surveyClient := genSurvey.NewSurveyServiceClient(surveyConn)

	// Инициализация HTTP сервера
	corsMiddleware := middleware.NewCORSMiddleware([]string{cfg.AllowedOrigin})
	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	r.Use(corsMiddleware.CorsMiddleware)

	r.Use(middleware.RequestLoggerMiddleware(logger))

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
	places.HandleFunc("/search", placesHandler.SearchPlacesHandler).Methods(http.MethodGet)
	places.HandleFunc("/{id}", placesHandler.GetPlaceHandler).Methods(http.MethodGet)
	places.HandleFunc("/category/{categoryName}", placesHandler.GetPlacesByCategoryHandler).Methods(http.MethodGet)

	categoriesHandler := httpCategories.NewCategoriesHandler(categoriesClient, logger)
	categories := r.PathPrefix("/categories").Subrouter()
	categories.HandleFunc("", categoriesHandler.GetCategoriesHandler).Methods(http.MethodGet)

	reviewsHandler := httpReviews.NewReviewHandler(reviewsClient, logger)
	reviews := places.PathPrefix("/{placeID}/reviews").Subrouter()
	reviews.Handle("", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(reviewsHandler.CreateReviewHandler), logger)).Methods(http.MethodPost)
	reviews.Handle("/{reviewID}", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(reviewsHandler.UpdateReviewHandler), logger)).Methods(http.MethodPut)
	reviews.Handle("/{reviewID}", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(reviewsHandler.DeleteReviewHandler), logger)).Methods(http.MethodDelete)
	reviews.HandleFunc("/{reviewID}", reviewsHandler.GetReviewHandler).Methods(http.MethodGet)
	reviews.HandleFunc("", reviewsHandler.GetReviewsByPlaceIDHandler).Methods(http.MethodGet)

	citiesHandler := httpCities.NewCitiesHandler(citiesClient, logger)
	cities := r.PathPrefix("/cities").Subrouter()
	cities.HandleFunc("/search", citiesHandler.SearchCitiesByNameHandler).Methods(http.MethodGet)
	cities.HandleFunc("/{id}", citiesHandler.SearchCityByIDHandler).Methods(http.MethodGet)

	searchHandler := httpSearch.NewSearchHandler(searchClient, logger)
	search := r.PathPrefix("/search").Subrouter()
	search.HandleFunc("", searchHandler.Search).Methods(http.MethodGet)

	//Маршруты для Users
	usersHandler := httpUsers.NewUserHandler(usersClient, jwtHandler, logger)
	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signup", usersHandler.SignUp).Methods(http.MethodPost)
	auth.HandleFunc("/login", usersHandler.Login).Methods(http.MethodPost)
	auth.Handle("/logout", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(usersHandler.Logout), logger)).Methods(http.MethodPost)

	users := r.PathPrefix("/users").Subrouter()
	users.Handle("/me", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(usersHandler.CurrentUser), logger)).Methods(http.MethodGet)

	user := users.PathPrefix("/{userID}").Subrouter()

	user.Handle("/avatars", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(usersHandler.UploadAvatar), logger)).Methods(http.MethodPut)
	user.Handle("/profile", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(usersHandler.GetProfile), logger)).Methods(http.MethodGet)
	user.Handle("/update/password", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(usersHandler.UpdatePassword), logger)).Methods(http.MethodPut)
	user.Handle("/profile", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(usersHandler.UpdateProfile), logger)).Methods(http.MethodPut)

	tripsHandler := httpTrips.NewTripHandler(tripsClient, logger)
	trips := r.PathPrefix("/trips").Subrouter()
	trips.Handle("/{id}", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(tripsHandler.GetTripHandler), logger)).Methods(http.MethodGet)
	trips.Handle("", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(tripsHandler.CreateTripHandler), logger)).Methods(http.MethodPost)
	trips.Handle("/{id}", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(tripsHandler.UpdateTripHandler), logger)).Methods(http.MethodPut)
	trips.Handle("/{id}", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(tripsHandler.DeleteTripHandler), logger)).Methods(http.MethodDelete)
	trips.Handle("/{id}", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(tripsHandler.AddPlaceToTripHandler), logger)).Methods(http.MethodPost)
	user.Handle("/trips", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(tripsHandler.GetTripsByUserIDHandler), logger)).Methods(http.MethodGet)
	trips.Handle("/{id}/photos", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(tripsHandler.AddPhotosToTripHandler), logger)).Methods(http.MethodPut)
	trips.Handle("/{id}/photos", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(tripsHandler.DeletePhotoHandler), logger)).Methods(http.MethodDelete)

	surveyHandler := httpSurvey.NewSurveyHandler(surveyClient, logger)
	survey := r.PathPrefix("/survey").Subrouter()
	survey.Handle("/stats/{id}", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(surveyHandler.GetSurveyStatsBySurveyId), logger)).Methods(http.MethodGet)
	survey.Handle("/{id}", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(surveyHandler.GetSurveyById), logger)).Methods(http.MethodGet)
	survey.Handle("/{id}", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(surveyHandler.CreateSurveyResponse), logger)).Methods(http.MethodPost)
	survey.Handle("/users/{id}", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(surveyHandler.GetSurveyStatsByUserId), logger)).Methods(http.MethodGet)

	httpSrv := &http.Server{Handler: r, Addr: fmt.Sprintf(":%d", cfg.Port)}
	go func() {
		logger.Info("HTTP server listening on :%d", cfg.Port)
		if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("failed to serve HTTP: %d", err)
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
