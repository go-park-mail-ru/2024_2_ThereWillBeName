package main

import (
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
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment")
	flag.Parse()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := &application{
		config: cfg,
		logger: logger,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", app.healthcheckHandler)
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
		logger.Fatal(err)
	}
}

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "STATUS: OK")
	if err != nil {
		fmt.Printf("ERROR: healthcheckHandler: %s\n", err)
	}
}
