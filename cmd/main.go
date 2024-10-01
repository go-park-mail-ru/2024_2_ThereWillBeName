package main

import (
	"2024_2_ThereWillBeName/internal/models"
	httpHandler "2024_2_ThereWillBeName/internal/pkg/auth/delivery/http"
	"2024_2_ThereWillBeName/internal/pkg/auth/repo"
	"2024_2_ThereWillBeName/internal/pkg/auth/usecase"
	"2024_2_ThereWillBeName/internal/pkg/jwt"
	"2024_2_ThereWillBeName/internal/pkg/middleware"
	"2024_2_ThereWillBeName/internal/pkg/places/delivery"
	placerepo "2024_2_ThereWillBeName/internal/pkg/places/repo"
	placeusecase "2024_2_ThereWillBeName/internal/pkg/places/usecase"
	"database/sql"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	var cfg models.Config
	flag.IntVar(&cfg.Port, "port", 8080, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment")
	flag.StringVar(&cfg.AllowedOrigin, "allowed-origin", "*", "Allowed origin")
	flag.StringVar(&cfg.ConnStr, "connStr", "host=tripdb port=5432 user=service password=test dbname=trip sslmode=disable", "PostgreSQL connection string")
	flag.Parse()

	newPlaceRepo := placerepo.NewPLaceRepository()
	placeUsecase := placeusecase.NewPlaceUsecase(newPlaceRepo)
	handler := delivery.NewPlacesHandler(placeUsecase)

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg.ConnStr)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	jwtSecret, err := generateSecretKey(32)
	if err != nil {
		logger.Fatal("Error generating secret key:", err)
	}

	authRepo := repo.NewAuthRepository(db)
	jwtHandler := jwt.NewJWT(string(jwtSecret))
	authUseCase := usecase.NewAuthUsecase(authRepo, jwtHandler)
	h := httpHandler.NewAuthHandler(authUseCase, jwtHandler)

	corsMiddleware := middleware.NewCORSMiddleware([]string{cfg.AllowedOrigin})

	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	r.Use(corsMiddleware.CorsMiddleware)
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})
	r.HandleFunc("/healthcheck", healthcheckHandler).Methods(http.MethodGet)

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signup", h.SignUp).Methods(http.MethodPost)
	auth.HandleFunc("/login", h.Login).Methods(http.MethodPost)
	auth.HandleFunc("/logout", h.Logout).Methods(http.MethodPost)
	auth.Handle("/users/me", middleware.MiddlewareAuth(jwtHandler, http.HandlerFunc(h.CurrentUser))).Methods(http.MethodGet)
	places := r.PathPrefix("/places").Subrouter()
	places.HandleFunc("", handler.GetPlaceHandler).Methods(http.MethodGet)
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

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "STATUS: OK")
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
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

func generateSecretKey(length int) (string, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}
