package gateway

import (
	"2024_2_ThereWillBeName/internal/pkg/trips/delivery/grpc/gen"
	tripshandler "2024_2_ThereWillBeName/internal/pkg/trips/delivery/http"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {
	grpcConn, err := grpc.Dial("localhost:50051")
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer grpcConn.Close()

	tripsClient := gen.NewTripsClient(grpcConn)

	// Настройка HTTP маршрутов
	r := mux.NewRouter()
	r.HandleFunc("/trips", func(w http.ResponseWriter, r *http.Request) {
		tripshandler.CreateTripHandler
	}).Methods(http.MethodPost)
	r.HandleFunc("/trips/{id}", func(w http.ResponseWriter, r *http.Request) {
		getTripHandler(w, r, tripsClient)
	}).Methods(http.MethodGet)

	fmt.Println("Starting HTTP Gateway on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
