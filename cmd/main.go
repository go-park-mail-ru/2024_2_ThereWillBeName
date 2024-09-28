package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	httpHandler "2024_2_ThereWillBeName/internal/pkg/auth/delivery/http"
	"2024_2_ThereWillBeName/internal/pkg/jwt"
	"2024_2_ThereWillBeName/internal/pkg/auth/usecase"
	"2024_2_ThereWillBeName/internal/pkg/auth/repo"
	"2024_2_ThereWillBeName/internal/pkg/auth"

	"math/rand"
	"encoding/hex"
)

type config struct {
	port int
	env  string
}
type application struct {
	config config
	logger *log.Logger
	db *sql.DB
	jwtSecret []byte
	authUseCase auth.AuthUsecase
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB()
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	jwtSecret, err := generateSecretKey(32)
	if err != nil {
		logger.Fatal("Error generating secret key:", err)
	}

	authRepo := repo.NewRepository(db) 
	jwtHandler := jwt.NewJWT(string(jwtSecret))
	authUseCase := usecase.NewAuthUsecase(authRepo, jwtHandler)
    h := httpHandler.NewHandler(authUseCase, jwtHandler)

	app := &application{
		config: cfg,
		logger: logger,
		db: db,
		jwtSecret: []byte(jwtSecret),
		authUseCase: authUseCase,
	}

	logger.Println("Successfully connected to the database!")

	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", app.healthcheckHandler)
	mux.HandleFunc("/signup", h.SignUp)
    mux.HandleFunc("/login", h.Login) 

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
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

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "STATUS: OK")
	if err != nil {
		fmt.Printf("ERROR: healthcheckHandler: %s\n", err)
	}
}

func openDB() (*sql.DB, error) {
	connStr := "host=localhost port=5433 user=test_user password=1234567890 dbname=testdb_tripadvisor sslmode=disable"
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

