package main

import (
	"2024_2_ThereWillBeName/internal/models"
	httpHandler "2024_2_ThereWillBeName/internal/pkg/auth/delivery/http"
	"2024_2_ThereWillBeName/internal/pkg/auth/repo"
	"2024_2_ThereWillBeName/internal/pkg/auth/usecase"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"2024_2_ThereWillBeName/internal/pkg/jwt"
	"2024_2_ThereWillBeName/internal/pkg/middleware"

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
	"log"
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
	flag.StringVar(&cfg.Env, "env", "development", "Environment")
	flag.StringVar(&cfg.AllowedOrigin, "allowed-origin", "*", "Allowed origin")
	flag.StringVar(&cfg.ConnStr, "connStr", "host=tripdb port=5432 user=service password=test dbname=trip sslmode=disable", "PostgreSQL connection string")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg.ConnStr)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	jwtSecret := os.Getenv("JWT_SECRET")

	if err != nil {
		logger.Fatal("Error generating secret key:", err)
	}

	authRepo := repo.NewAuthRepository(db)
	jwtHandler := jwt.NewJWT(string(jwtSecret))
	authUseCase := usecase.NewAuthUsecase(authRepo, jwtHandler)
	h := httpHandler.NewAuthHandler(authUseCase, jwtHandler)

	reviewsRepo := reviewrepo.NewReviewRepository(db)
	reviewUsecase := reviewusecase.NewReviewsUsecase(reviewsRepo)
	reviewHandler := reviewhandler.NewReviewHandler(reviewUsecase)
	placeRepo := placeRepo.NewPLaceRepository(db)
	placeUsecase := placeUsecase.NewPlaceUsecase(placeRepo)
	placeHandler := delivery.NewPlacesHandler(placeUsecase)
	tripsRepo := triprepo.NewTripRepository(db)
	tripUsecase := tripusecase.NewTripsUsecase(tripsRepo)
	tripHandler := triphandler.NewTripHandler(tripUsecase)

	corsMiddleware := middleware.NewCORSMiddleware([]string{cfg.AllowedOrigin})

	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	r.Use(corsMiddleware.CorsMiddleware)
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := httpresponse.ErrorResponse{
			Message: "Not found",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusNotFound)
	})
	r.HandleFunc("/healthcheck", healthcheckHandler).Methods(http.MethodGet)

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signup", h.SignUp).Methods(http.MethodPost)
	auth.HandleFunc("/login", h.Login).Methods(http.MethodPost)
	auth.HandleFunc("/logout", h.Logout).Methods(http.MethodPost)
	users := r.PathPrefix("/users").Subrouter()
	users.Handle("/me", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(h.CurrentUser))).Methods(http.MethodGet)

	user := users.PathPrefix("/{userID}").Subrouter()

	places := r.PathPrefix("/places").Subrouter()
	places.HandleFunc("", placeHandler.GetPlacesHandler).Methods(http.MethodGet)
	// places.HandleFunc("", placeHandler.PostPlaceHandler).Methods(http.MethodPost)
	places.Handle("", middleware.CSRFMiddleware(http.HandlerFunc(placeHandler.PostPlaceHandler))).Methods(http.MethodPost)
	places.HandleFunc("/search/{placeName}", placeHandler.SearchPlacesHandler).Methods(http.MethodGet)
	places.HandleFunc("/{id}", placeHandler.GetPlaceHandler).Methods(http.MethodGet)
	// places.HandleFunc("/{id}", placeHandler.PutPlaceHandler).Methods(http.MethodPut)
	places.Handle("/{id}", middleware.CSRFMiddleware(http.HandlerFunc(placeHandler.PutPlaceHandler))).Methods(http.MethodPut)
	// places.HandleFunc("/{id}", placeHandler.DeletePlaceHandler).Methods(http.MethodDelete)
	places.Handle("/{id}", middleware.CSRFMiddleware(http.HandlerFunc(placeHandler.DeletePlaceHandler))).Methods(http.MethodDelete)

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	reviews := r.PathPrefix("/reviews").Subrouter()
	// reviews.HandleFunc("/", reviewHandler.CreateReviewHandler).Methods(http.MethodPost)
	reviews.Handle("/", middleware.CSRFMiddleware(http.HandlerFunc(reviewHandler.CreateReviewHandler))).Methods(http.MethodPost)
	// reviews.HandleFunc("/{id}", reviewHandler.UpdateReviewHandler).Methods(http.MethodPut)
	reviews.Handle("/{id}", middleware.CSRFMiddleware(http.HandlerFunc(reviewHandler.UpdateReviewHandler))).Methods(http.MethodPut)
	// reviews.HandleFunc("/{id}", reviewHandler.DeleteReviewHandler).Methods(http.MethodDelete)
	reviews.Handle("/{id}", middleware.CSRFMiddleware(http.HandlerFunc(reviewHandler.DeleteReviewHandler))).Methods(http.MethodDelete)
	reviews.HandleFunc("/{id}", reviewHandler.GetReviewHandler).Methods(http.MethodGet)
	reviews.HandleFunc("/reviews/{reviewID}", reviewHandler.GetReviewsByPlaceIDHandler).Methods(http.MethodGet)

	trips := r.PathPrefix("/trips").Subrouter()
	// trips.HandleFunc("", tripHandler.CreateTripHandler).Methods(http.MethodPost)
	trips.Handle("", middleware.CSRFMiddleware(http.HandlerFunc(tripHandler.CreateTripHandler))).Methods(http.MethodPost)
	// trips.HandleFunc("/{id}", tripHandler.UpdateTripHandler).Methods(http.MethodPut)
	trips.Handle("/{id}", middleware.CSRFMiddleware(http.HandlerFunc(tripHandler.UpdateTripHandler))).Methods(http.MethodPut)
	// trips.HandleFunc("/{id}", tripHandler.DeleteTripHandler).Methods(http.MethodDelete)
	trips.Handle("/{id}", middleware.CSRFMiddleware(http.HandlerFunc(tripHandler.DeleteTripHandler))).Methods(http.MethodDelete)
	trips.HandleFunc("/{id}", tripHandler.GetTripHandler).Methods(http.MethodGet)
	user.HandleFunc("/trips", tripHandler.GetTripsByUserIDHandler).Methods(http.MethodGet)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.Env, srv.Addr)
	err = srv.ListenAndServe()
	if err != nil {
		logger.Println(fmt.Errorf("Failed to start server: %v", err))
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
	_, err := fmt.Fprintf(w, "STATUS: OK")
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
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
