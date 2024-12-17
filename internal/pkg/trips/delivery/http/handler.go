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

type TripData struct {
	UserID      uint   `json:"user_id"`
	Name        string `json:"name"`
	CityID      uint   `json:"city_id"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Private     bool   `json:"private_trip"`
}

type TripResponse struct {
	Trip  *tripsGen.Trip          `json:"trip"`
	Users []*tripsGen.UserProfile `json:"users"`
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

func ErrorCheck(err error, action string, logger *slog.Logger, ctx context.Context) (httpresponse.Response, int) {
	logContext := log.AppendCtx(ctx, slog.String("action", action))
	logContext = log.AppendCtx(logContext, slog.Any("error", err.Error()))

	if errors.Is(err, models.ErrNotFound) {

		logger.ErrorContext(logContext, fmt.Sprintf("Error during %s operation", action))

		response := httpresponse.Response{
			Message: "Invalid request",
		}
		return response, http.StatusNotFound
	}
	logger.ErrorContext(logContext, fmt.Sprintf("Failed to %s trips", action))
	response := httpresponse.Response{
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
// @Failure 400 {object} httpresponses.Response "Invalid request"
// @Failure 403 {object} httpresponses.Response "Token is missing"
// @Failure 403 {object} httpresponses.Response "Invalid token"
// @Failure 404 {object} httpresponses.Response "Invalid request"
// @Failure 422 {object} httpresponses.Response
// @Failure 500 {object} httpresponses.Response "Failed to create trip"
// @Router /trips [post]
func (h *TripHandler) CreateTripHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self';")
	logCtx := r.Context()
	h.logger.DebugContext(logCtx, "Handling request for creating a trip")

	_, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.WarnContext(logCtx, "Failed to retrieve user ID from context")

		response := httpresponse.Response{
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
		response := httpresponse.Response{
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

	httpresponse.SendJSONResponse(logCtx, w, httpresponse.Response{"Trip created successfully"}, http.StatusCreated, h.logger)
}

// UpdateTripHandler godoc
// @Summary Update an existing trip
// @Description Update trip details by trip ID
// @Accept json
// @Produce json
// @Param id path int true "Trip ID"
// @Param tripData body TripData true "Updated trip details"
// @Success 200 {object} models.Trip "Trip updated successfully"
// @Failure 400 {object} httpresponses.Response "Invalid trip ID"
// @Failure 400 {object} httpresponses.Response "Invalid trip data"
// @Failure 403 {object} httpresponses.Response "Token is missing"
// @Failure 403 {object} httpresponses.Response "Invalid token"
// @Failure 404 {object} httpresponses.Response "Trip not found"
// @Failure 422 {object} httpresponses.Response
// @Failure 500 {object} httpresponses.Response "Failed to update trip"
// @Router /trips/{id} [put]
func (h *TripHandler) UpdateTripHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self';")
	logCtx := r.Context()

	_, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.WarnContext(logCtx, "Failed to retrieve user ID from context")

		response := httpresponse.Response{
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
		response := httpresponse.Response{
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
		response := httpresponse.Response{
			Message: "User cannot edit this trip",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusForbidden, h.logger)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&tripData)
	if err != nil {
		h.logger.WarnContext(logCtx, "Failed to decode trip data", slog.String("trip_data", fmt.Sprintf("%+v", tripData)), slog.String("error", err.Error()))
		response := httpresponse.Response{
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

	httpresponse.SendJSONResponse(logCtx, w, httpresponse.Response{"Trip updated successfully"}, http.StatusOK, h.logger)
}

// DeleteTripHandler godoc
// @Summary Delete a trip
// @Description Delete a trip by trip ID
// @Produce json
// @Param id path int true "Trip ID"
// @Success 204 "Trip deleted successfully"
// @Failure 400 {object} httpresponses.Response "Invalid trip ID"
// @Failure 403 {object} httpresponses.Response "Token is missing"
// @Failure 403 {object} httpresponses.Response "Invalid token"
// @Failure 404 {object} httpresponses.Response "Trip not found"
// @Failure 500 {object} httpresponses.Response "Failed to delete trip"
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

		response := httpresponse.Response{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.WarnContext(logCtx, "Failed to parse trip ID", slog.String("error", err.Error()))
		response := httpresponse.Response{
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

	httpresponse.SendJSONResponse(logCtx, w, httpresponse.Response{"Trip deleted successfully"}, http.StatusNoContent, h.logger)
}

// GetTripsByUserIDHandler godoc
// @Summary Retrieve trips by user ID
// @Description Get all trips for a specific user
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {array} models.Trip "List of trips"
// @Failure 400 {object} httpresponses.Response "Invalid user ID"
// @Failure 403 {object} httpresponses.Response "Token is missing"
// @Failure 403 {object} httpresponses.Response "Invalid token"
// @Failure 404 {object} httpresponses.Response "Invalid user ID"
// @Failure 404 {object} httpresponses.Response "Trips not found"
// @Failure 500 {object} httpresponses.Response "Failed to retrieve trips"
// @Router /users/{userID}/trips [get]
func (h *TripHandler) GetTripsByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()

	userID, ok := r.Context().Value(middleware.IdKey).(uint)

	logCtx = log.AppendCtx(logCtx, slog.Int("user_id", int(userID)))
	h.logger.DebugContext(logCtx, "Handling request for getting trips by user ID")

	if !ok {

		h.logger.WarnContext(logCtx, "Failed to retrieve user ID from context")

		response := httpresponse.Response{
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
			response := httpresponse.Response{
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
	tripResponse := make(models.TripList, len(trip.Trips))
	for i, trip := range trip.Trips {
		tripResponse[i] = models.Trip{
			ID:          uint(trip.Id),
			UserID:      uint(trip.UserId),
			Name:        trip.Name,
			Description: trip.Description,
			CityID:      uint(trip.CityId),
			StartDate:   trip.StartDate,
			EndDate:     trip.EndDate,
			Private:     trip.Private,
			Photos:      trip.Photos,
			CreatedAt:   trip.CreatedAt.AsTime(),
		}
	}
	httpresponse.SendJSONResponse(logCtx, w, tripResponse, http.StatusOK, h.logger)
}

// GetTripHandler godoc
// @Summary Retrieve a trip by ID
// @Description Get trip details by trip ID
// @Produce json
// @Param id path int true "Trip ID"
// @Success 200 {object} models.Trip "Trip details"
// @Failure 400 {object} httpresponses.Response "Invalid trip ID"
// @Failure 404 {object} httpresponses.Response "Trip not found"
// @Failure 500 {object} httpresponses.Response "Failed to retrieve trip"
// @Router /trips/{id} [get]
func (h *TripHandler) GetTripHandler(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()

	vars := mux.Vars(r)
	tripIDStr := vars["id"]

	logCtx = log.AppendCtx(logCtx, slog.String("tripID", tripIDStr))
	h.logger.DebugContext(logCtx, "Handling request for getting trip by ID")

	tripID, err := strconv.ParseUint(tripIDStr, 10, 64)
	if err != nil {
		h.logger.WarnContext(logCtx, "Failed to parse trip ID", slog.String("error", err.Error()))
		response := httpresponse.Response{
			Message: "Invalid trip ID",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	tripResponse, err := h.client.GetTrip(r.Context(), &tripsGen.GetTripRequest{TripId: uint32(tripID)})
	if err != nil {
		response, status := ErrorCheck(err, "retrieve", h.logger, logCtx)
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}

	trip := models.Trip{
		ID:          uint(tripResponse.Trip.Id),
		UserID:      uint(tripResponse.Trip.UserId),
		Name:        tripResponse.Trip.Name,
		Description: tripResponse.Trip.Description,
		CityID:      uint(tripResponse.Trip.CityId),
		StartDate:   tripResponse.Trip.StartDate,
		EndDate:     tripResponse.Trip.EndDate,
		Private:     tripResponse.Trip.Private,
		Photos:      tripResponse.Trip.Photos,
	}

	var users []models.UserProfile
	for _, user := range tripResponse.Users {
		users = append(users, models.UserProfile{
			Login:      user.Login,
			AvatarPath: user.AvatarPath,
			Email:      user.Email,
		})
	}

	h.logger.DebugContext(logCtx, "Successfully got trip by ID")

	// Ответ с данными поездки и пользователей
	response := models.TripResponse{
		Trip:  trip,
		Users: users,
	}
	httpresponse.SendJSONResponse(logCtx, w, response, http.StatusOK, h.logger)
}

// AddPlaceToTripHandler godoc
// @Summary Add a place to a trip
// @Description Add a place with given place_id to a trip
// @Produce json
// @Param id path int true "Trip ID"
// @Param place_id body int true "Place ID"
// @Success 201 "Place added to trip successfully"
// @Failure 400 {object} httpresponses.Response "Invalid trip ID"
// @Failure 400 {object} httpresponses.Response "Invalid place ID"
// @Failure 404 {object} httpresponses.Response "Trip not found"
// @Failure 404 {object} httpresponses.Response "Place not found"
// @Failure 500 {object} httpresponses.Response "Failed to add place to trip"
// @Router /trips/{id} [post]
func (h *TripHandler) AddPlaceToTripHandler(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()

	vars := mux.Vars(r)
	tripIDStr, ok := vars["id"]

	logCtx = log.AppendCtx(logCtx, slog.String("tripID", tripIDStr))
	h.logger.DebugContext(logCtx, "Handling request for adding places to trip")

	if !ok {
		response := httpresponse.Response{
			Message: "Invalid trip ID",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	_, ok = r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.WarnContext(logCtx, "Failed to retrieve user ID from context")

		response := httpresponse.Response{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	tripID, err := strconv.ParseUint(tripIDStr, 10, 64)
	if err != nil {
		response := httpresponse.Response{
			Message: "Invalid trip ID",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}
	var req models.AddPlaceRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.logger.WarnContext(logCtx, "Failed to parse trip ID", slog.String("error", err.Error()))
		response := httpresponse.Response{
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

	httpresponse.SendJSONResponse(logCtx, w, httpresponse.Response{"Place added to trip successfully"}, http.StatusCreated, h.logger)
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
		response := httpresponse.Response{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	tripID, err := strconv.ParseUint(tripIDStr, 10, 64)
	if err != nil {
		h.logger.WarnContext(logCtx, "Failed to parse trip ID", slog.String("error", err.Error()))
		response := httpresponse.Response{
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
		response := httpresponse.Response{
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

	photoResponse := make(models.PhotoList, len(resp.Photos))
	for i, photo := range resp.Photos {
		photoResponse[i] = models.Photo{
			Path: photo.PhotoPath,
		}
	}

	httpresponse.SendJSONResponse(logCtx, w, photoResponse, http.StatusCreated, h.logger)
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
		response := httpresponse.Response{
			Message: "Invalid request body",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}
	tripID, err := strconv.ParseUint(tripIDStr, 10, 32)
	if err != nil {
		h.logger.WarnContext(logCtx, "Failed to parse trip ID", slog.Any("error", err.Error()))
		httpresponse.SendJSONResponse(logCtx, w, httpresponse.Response{Message: "Invalid trip ID"}, http.StatusBadRequest, h.logger)
		return
	}

	req := &tripsGen.DeletePhotoRequest{
		TripId:    uint32(tripID),
		PhotoPath: photoPath.PhotoPath,
	}

	_, err = h.client.DeletePhotoFromTrip(r.Context(), req)
	if err != nil {
		h.logger.ErrorContext(logCtx, "Failed to delete photo from trip", slog.Any("error", err.Error()))
		httpresponse.SendJSONResponse(logCtx, w, httpresponse.Response{Message: "Failed to delete photo"}, http.StatusInternalServerError, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully deleted a photo from trip")

	httpresponse.SendJSONResponse(logCtx, w, httpresponse.Response{"Photo deleted successfully"}, http.StatusOK, h.logger)
}

func (h *TripHandler) CreateSharingLinkHandler(w http.ResponseWriter, r *http.Request) {
	tripIDStr := mux.Vars(r)["id"]
	sharingOption := r.URL.Query().Get("sharing_option")
	urlBase := "therewillbetrip.ru/trips/"
	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for creating a sharing link for a trip", slog.String("tripID", tripIDStr))
	tripID, err := strconv.ParseUint(tripIDStr, 10, 32)
	if err != nil {
		httpresponse.SendJSONResponse(logCtx, w, httpresponse.Response{Message: "Invalid trip ID"}, http.StatusBadRequest, h.logger)
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
		response := models.CreateSharingLinkResponse{
			URL: urlBase + token.Token.Token,
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusOK, h.logger)
		return
	}
	newToken, err := generateToken()
	if err != nil {
		h.logger.Warn("Failed to generate token", slog.String("error", err.Error()))
		response := httpresponse.Response{
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
		httpresponse.SendJSONResponse(logCtx, w, httpresponse.Response{Message: "Failed to create sharing link"}, http.StatusInternalServerError, h.logger)
		return
	}
	response := models.CreateSharingLinkResponse{
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

		response := httpresponse.Response{
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
	tripResponse := models.Trip{
		ID:          uint(trip.Trip.Id),
		UserID:      userID,
		Name:        trip.Trip.Name,
		Description: trip.Trip.Description,
		CityID:      uint(trip.Trip.CityId),
		StartDate:   trip.Trip.StartDate,
		EndDate:     trip.Trip.EndDate,
		Private:     trip.Trip.Private,
		Photos:      trip.Trip.Photos,
		CreatedAt:   trip.Trip.CreatedAt.AsTime(),
	}

	httpresponse.SendJSONResponse(logCtx, w, tripResponse, http.StatusOK, h.logger)
}
