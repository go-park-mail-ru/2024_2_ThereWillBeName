package http

import (
	"2024_2_ThereWillBeName/internal/models"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"2024_2_ThereWillBeName/internal/pkg/trips"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TripHandler struct {
	uc trips.TripsUsecase
}

func NewTripHandler(uc trips.TripsUsecase) *TripHandler {
	return &TripHandler{uc}
}

// CreateTripHandler godoc
// @Summary Create a new trip
// @Description Create a new trip with given fields
// @Accept json
// @Produce json
// @Param trip body models.Trip true "Trip details"
// @Success 201 {object} models.Trip "Trip created successfully"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid request"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to create trip"
// @Router /trips/create [post]
func (h *TripHandler) CreateTripHandler(w http.ResponseWriter, r *http.Request) {
	var trip models.Trip
	err := json.NewDecoder(r.Body).Decode(&trip)
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	err = h.uc.CreateTrip(context.Background(), trip)
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Failed to create trip",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateTripHandler godoc
// @Summary Update an existing trip
// @Description Update trip details by trip ID
// @Accept json
// @Produce json
// @Param id path int true "Trip ID"
// @Param trip body models.Trip true "Updated trip details"
// @Success 200 {object} models.Trip "Trip updated successfully"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid request"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to update trip"
// @Router /trips/{id} [put]
func (h *TripHandler) UpdateTripHandler(w http.ResponseWriter, r *http.Request) {
	var trip models.Trip
	vars := mux.Vars(r)
	tripID, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&trip)
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}
	if tripID < 0 {
		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}
	trip.ID = uint(tripID)
	err = h.uc.UpdateTrip(context.Background(), trip)
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Failed to update trip",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteTripHandler godoc
// @Summary Delete a trip
// @Description Delete a trip by trip ID
// @Produce json
// @Param id path int true "Trip ID"
// @Success 204 "Trip deleted successfully"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid request"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to delete trip"
// @Router /trips/{id} [delete]
func (h *TripHandler) DeleteTripHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	err = h.uc.DeleteTrip(context.Background(), uint(id))
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Failed to delete trip",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ReadTripsByUserIDHandler godoc
// @Summary Retrieve trips by user ID
// @Description Get all trips for a specific user
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {array} models.Trip "List of trips"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid request"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to retrieve trips"
// @Router /trips/{userID} [get]
func (h *TripHandler) ReadTripsByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["userID"]
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	trips, err := h.uc.ReadTripsByUserID(context.Background(), uint(userID))
	if err != nil {
		http.Error(w, "Failed to retrieve trips", http.StatusInternalServerError)
		return
	}

	httpresponse.SendJSONResponse(w, trips, http.StatusOK)
}

// ReadTripHandler godoc
// @Summary Retrieve a trip by ID
// @Description Get trip details by trip ID
// @Produce json
// @Param id path int true "Trip ID"
// @Success 200 {object} models.Trip "Trip details"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid request"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to retrieve trip"
// @Router /trips/trip/{id} [get]
func (h *TripHandler) ReadTripHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tripIDStr := vars["id"]
	tripID, err := strconv.ParseUint(tripIDStr, 10, 64)
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	trip, err := h.uc.ReadTrip(context.Background(), uint(tripID))
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Failed to retrieve trip",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusInternalServerError)
		return
	}

	httpresponse.SendJSONResponse(w, trip, http.StatusOK)
}
