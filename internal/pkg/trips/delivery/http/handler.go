package http

import (
	"2024_2_ThereWillBeName/internal/models"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	log "2024_2_ThereWillBeName/internal/pkg/logger"
	"2024_2_ThereWillBeName/internal/pkg/middleware"
	tripsGen "2024_2_ThereWillBeName/internal/pkg/trips/delivery/grpc/gen"
	"2024_2_ThereWillBeName/internal/validator"

	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log/slog"
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
	client tripsGen.TripsClient
	logger *slog.Logger
}

func NewTripHandler(client tripsGen.TripsClient, logger *slog.Logger) *TripHandler {
	return &TripHandler{client, logger}
}

func ErrorCheck(err error, action string, logger *slog.Logger, ctx context.Context) (httpresponse.ErrorResponse, int) {
	if errors.Is(err, models.ErrNotFound) {

		logContext := log.AppendCtx(ctx, slog.String("action", action))
		logger.ErrorContext(logContext, fmt.Sprintf("Error during %s operation", action), slog.Any("error", err.Error()))

		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		return response, http.StatusNotFound
	}
	logContext := log.AppendCtx(ctx, slog.String("action", action))
	logger.ErrorContext(logContext, fmt.Sprintf("Failed to %s trips", action), slog.Any("error", err.Error()))
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
// @Failure 403 {object} httpresponses.ErrorResponse "Token is missing"
// @Failure 403 {object} httpresponses.ErrorResponse "Invalid token"
// @Failure 404 {object} httpresponses.ErrorResponse "Invalid request"
// @Failure 422 {object} httpresponses.ErrorResponse
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to create trip"
// @Router /trips [post]
func (h *TripHandler) CreateTripHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self';")
	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for creating a trip")

	_, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.Warn("Failed to retrieve user ID from context")

		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized, h.logger)
		return
	}

	var tripData TripData
	err := json.NewDecoder(r.Body).Decode(&tripData)

	if err != nil {
		h.logger.Warn("Failed to decode trip data",
			slog.String("error", err.Error()),
			slog.String("trip_data", fmt.Sprintf("%+v", tripData)))
		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
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

	v := validator.New()
	if models.ValidateTrip(v, &trip); !v.Valid() {
		httpresponse.SendJSONResponse(w, nil, http.StatusUnprocessableEntity, h.logger)
		return
	}

	trip.Name = template.HTMLEscapeString(trip.Name)
	trip.Description = template.HTMLEscapeString(trip.Description)

	// err = h.uc.CreateTrip(context.Background(), trip)
	_, err = h.client.CreateTrip(r.Context(), &tripsGen.CreateTripRequest{Trip: &tripsGen.Trip{
		Id:          uint32(trip.ID),
		UserId:      uint32(trip.UserID),
		Name:        trip.Name,
		Description: trip.Description,
		CityId:      uint32(trip.CityID),
		StartDate:   trip.StartDate,
		EndDate:     trip.EndDate,
		Private:     trip.Private,
	}})
	if err != nil {
		response, status := ErrorCheck(err, "create", h.logger, context.Background())
		httpresponse.SendJSONResponse(w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully created a trip")

	httpresponse.SendJSONResponse(w, "Trip created successfully", http.StatusCreated, h.logger)
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
// @Failure 403 {object} httpresponses.ErrorResponse "Token is missing"
// @Failure 403 {object} httpresponses.ErrorResponse "Invalid token"
// @Failure 404 {object} httpresponses.ErrorResponse "Trip not found"
// @Failure 422 {object} httpresponses.ErrorResponse
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to update trip"
// @Router /trips/{id} [put]
func (h *TripHandler) UpdateTripHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self';")
	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for updating a trip")

	_, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.Warn("Failed to retrieve user ID from context")

		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized, h.logger)
		return
	}

	var tripData TripData

	vars := mux.Vars(r)
	tripID, err := strconv.Atoi(vars["id"])
	if err != nil || tripID < 0 {
		h.logger.Warn("Failed to parse trip ID", slog.Int("tripID", tripID), slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid trip ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&tripData)
	if err != nil {
		h.logger.Warn("Failed to decode trip data", slog.String("trip_data", fmt.Sprintf("%+v", tripData)), slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid trip data",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
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

	v := validator.New()
	if models.ValidateTrip(v, &trip); !v.Valid() {
		httpresponse.SendJSONResponse(w, nil, http.StatusUnprocessableEntity, h.logger)
		return
	}

	trip.Name = template.HTMLEscapeString(trip.Name)
	trip.Description = template.HTMLEscapeString(trip.Description)

	// err = h.uc.UpdateTrip(context.Background(), trip)
	_, err = h.client.UpdateTrip(r.Context(), &tripsGen.UpdateTripRequest{Trip: &tripsGen.Trip{
		Id:          uint32(trip.ID),
		UserId:      uint32(trip.UserID),
		Name:        trip.Name,
		Description: trip.Description,
		CityId:      uint32(trip.CityID),
		StartDate:   trip.StartDate,
		EndDate:     trip.EndDate,
		Private:     trip.Private,
	}})
	if err != nil {
		logCtx := log.AppendCtx(context.Background(), slog.Int("tripID", tripID))
		response, status := ErrorCheck(err, "update", h.logger, logCtx)
		httpresponse.SendJSONResponse(w, response, status, h.logger)
		return
	}
	h.logger.DebugContext(logCtx, "Successfully updated a trip")

	httpresponse.SendJSONResponse(w, "Trip updated successfully", http.StatusOK, h.logger)
}

// DeleteTripHandler godoc
// @Summary Delete a trip
// @Description Delete a trip by trip ID
// @Produce json
// @Param id path int true "Trip ID"
// @Success 204 "Trip deleted successfully"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid trip ID"
// @Failure 403 {object} httpresponses.ErrorResponse "Token is missing"
// @Failure 403 {object} httpresponses.ErrorResponse "Invalid token"
// @Failure 404 {object} httpresponses.ErrorResponse "Trip not found"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to delete trip"
// @Router /trips/{id} [delete]
func (h *TripHandler) DeleteTripHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for deleting a trip", slog.String("tripID", idStr))

	_, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.Warn("Failed to retrieve user ID from context")

		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized, h.logger)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.Warn("Failed to parse trip ID", slog.String("tripID", idStr), slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid trip ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}

	// err = h.uc.DeleteTrip(context.Background(), uint(id))
	_, err = h.client.DeleteTrip(r.Context(), &tripsGen.DeleteTripRequest{Id: uint32(id)})
	if err != nil {
		logCtx := log.AppendCtx(context.Background(), slog.String("tripID", idStr))
		response, status := ErrorCheck(err, "delete", h.logger, logCtx)
		httpresponse.SendJSONResponse(w, response, status, h.logger)
		return
	}
	h.logger.DebugContext(logCtx, "Successfully deleted a trip")

	httpresponse.SendJSONResponse(w, "Trip deleted successfully", http.StatusNoContent, h.logger)
}

// GetTripsByUserIDHandler godoc
// @Summary Retrieve trips by user ID
// @Description Get all trips for a specific user
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {array} models.Trip "List of trips"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid user ID"
// @Failure 403 {object} httpresponses.ErrorResponse "Token is missing"
// @Failure 403 {object} httpresponses.ErrorResponse "Invalid token"
// @Failure 404 {object} httpresponses.ErrorResponse "Invalid user ID"
// @Failure 404 {object} httpresponses.ErrorResponse "Trips not found"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to retrieve trips"
// @Router /users/{userID}/trips [get]
func (h *TripHandler) GetTripsByUserIDHandler(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.IdKey).(uint)

	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for getting trips by user ID", slog.Int("placeID", int(userID)))

	if !ok {

		h.logger.Warn("Failed to retrieve user ID from context")

		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized, h.logger)
		return
	}

	var err error
	pageStr := r.URL.Query().Get("page")
	page := 1
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			response := httpresponse.ErrorResponse{
				Message: "Invalid page number",
			}
			httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
			return
		}
	}
	limit := 10
	offset := limit * (page - 1)
	// trip, err := h.uc.GetTripsByUserID(context.Background(), uint(userID), limit, offset)
	trip, err := h.client.GetTripsByUserID(r.Context(), &tripsGen.GetTripsByUserIDRequest{
		UserId: uint32(userID),
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		logCtx := log.AppendCtx(context.Background(), slog.Int("userID", int(userID)))
		response, status := ErrorCheck(err, "retrieve", h.logger, logCtx)
		httpresponse.SendJSONResponse(w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully got trips by user ID")
	tripArr := trip.Trips

	httpresponse.SendJSONResponse(w, tripArr, http.StatusOK, h.logger)
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

	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for getting trip by ID", slog.String("tripID", tripIDStr))

	_, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.Warn("Failed to retrieve user ID from context")

		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized, h.logger)
		return
	}

	tripID, err := strconv.ParseUint(tripIDStr, 10, 64)
	if err != nil {
		h.logger.Warn("Failed to parse trip ID", slog.String("tripID", tripIDStr), slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid trip ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}

	// trip, err := h.uc.GetTrip(context.Background(), uint(tripID))
	trip, err := h.client.GetTrip(r.Context(), &tripsGen.GetTripRequest{TripId: uint32(tripID)})
	if err != nil {
		logCtx := log.AppendCtx(context.Background(), slog.String("tripID", tripIDStr))
		response, status := ErrorCheck(err, "retrieve", h.logger, logCtx)
		httpresponse.SendJSONResponse(w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully got trip by ID")
	tripResponse := trip.Trip

	httpresponse.SendJSONResponse(w, tripResponse, http.StatusOK, h.logger)
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
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}

	_, ok = r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.Warn("Failed to retrieve user ID from context")

		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized, h.logger)
		return
	}

	tripID, err := strconv.ParseUint(tripIDStr, 10, 64)
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Invalid trip ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}
	var req AddPlaceRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.logger.Warn("Failed to parse trip ID", slog.String("tripID", tripIDStr), slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid place ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}

	// err = h.uc.AddPlaceToTrip(context.Background(), uint(tripID), req.PlaceID)
	_, err = h.client.AddPlaceToTrip(r.Context(), &tripsGen.AddPlaceToTripRequest{
		TripId:  uint32(tripID),
		PlaceId: uint32(req.PlaceID),
	})
	if err != nil {
		response, status := ErrorCheck(err, "add place", h.logger, context.Background())
		httpresponse.SendJSONResponse(w, response, status, h.logger)
		return
	}

	httpresponse.SendJSONResponse(w, "Place added to trip successfully", http.StatusCreated, h.logger)
}

