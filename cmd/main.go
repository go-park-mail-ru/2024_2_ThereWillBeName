package main

import (
	"TripAdvisor/pkg/auth"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type config struct {
	port int
	env  string
}
type application struct {
	config      config
	logger      *log.Logger
	repousecase auth.PlaceUsecase
}

func main() {
	connStr := "user=postgres password=mypassword host=localhost port=5432 dbname=landmarks sslmode=disable"
	repos := auth.NewRepository(connStr)
	repoUsecase := auth.NewRepoUsecase(repos)
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment")
	flag.Parse()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := &application{
		config:      cfg,
		logger:      logger,
		repousecase: repoUsecase,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", app.healthcheckHandler)
	mux.HandleFunc("/place", app.getPlaceHandler)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Errorf("Failed to start server: %v", err)
		os.Exit(1)
	}
}

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "STATUS: OK")
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		fmt.Errorf("ERROR: healthcheckHandler: %s\n", err)
	}
}

func (app *application) getPlaceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	places, err := app.repousecase.GetPlace()
	if err != nil {
		http.Error(w, "Не удалось получить список достопримечательностей", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(places)
	if err != nil {
		http.Error(w, "Не удалось преобразовать в json", http.StatusInternalServerError)
		return
	}
}
