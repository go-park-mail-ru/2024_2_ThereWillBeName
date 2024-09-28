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
)

type application struct {
	config models.Config
	logger *log.Logger
}

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
	app := &application{
		logger: logger,
		config: cfg,
	}
	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.Use(middleware.CORSMiddleware)
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})
	healthcheck := r.PathPrefix("/healthcheck").Subrouter()
	healthcheck.HandleFunc("", app.HealthcheckHandler).Methods(http.MethodGet)
	placecheck := r.PathPrefix("/v1/places").Subrouter()
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

func (app *application) HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "STATUS: OK")
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		fmt.Errorf("ERROR: healthcheckHandler: %s\n", err)
	}
}
