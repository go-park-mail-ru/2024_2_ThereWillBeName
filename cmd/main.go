package main

import (
	"TripAdvisor/internal/models"
	"TripAdvisor/internal/pkg/middleware"
	"TripAdvisor/internal/pkg/places/delivery"
	"TripAdvisor/internal/pkg/places/repo"
	"TripAdvisor/internal/pkg/places/usecase"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/gorilla/mux"
)

func main() {
	newPlaceRepo := repo.NewPLaceRepository()
	placeUsecase := usecase.NewPlaceUsecase(newPlaceRepo)
	handler := delivery.NewPlacesHandler(placeUsecase)

	var cfg models.Config
	flag.IntVar(&cfg.Port, "port", 8080, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment")
	flag.StringVar(&cfg.AllowedOrigin, "allowed-origin", "*", "Allowed origin")
	flag.Parse()

	corsMiddleware := middleware.NewCORSMiddleware([]string{cfg.AllowedOrigin})

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	r.Use(corsMiddleware.CorsMiddleware)
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})
	healthcheck := r.PathPrefix("/healthcheck").Subrouter()
	healthcheck.HandleFunc("", healthcheckHandler).Methods(http.MethodGet)
	placecheck := r.PathPrefix("/places").Subrouter()
	placecheck.HandleFunc("", handler.GetPlaceHandler).Methods(http.MethodGet)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.Env, srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Errorf("Failed to start server: %v", err)
		os.Exit(1)
	}
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "STATUS: OK")
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		fmt.Errorf("ERROR: healthcheckHandler: %s\n", err)
	}
}
