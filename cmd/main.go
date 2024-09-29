package main

import (
	httpHandler "2024_2_ThereWillBeName/internal/pkg/auth/delivery/http"
	"2024_2_ThereWillBeName/internal/pkg/auth/repo"
	"2024_2_ThereWillBeName/internal/pkg/auth/usecase"
	"2024_2_ThereWillBeName/internal/pkg/jwt"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"

	"encoding/hex"
	"math/rand"

	"github.com/gorilla/mux"
)

type config struct {
	port    int
	env     string
	connStr string
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment")
	flag.StringVar(&cfg.connStr, "connStr", "host=localhost port=5433 user=test_user password=1234567890 dbname=testdb_tripadvisor sslmode=disable", "PostgreSQL connection string")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg.connStr)
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

	logger.Println("Successfully connected to the database!")

	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})
	r.HandleFunc("/healthcheck", healthcheckHandler).Methods(http.MethodGet)
	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signup", h.SignUp).Methods(http.MethodPost)
	auth.HandleFunc("/login", h.Login).Methods(http.MethodPost)
	auth.HandleFunc("/logout", h.Logout).Methods(http.MethodPost)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "STATUS: OK")
	if err != nil {
		fmt.Printf("ERROR: healthcheckHandler: %s\n", err)
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
