package http

import (
	"2024_2_ThereWillBeName/internal/models"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"2024_2_ThereWillBeName/internal/pkg/trips"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AddPlaceRequest struct {
	PlaceID uint `json:"place_id"`
}

type TripData struct {
	UserID      uint   `json:"user_id"`
	Name        string `json:"name"`
	CityID      uint   `json:"city_id"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Private     bool   `json:"private_trip"`
}

type TripHandler struct {
	uc trips.TripsUsecase
}

func NewTripHandler(uc trips.TripsUsecase) *TripHandler {
	return &TripHandler{uc}
}

func ErrorCheck(err error, action string) (httpresponse.ErrorResponse, int) {
	if errors.Is(err, models.ErrNotFound) {
		log.Printf("%s error: %s", action, err)
		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		return response, http.StatusNotFound
	}
	log.Printf("%s error: %s", action, err)
	response := httpresponse.ErrorResponse{
		Message: fmt.Sprintf("Failed to %s trip", action),
	}
	return response, http.StatusInternalServerError
}

// CreateTripHandler godoc
// @Summary Create a new trip
// @Description Create a new trip with given fields
// @Accept json
// @Produce json
// @Param tripData body TripData true "Trip details"
// @Success 201 {object} models.Trip "Trip created successfully"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid request"
// @Failure 404 {object} httpresponses.ErrorResponse "Invalid request"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to create trip"
// @Router /trips [post]
func (h *TripHandler) CreateTripHandler(w http.ResponseWriter, r *http.Request) {
	var tripData TripData
	err := json.NewDecoder(r.Body).Decode(&tripData)
	if err != nil {
		log.Printf("create error: %s", err)
		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}
	trip := models.Trip{
		UserID:      tripData.UserID,
		Name:        tripData.Name,
		Description: tripData.Description,
		CityID:      tripData.CityID,
		StartDate:   tripData.StartDate,
		EndDate:     tripData.EndDate,
		Private:     tripData.Private,
	}
	err = h.uc.CreateTrip(context.Background(), trip)
	if err != nil {
		response, status := ErrorCheck(err, "create")
		httpresponse.SendJSONResponse(w, response, status)
		return
	}

	httpresponse.SendJSONResponse(w, "Trip created successfully", http.StatusCreated)
}

// UpdateTripHandler godoc
// @Summary Update an existing trip
// @Description Update trip details by trip ID
// @Accept json
// @Produce json
// @Param id path int true "Trip ID"
// @Param tripData body TripData true "Updated trip details"
// @Success 200 {object} models.Trip "Trip updated successfully"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid trip ID"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid trip data"
// @Failure 404 {object} httpresponses.ErrorResponse "Trip not found"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to update trip"
// @Router /trips/{id} [put]
func (h *TripHandler) UpdateTripHandler(w http.ResponseWriter, r *http.Request) {
	var tripData TripData
	vars := mux.Vars(r)
	tripID, err := strconv.Atoi(vars["id"])
	if err != nil || tripID < 0 {
		log.Printf("update error: %s", err)
		response := httpresponse.ErrorResponse{
			Message: "Invalid trip ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&tripData)
	if err != nil {
		log.Printf("update error: %s", err)
		response := httpresponse.ErrorResponse{
			Message: "Invalid trip data",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}
	trip := models.Trip{
		ID:          uint(tripID),
		UserID:      tripData.UserID,
		Name:        tripData.Name,
		Description: tripData.Description,
		CityID:      tripData.CityID,
		StartDate:   tripData.StartDate,
		EndDate:     tripData.EndDate,
		Private:     tripData.Private,
	}
	err = h.uc.UpdateTrip(context.Background(), trip)
	if err != nil {
		response, status := ErrorCheck(err, "update")
		httpresponse.SendJSONResponse(w, response, status)
		return
	}

	httpresponse.SendJSONResponse(w, "Trip updated successfully", http.StatusOK)
}

// DeleteTripHandler godoc
// @Summary Delete a trip
// @Description Delete a trip by trip ID
// @Produce json
// @Param id path int true "Trip ID"
// @Success 204 "Trip deleted successfully"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid trip ID"
// @Failure 404 {object} httpresponses.ErrorResponse "Trip not found"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to delete trip"
// @Router /trips/{id} [delete]
func (h *TripHandler) DeleteTripHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Printf("delete error: %s", err)
		response := httpresponse.ErrorResponse{
			Message: "Invalid trip ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	err = h.uc.DeleteTrip(context.Background(), uint(id))
	if err != nil {
		response, status := ErrorCheck(err, "delete")
		httpresponse.SendJSONResponse(w, response, status)
		return
	}

	httpresponse.SendJSONResponse(w, "Trip deleted successfully", http.StatusNoContent)
}

// GetTripsByUserIDHandler godoc
// @Summary Retrieve trips by user ID
// @Description Get all trips for a specific user
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {array} models.Trip "List of trips"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid user ID"
// @Failure 404 {object} httpresponses.ErrorResponse "Invalid user ID"
// @Failure 404 {object} httpresponses.ErrorResponse "Trips not found"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to retrieve trips"
// @Router /users/{userID}/trips [get]
func (h *TripHandler) GetTripsByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["userID"]
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		log.Printf("retrieve error: %s", err)
		response := httpresponse.ErrorResponse{
			Message: "Invalid user ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}
	pageStr := r.URL.Query().Get("page")
	page := 1
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			log.Printf("retrieve error: %s", err)
			response := httpresponse.ErrorResponse{
				Message: "Invalid page number",
			}
			httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
			return
		}
	}
	limit := 10
	offset := limit * (page - 1)
	trip, err := h.uc.GetTripsByUserID(context.Background(), uint(userID), limit, offset)
	if err != nil {
		response, status := ErrorCheck(err, "retrieve")
		httpresponse.SendJSONResponse(w, response, status)
		return
	}

	httpresponse.SendJSONResponse(w, trip, http.StatusOK)
}

// GetTripHandler godoc
// @Summary Retrieve a trip by ID
// @Description Get trip details by trip ID
// @Produce json
// @Param id path int true "Trip ID"
// @Success 200 {object} models.Trip "Trip details"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid trip ID"
// @Failure 404 {object} httpresponses.ErrorResponse "Trip not found"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to retrieve trip"
// @Router /trips/{id} [get]
func (h *TripHandler) GetTripHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tripIDStr := vars["id"]
	tripID, err := strconv.ParseUint(tripIDStr, 10, 64)
	if err != nil {
		log.Printf("retrieve error: %s", err)
		response := httpresponse.ErrorResponse{
			Message: "Invalid trip ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	trip, err := h.uc.GetTrip(context.Background(), uint(tripID))
	if err != nil {
		response, status := ErrorCheck(err, "retrieve")
		httpresponse.SendJSONResponse(w, response, status)
		return
	}

	httpresponse.SendJSONResponse(w, trip, http.StatusOK)
}

// AddPlaceToTripHandler godoc
// @Summary Add a place to a trip
// @Description Add a place with given place_id to a trip
// @Produce json
// @Param id path int true "Trip ID"
// @Param place_id body int true "Place ID"
// @Success 201 "Place added to trip successfully"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid trip ID"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid place ID"
// @Failure 404 {object} httpresponses.ErrorResponse "Trip not found"
// @Failure 404 {object} httpresponses.ErrorResponse "Place not found"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to add place to trip"
// @Router /trips/{id} [post]
func (h *TripHandler) AddPlaceToTripHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tripIDStr, ok := vars["id"]
	if !ok {
		response := httpresponse.ErrorResponse{
			Message: "Invalid trip ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	tripID, err := strconv.ParseUint(tripIDStr, 10, 64)
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Invalid trip ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}
	var req AddPlaceRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("add place error: %s", err)
		response := httpresponse.ErrorResponse{
			Message: "Invalid place ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	err = h.uc.AddPlaceToTrip(context.Background(), uint(tripID), req.PlaceID)
	if err != nil {
		response, status := ErrorCheck(err, "add place")
		httpresponse.SendJSONResponse(w, response, status)
		return
	}

	httpresponse.SendJSONResponse(w, "Place added to trip successfully", http.StatusCreated)
}
