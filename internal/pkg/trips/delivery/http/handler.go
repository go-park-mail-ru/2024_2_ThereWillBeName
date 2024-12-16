package http

import (
	"2024_2_ThereWillBeName/internal/models"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	log "2024_2_ThereWillBeName/internal/pkg/logger"
	"2024_2_ThereWillBeName/internal/pkg/middleware"
	tripsGen "2024_2_ThereWillBeName/internal/pkg/trips/delivery/grpc/gen"
	"2024_2_ThereWillBeName/internal/validator"

	"context"
	"crypto/rand"
	"encoding/base64"
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

type CreateSharingLinkResponse struct {
	URL string `json:"url"`
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

func generateToken() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func ErrorCheck(err error, action string, logger *slog.Logger, ctx context.Context) (httpresponse.ErrorResponse, int) {
	logContext := log.AppendCtx(ctx, slog.String("action", action))
	logContext = log.AppendCtx(logContext, slog.Any("error", err.Error()))

	if errors.Is(err, models.ErrNotFound) {

		logger.ErrorContext(logContext, fmt.Sprintf("Error during %s operation", action))

		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		return response, http.StatusNotFound
	}
	logger.ErrorContext(logContext, fmt.Sprintf("Failed to %s trips", action))
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
	logCtx := r.Context()
	h.logger.DebugContext(logCtx, "Handling request for creating a trip")

	_, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.WarnContext(logCtx, "Failed to retrieve user ID from context")

		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	var tripData TripData
	err := json.NewDecoder(r.Body).Decode(&tripData)

	if err != nil {
		h.logger.WarnContext(logCtx, "Failed to decode trip data",
			slog.String("error", err.Error()),
			slog.String("trip_data", fmt.Sprintf("%+v", tripData)))
		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
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
		h.logger.WarnContext(logCtx, "Review data is not valid")
		httpresponse.SendJSONResponse(logCtx, w, nil, http.StatusUnprocessableEntity, h.logger)
		return
	}

	trip.Name = template.HTMLEscapeString(trip.Name)
	trip.Description = template.HTMLEscapeString(trip.Description)

	tripRequest := &tripsGen.Trip{
		Id:          uint32(trip.ID),
		UserId:      uint32(trip.UserID),
		Name:        trip.Name,
		Description: trip.Description,
		CityId:      uint32(trip.CityID),
		StartDate:   trip.StartDate,
		EndDate:     trip.EndDate,
		Private:     trip.Private,
	}

	h.logger.DebugContext(logCtx, "Trip request details", slog.Any("tripRequest", tripRequest))

	_, err = h.client.CreateTrip(r.Context(), &tripsGen.CreateTripRequest{Trip: tripRequest})
	if err != nil {
		response, status := ErrorCheck(err, "create", h.logger, logCtx)
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully created a trip")

	httpresponse.SendJSONResponse(logCtx, w, "Trip created successfully", http.StatusCreated, h.logger)
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
	logCtx := r.Context()

	_, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.WarnContext(logCtx, "Failed to retrieve user ID from context")

		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	var tripData TripData

	vars := mux.Vars(r)
	tripID, err := strconv.Atoi(vars["id"])
	logCtx = log.AppendCtx(logCtx, slog.Int("trip_id", tripID))
	h.logger.DebugContext(logCtx, "Handling request for updating a trip")

	if err != nil || tripID < 0 {
		h.logger.WarnContext(logCtx, "Failed to parse trip ID", slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid trip ID",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	optionRequest := &tripsGen.GetSharingOptionRequest{
		TripId: uint32(tripID),
		UserId: uint32(r.Context().Value(middleware.IdKey).(uint)),
	}
	sharingOption, err := h.client.GetSharingOption(r.Context(), optionRequest)
	if err != nil {
		response, status := ErrorCheck(err, "retrieve sharing option", h.logger, logCtx)
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}

	if sharingOption.SharingOption != "editing" {
		response := httpresponse.ErrorResponse{
			Message: "User cannot edit this trip",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusForbidden, h.logger)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&tripData)
	if err != nil {
		h.logger.WarnContext(logCtx, "Failed to decode trip data", slog.String("trip_data", fmt.Sprintf("%+v", tripData)), slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid trip data",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
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
		h.logger.WarnContext(logCtx, "Trip data is not valid")
		httpresponse.SendJSONResponse(logCtx, w, nil, http.StatusUnprocessableEntity, h.logger)
		return
	}

	trip.Name = template.HTMLEscapeString(trip.Name)
	trip.Description = template.HTMLEscapeString(trip.Description)

	tripRequest := &tripsGen.Trip{
		Id:          uint32(tripID),
		UserId:      uint32(trip.UserID),
		Name:        trip.Name,
		Description: trip.Description,
		CityId:      uint32(trip.CityID),
		StartDate:   trip.StartDate,
		EndDate:     trip.EndDate,
		Private:     trip.Private,
	}

	h.logger.DebugContext(logCtx, "Trip request details", slog.Any("tripRequest", tripRequest))

	_, err = h.client.UpdateTrip(r.Context(), &tripsGen.UpdateTripRequest{Trip: tripRequest})
	if err != nil {
		response, status := ErrorCheck(err, "update", h.logger, logCtx)
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}
	h.logger.DebugContext(logCtx, "Successfully updated a trip")

	httpresponse.SendJSONResponse(logCtx, w, "Trip updated successfully", http.StatusOK, h.logger)
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
	logCtx := r.Context()

	vars := mux.Vars(r)
	idStr := vars["id"]

	logCtx = log.AppendCtx(logCtx, slog.String("tripID", idStr))
	h.logger.DebugContext(logCtx, "Handling request for deleting a trip")

	_, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.WarnContext(logCtx, "Failed to retrieve user ID from context")

		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.WarnContext(logCtx, "Failed to parse trip ID", slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid trip ID",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	_, err = h.client.DeleteTrip(r.Context(), &tripsGen.DeleteTripRequest{Id: uint32(id)})
	if err != nil {
		response, status := ErrorCheck(err, "delete", h.logger, logCtx)
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}
	h.logger.DebugContext(logCtx, "Successfully deleted a trip")

	httpresponse.SendJSONResponse(logCtx, w, "Trip deleted successfully", http.StatusNoContent, h.logger)
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
	logCtx := r.Context()

	userID, ok := r.Context().Value(middleware.IdKey).(uint)

	logCtx = log.AppendCtx(logCtx, slog.Int("user_id", int(userID)))
	h.logger.DebugContext(logCtx, "Handling request for getting trips by user ID")

	if !ok {

		h.logger.WarnContext(logCtx, "Failed to retrieve user ID from context")

		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
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
			httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
			return
		}
	}
	limit := 10
	offset := limit * (page - 1)

	trip, err := h.client.GetTripsByUserID(r.Context(), &tripsGen.GetTripsByUserIDRequest{
		UserId: uint32(userID),
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		response, status := ErrorCheck(err, "retrieve", h.logger, logCtx)
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully got trips by user ID", slog.Int("trips_count", len(trip.Trips)))
	tripArr := trip.Trips

	httpresponse.SendJSONResponse(logCtx, w, tripArr, http.StatusOK, h.logger)
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
	logCtx := r.Context()

	vars := mux.Vars(r)
	tripIDStr := vars["id"]

	logCtx = log.AppendCtx(logCtx, slog.String("tripID", tripIDStr))
	h.logger.DebugContext(logCtx, "Handling request for getting trip by ID")

	_, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.WarnContext(logCtx, "Failed to retrieve user ID from context")

		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	tripID, err := strconv.ParseUint(tripIDStr, 10, 64)
	if err != nil {
		h.logger.WarnContext(logCtx, "Failed to parse trip ID", slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid trip ID",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	trip, err := h.client.GetTrip(r.Context(), &tripsGen.GetTripRequest{TripId: uint32(tripID)})
	if err != nil {
		response, status := ErrorCheck(err, "retrieve", h.logger, logCtx)
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully got trip by ID")
	tripResponse := trip.Trip

	httpresponse.SendJSONResponse(logCtx, w, tripResponse, http.StatusOK, h.logger)
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
	logCtx := r.Context()

	vars := mux.Vars(r)
	tripIDStr, ok := vars["id"]

	logCtx = log.AppendCtx(logCtx, slog.String("tripID", tripIDStr))
	h.logger.DebugContext(logCtx, "Handling request for adding places to trip")

	if !ok {
		response := httpresponse.ErrorResponse{
			Message: "Invalid trip ID",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	_, ok = r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.WarnContext(logCtx, "Failed to retrieve user ID from context")

		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	tripID, err := strconv.ParseUint(tripIDStr, 10, 64)
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Invalid trip ID",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}
	var req AddPlaceRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.logger.WarnContext(logCtx, "Failed to parse trip ID", slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid place ID",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Adding place with ID", slog.Int("place_id", int(req.PlaceID)))

	_, err = h.client.AddPlaceToTrip(r.Context(), &tripsGen.AddPlaceToTripRequest{
		TripId:  uint32(tripID),
		PlaceId: uint32(req.PlaceID),
	})
	if err != nil {
		response, status := ErrorCheck(err, "add place", h.logger, context.Background())
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully added place to trip")

	httpresponse.SendJSONResponse(logCtx, w, "Place added to trip successfully", http.StatusCreated, h.logger)
}

func (h *TripHandler) AddPhotosToTripHandler(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()

	vars := mux.Vars(r)
	tripIDStr := vars["id"]

	logCtx = log.AppendCtx(logCtx, slog.String("tripID", tripIDStr))
	h.logger.DebugContext(logCtx, "Handling request for adding photos to a trip")

	_, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {
		h.logger.WarnContext(logCtx, "Failed to retrieve user ID from context")
		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	tripID, err := strconv.ParseUint(tripIDStr, 10, 64)
	if err != nil {
		h.logger.WarnContext(logCtx, "Failed to parse trip ID", slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid trip ID",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	var photosRequest struct {
		Photos []string `json:"photos"`
	}

	err = json.NewDecoder(r.Body).Decode(&photosRequest)
	if err != nil {
		h.logger.WarnContext(logCtx, "Failed to decode photos request body", slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid request body",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}
	resp, err := h.client.AddPhotosToTrip(r.Context(), &tripsGen.AddPhotosToTripRequest{
		TripId: uint32(tripID),
		Photos: photosRequest.Photos,
	})
	if err != nil {
		response, status := ErrorCheck(err, "add photos", h.logger, logCtx)
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully added photos to the trip")

	httpresponse.SendJSONResponse(logCtx, w, resp.Photos, http.StatusCreated, h.logger)
}

func (h *TripHandler) DeletePhotoHandler(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()

	tripIDStr := mux.Vars(r)["id"]
	var photoPath struct {
		PhotoPath string `json:"photo_path"`
	}

	logCtx = log.AppendCtx(logCtx, slog.String("tripID", tripIDStr))
	h.logger.DebugContext(logCtx, "Handling request for deleting photo from a trip")
	err := json.NewDecoder(r.Body).Decode(&photoPath)
	if err != nil {
		h.logger.WarnContext(logCtx, "Failed to decode photos delete request body", slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid request body",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}
	tripID, err := strconv.ParseUint(tripIDStr, 10, 32)
	if err != nil {
		h.logger.WarnContext(logCtx, "Failed to parse trip ID", slog.Any("error", err.Error()))
		httpresponse.SendJSONResponse(logCtx, w, httpresponse.ErrorResponse{Message: "Invalid trip ID"}, http.StatusBadRequest, h.logger)
		return
	}

	req := &tripsGen.DeletePhotoRequest{
		TripId:    uint32(tripID),
		PhotoPath: photoPath.PhotoPath,
	}

	_, err = h.client.DeletePhotoFromTrip(r.Context(), req)
	if err != nil {
		h.logger.ErrorContext(logCtx, "Failed to delete photo from trip", slog.Any("error", err.Error()))
		httpresponse.SendJSONResponse(logCtx, w, httpresponse.ErrorResponse{Message: "Failed to delete photo"}, http.StatusInternalServerError, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully deleted a photo from trip")

	httpresponse.SendJSONResponse(logCtx, w, map[string]string{"message": "Photo deleted successfully"}, http.StatusOK, h.logger)
}

func (h *TripHandler) CreateSharingLinkHandler(w http.ResponseWriter, r *http.Request) {
	tripIDStr := mux.Vars(r)["id"]
	sharingOption := r.URL.Query().Get("sharing_option")
	urlBase := "therewillbetrip.ru/trips/"
	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for creating a sharing link for a trip", slog.String("tripID", tripIDStr))
	tripID, err := strconv.ParseUint(tripIDStr, 10, 32)
	if err != nil {
		httpresponse.SendJSONResponse(logCtx, w, httpresponse.ErrorResponse{Message: "Invalid trip ID"}, http.StatusBadRequest, h.logger)
		return
	}
	req := &tripsGen.GetSharingTokenRequest{
		TripId: uint32(tripID),
	}
	token, err := h.client.GetSharingToken(r.Context(), req)
	if err != nil {
		response, status := ErrorCheck(err, "get sharing token", h.logger, context.Background())
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}
	if token.Token.Token != "" {
		response := CreateSharingLinkResponse{
			URL: urlBase + token.Token.Token,
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusOK, h.logger)
		return
	}
	newToken, err := generateToken()
	if err != nil {
		h.logger.Warn("Failed to generate token", slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid request body",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	newReq := &tripsGen.CreateSharingLinkRequest{
		TripId:        uint32(tripID),
		Token:         newToken,
		SharingOption: sharingOption,
	}

	_, err = h.client.CreateSharingLink(r.Context(), newReq)
	if err != nil {
		httpresponse.SendJSONResponse(logCtx, w, httpresponse.ErrorResponse{Message: "Failed to create sharing link"}, http.StatusInternalServerError, h.logger)
		return
	}
	response := CreateSharingLinkResponse{
		URL: urlBase + newToken,
	}

	httpresponse.SendJSONResponse(logCtx, w, response, http.StatusOK, h.logger)
}

func (h *TripHandler) GetTripBySharingToken(w http.ResponseWriter, r *http.Request) {
	token := mux.Vars(r)["sharing_token"]
	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)

	req := &tripsGen.GetTripBySharingTokenRequest{
		Token: token,
	}
	trip, err := h.client.GetTripBySharingToken(r.Context(), req)
	if err != nil {
		response, status := ErrorCheck(err, "retrieve trip by sharing token", h.logger, logCtx)
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}
	userID, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.WarnContext(logCtx, "Failed to retrieve user ID from context")

		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}
	addUserReq := &tripsGen.AddUserToTripRequest{
		TripId: trip.Trip.Id,
		UserId: uint32(userID),
	}
	_, err = h.client.AddUserToTrip(r.Context(), addUserReq)
	if err != nil {
		response, status := ErrorCheck(err, "add user to trip", h.logger, logCtx)
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}
	h.logger.DebugContext(logCtx, "Successfully got trip by sharing token")
	tripResponse := trip.Trip

	httpresponse.SendJSONResponse(logCtx, w, tripResponse, http.StatusOK, h.logger)
}
