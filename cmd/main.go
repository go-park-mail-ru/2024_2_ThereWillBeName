package main

import (
	"2024_2_ThereWillBeName/internal/models"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"2024_2_ThereWillBeName/internal/pkg/jwt"
	"2024_2_ThereWillBeName/internal/pkg/logger"
	"2024_2_ThereWillBeName/internal/pkg/middleware"
	httpHandler "2024_2_ThereWillBeName/internal/pkg/user/delivery/http"
	userRepo "2024_2_ThereWillBeName/internal/pkg/user/repo"
	userUsecase "2024_2_ThereWillBeName/internal/pkg/user/usecase"
	"log/slog"
	"strconv"

	citieshandler "2024_2_ThereWillBeName/internal/pkg/cities/delivery/http"
	citiesrepo "2024_2_ThereWillBeName/internal/pkg/cities/repo"
	citiesusecase "2024_2_ThereWillBeName/internal/pkg/cities/usecase"
	delivery "2024_2_ThereWillBeName/internal/pkg/places/delivery/http"
	placeRepo "2024_2_ThereWillBeName/internal/pkg/places/repo"
	placeUsecase "2024_2_ThereWillBeName/internal/pkg/places/usecase"
	reviewhandler "2024_2_ThereWillBeName/internal/pkg/reviews/delivery/http"
	reviewrepo "2024_2_ThereWillBeName/internal/pkg/reviews/repo"
	reviewusecase "2024_2_ThereWillBeName/internal/pkg/reviews/usecase"
	triphandler "2024_2_ThereWillBeName/internal/pkg/trips/delivery/http"
	triprepo "2024_2_ThereWillBeName/internal/pkg/trips/repo"
	tripusecase "2024_2_ThereWillBeName/internal/pkg/trips/usecase"

	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "2024_2_ThereWillBeName/docs"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	var cfg models.Config
	flag.IntVar(&cfg.Port, "port", 8080, "API server port")
	flag.StringVar(&cfg.Env, "env", "production", "Environment")
	flag.StringVar(&cfg.AllowedOrigin, "allowed-origin", "*", "Allowed origin")
	flag.StringVar(&cfg.ConnStr, "connStr", "host=tripdb port=5432 user=service password=test dbname=trip sslmode=disable", "PostgreSQL connection string")
	flag.Parse()

	logger := setupLogger()
	defer logger.Info("Server stopped")

	db, err := openDB(cfg.ConnStr)
	if err != nil {
		logger.Error("Failed to open database", slog.Any("error", err))
	}
	logger.Info("Connected to database successfully")
	defer db.Close()

	jwtSecret := os.Getenv("JWT_SECRET")
	storagePath := os.Getenv("AVATAR_STORAGE_PATH")

	userRepo := userRepo.NewAuthRepository(db)
	jwtHandler := jwt.NewJWT(string(jwtSecret), logger)
	userUseCase := userUsecase.NewUserUsecase(userRepo, storagePath)
	h := httpHandler.NewUserHandler(userUseCase, jwtHandler, logger)

	reviewsRepo := reviewrepo.NewReviewRepository(db)
	reviewUsecase := reviewusecase.NewReviewsUsecase(reviewsRepo)
	reviewHandler := reviewhandler.NewReviewHandler(reviewUsecase, logger)
	placeRepo := placeRepo.NewPLaceRepository(db)
	placeUsecase := placeUsecase.NewPlaceUsecase(placeRepo)
	placeHandler := delivery.NewPlacesHandler(placeUsecase, logger)
	tripsRepo := triprepo.NewTripRepository(db)
	tripUsecase := tripusecase.NewTripsUsecase(tripsRepo)
	tripHandler := triphandler.NewTripHandler(tripUsecase, logger)
	citiesRepo := citiesrepo.NewCitiesRepository(db)
	citiesUsecase := citiesusecase.NewCitiesUsecase(citiesRepo)
	citiesHandler := citieshandler.NewCitiesHandler(citiesUsecase, logger)

	corsMiddleware := middleware.NewCORSMiddleware([]string{cfg.AllowedOrigin})

	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	r.Use(corsMiddleware.CorsMiddleware)
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := httpresponse.ErrorResponse{
			Message: "Not found",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusNotFound, logger)
	})
	r.HandleFunc("/healthcheck", healthcheckHandler).Methods(http.MethodGet)

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signup", h.SignUp).Methods(http.MethodPost)
	auth.HandleFunc("/login", h.Login).Methods(http.MethodPost)
	auth.HandleFunc("/logout", h.Logout).Methods(http.MethodPost)
	users := r.PathPrefix("/users").Subrouter()
	users.Handle("/me", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(h.CurrentUser), logger)).Methods(http.MethodGet)

	user := users.PathPrefix("/{userID}").Subrouter()

	user.Handle("/avatars", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(h.UploadAvatar), logger)).Methods(http.MethodPut)
	user.Handle("/profile", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(h.GetProfile), logger)).Methods(http.MethodGet)

	places := r.PathPrefix("/places").Subrouter()
	places.HandleFunc("", placeHandler.GetPlacesHandler).Methods(http.MethodGet)
	places.HandleFunc("", placeHandler.PostPlaceHandler).Methods(http.MethodPost)
	places.HandleFunc("/search/{placeName}", placeHandler.SearchPlacesHandler).Methods(http.MethodGet)
	places.HandleFunc("/{id}", placeHandler.GetPlaceHandler).Methods(http.MethodGet)
	places.HandleFunc("/{id}", placeHandler.PutPlaceHandler).Methods(http.MethodPut)
	places.HandleFunc("/{id}", placeHandler.DeletePlaceHandler).Methods(http.MethodDelete)

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	reviews := places.PathPrefix("/{placeID}/reviews").Subrouter()
	reviews.HandleFunc("", reviewHandler.CreateReviewHandler).Methods(http.MethodPost)
	reviews.HandleFunc("/{reviewID}", reviewHandler.UpdateReviewHandler).Methods(http.MethodPut)
	reviews.HandleFunc("/{reviewID}", reviewHandler.DeleteReviewHandler).Methods(http.MethodDelete)
	reviews.HandleFunc("/{reviewID}", reviewHandler.GetReviewHandler).Methods(http.MethodGet)
	reviews.HandleFunc("", reviewHandler.GetReviewsByPlaceIDHandler).Methods(http.MethodGet)

	trips := r.PathPrefix("/trips").Subrouter()
	trips.HandleFunc("", tripHandler.CreateTripHandler).Methods(http.MethodPost)
	trips.HandleFunc("/{id}", tripHandler.UpdateTripHandler).Methods(http.MethodPut)
	trips.HandleFunc("/{id}", tripHandler.DeleteTripHandler).Methods(http.MethodDelete)
	trips.HandleFunc("/{id}", tripHandler.GetTripHandler).Methods(http.MethodGet)
	user.HandleFunc("/trips", tripHandler.GetTripsByUserIDHandler).Methods(http.MethodGet)

	cities := r.PathPrefix("/cities").Subrouter()
	cities.HandleFunc("/search", citiesHandler.SearchCitiesByNameHandler).Methods(http.MethodGet)
	cities.HandleFunc("/{id}", citiesHandler.SearchCityByIDHandler).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Info("starting server", "environment", cfg.Env, "address", srv.Addr)
	err = srv.ListenAndServe()
	if err != nil {
		logger.Error("Failed to start server", slog.Any("error", err))
		os.Exit(1)
	}
}

// healthcheckHandler godoc
// @Summary Health check
// @Description Check the health status of the service
// @Produce text/plain
// @Success 200 {string} string "STATUS: OK"
// @Failure 400 {object} httpresponses.ErrorResponse "Bad Request"
// @Router /healthcheck [get]
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

func openDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
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