func (h *TripHandler) AddPhotosToTripHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tripIDStr := vars["id"]

	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for adding photos to a trip", slog.String("tripID", tripIDStr))

	_, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {
		h.logger.Warn("Failed to retrieve user ID from context")
		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized, h.logger)
		return
	}

	tripID, err := strconv.ParseUint(tripIDStr, 10, 64)
	if err != nil {
		h.logger.Warn("Failed to parse trip ID", slog.String("tripID", tripIDStr), slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid trip ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}

	var photosRequest struct {
		Photos []string `json:"photos"`
	}

	err = json.NewDecoder(r.Body).Decode(&photosRequest)
	if err != nil {
		h.logger.Warn("Failed to decode photos request body", slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid request body",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}

	resp, err := h.client.AddPhotosToTrip(r.Context(), &tripsGen.AddPhotosToTripRequest{
		TripId: uint32(tripID),
		Photos: photosRequest.Photos,
	})
	if err != nil {
		logCtx := log.AppendCtx(context.Background(), slog.String("tripID", tripIDStr))
		response, status := ErrorCheck(err, "add photos", h.logger, logCtx)
		httpresponse.SendJSONResponse(w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully added photos to the trip")

	httpresponse.SendJSONResponse(w, resp.Photos, http.StatusCreated, h.logger)
}
