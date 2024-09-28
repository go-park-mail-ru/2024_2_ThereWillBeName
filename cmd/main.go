package main

import (
	"TripAdvisor/internal/models"
	"TripAdvisor/internal/pkg/places/delivery"
	"TripAdvisor/internal/pkg/places/repo"
	"TripAdvisor/internal/pkg/places/usecase"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	//connStr := "user=postgres password=mypassword host=localhost port=5432 dbname=landmarks sslmode=disable"
	//db, err := sql.Open("postgres", connStr)
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()

	newPlaceRepo := repo.NewRepository()
	placeUsecase := usecase.NewPlaceUsecase(newPlaceRepo)
	handler := delivery.NewPlacesHandler(placeUsecase)

	var cfg models.Config
	flag.IntVar(&cfg.Port, "port", 8080, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", HealthcheckHandler)
	mux.HandleFunc("/api/v1/places", handler.GetPlaceHandler)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      mux,
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

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "STATUS: OK")
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		fmt.Errorf("ERROR: healthcheckHandler: %s\n", err)
	}
}
