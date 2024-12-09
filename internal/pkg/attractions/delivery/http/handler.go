package http

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/attractions/delivery/grpc/gen"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	log "2024_2_ThereWillBeName/internal/pkg/logger"
	"errors"
	"html/template"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PlacesHandler struct {
	client gen.AttractionsClient
	logger *slog.Logger
}

func NewPlacesHandler(client gen.AttractionsClient, logger *slog.Logger) *PlacesHandler {
	return &PlacesHandler{client, logger}
}

// GetPlaceHandler godoc
// @Summary Get a list of attractions
// @Description Retrieve a list of attractions from the database
// @Produce json
// @Success 200 {array} models.GetPlace "List of attractions"
// @Failure 400 {object} httpresponses.ErrorResponse "Bad request"
// @Failure 500 {object} httpresponses.ErrorResponse "Internal Server Error"
// @Router /attractions [get]
func (h *PlacesHandler) GetPlacesHandler(w http.ResponseWriter, r *http.Request) {
	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for searching attractions")

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest, h.logger)
		h.logger.Warn("Invalid offset parameter", slog.String("error", err.Error()))
		return
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest, h.logger)
		h.logger.Warn("Invalid limit parameter", slog.String("error", err.Error()))
		return
	}
	places, err := h.client.GetPlaces(r.Context(), &gen.GetPlacesRequest{Limit: int32(limit), Offset: int32(offset)})
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError, h.logger)
		h.logger.Error("Couldn't get list of attractions",
			slog.Int("limit", limit),
			slog.Int("offset", offset),
			slog.String("error", err.Error()))
		return
	}

	h.logger.DebugContext(logCtx, "Successfully retrieved attractions")

	httpresponse.SendJSONResponse(w, places.Places, http.StatusOK, h.logger)
}

// PostPlaceHandler godoc
// @Summary Create a new place
// @Description Add a new place to the database
// @Accept json
// @Produce json
// @Param place body models.CreatePlace true "Place data"
// @Success 201 {object} httpresponses.ErrorResponse "Place successfully created"
// @Failure 400 {object} httpresponses.ErrorResponse
// @Failure 403 {object} httpresponses.ErrorResponse "Token is missing"
// @Failure 403 {object} httpresponses.ErrorResponse "Invalid token"
// @Failure 422 {object} httpresponses.ErrorResponse
// @Failure 500 {object} httpresponses.ErrorResponse
// @Router /attractions [post]
//func (h *PlacesHandler) PostPlaceHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self';")
//	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
//	h.logger.DebugContext(logCtx, "Handling request for creating a place")
//
//	var place models.CreatePlace
//	if err := json.NewDecoder(r.Body).Decode(&place); err != nil {
//		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest, h.logger)
//		h.logger.Warn("Failed to decode place data",
//			slog.String("error", err.Error()),
//			slog.String("place_data", fmt.Sprintf("%+v", place)))
//		return
//	}
//	v := validator.New()
//	if models.ValidateCreatePlace(v, &place); !v.Valid() {
//		httpresponse.SendJSONResponse(w, nil, http.StatusUnprocessableEntity, h.logger)
//		return
//	}
//
//	place.Name = template.HTMLEscapeString(place.Name)
//	place.ImagePath = template.HTMLEscapeString(place.ImagePath)
//	place.Description = template.HTMLEscapeString(place.Description)
//	place.Address = template.HTMLEscapeString(place.Address)
//	place.PhoneNumber = template.HTMLEscapeString(place.PhoneNumber)
//
//	if err := h.uc.CreatePlace(r.Context(), place); err != nil {
//		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError, h.logger)
//		h.logger.Error("Failed to create place",
//			slog.String("placeName", place.Name),
//			slog.String("error", err.Error()))
//		return
//	}
//
//	h.logger.DebugContext(logCtx, "Successfully created place")
//
//	httpresponse.SendJSONResponse(w, "Place succesfully created", http.StatusCreated, h.logger)
//}
//
//// PutPlaceHandler godoc
//// @Summary Update an existing place
//// @Description Update the details of an existing place in the database
//// @Accept json
//// @Produce json
//// @Param place body models.UpdatePlace true "Updated place data"
//// @Success 200 {object} httpresponses.ErrorResponse "Place successfully updated"
//// @Failure 400 {object} httpresponses.ErrorResponse
//// @Failure 403 {object} httpresponses.ErrorResponse "Token is missing"
//// @Failure 403 {object} httpresponses.ErrorResponse "Invalid token"
//// @Failure 422 {object} httpresponses.ErrorResponse
//// @Failure 500 {object} httpresponses.ErrorResponse
//// @Router /attractions/{id} [put]
//func (h *PlacesHandler) PutPlaceHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self';")
//	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
//	h.logger.DebugContext(logCtx, "Handling request for updating a place")
//
//	var place models.UpdatePlace
//
//	if err := json.NewDecoder(r.Body).Decode(&place); err != nil {
//		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest, h.logger)
//		h.logger.Warn("Failed to decode place data",
//			slog.String("error", err.Error()),
//			slog.String("place_data", fmt.Sprintf("%+v", place)))
//		return
//	}
//
//	v := validator.New()
//	if models.ValidateUpdatePlace(v, &place); !v.Valid() {
//		httpresponse.SendJSONResponse(w, nil, http.StatusUnprocessableEntity, h.logger)
//		return
//	}
//
//	place.Name = template.HTMLEscapeString(place.Name)
//	place.ImagePath = template.HTMLEscapeString(place.ImagePath)
//	place.Description = template.HTMLEscapeString(place.Description)
//	place.Address = template.HTMLEscapeString(place.Address)
//	place.PhoneNumber = template.HTMLEscapeString(place.PhoneNumber)
//
//	if err := h.uc.UpdatePlace(r.Context(), place); err != nil {
//		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError, h.logger)
//		h.logger.Error("Failed to update place",
//			slog.Int("placeID", place.ID),
//			slog.String("error", err.Error()))
//		return
//	}
//
//	h.logger.DebugContext(logCtx, "Successfully updated place")
//
//	httpresponse.SendJSONResponse(w, "place successfully updated", http.StatusOK, h.logger)
//}
//
//// DeletePlaceHandler godoc
//// @Summary Delete an existing place
//// @Description Remove a place from the database by its name
//// @Produce json
//// @Param name body string true "Name of the place to be deleted"
//// @Success 200 {object} httpresponses.ErrorResponse "Place successfully deleted"
//// @Failure 400 {object} httpresponses.ErrorResponse
//// @Failure 403 {object} httpresponses.ErrorResponse "Token is missing"
//// @Failure 403 {object} httpresponses.ErrorResponse "Invalid token"
//// @Failure 500 {object} httpresponses.ErrorResponse
//// @Router /attractions/{id} [delete]
//func (h *PlacesHandler) DeletePlaceHandler(w http.ResponseWriter, r *http.Request) {
//
//	idStr := mux.Vars(r)["id"]
//
//	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
//	h.logger.DebugContext(logCtx, "Handling request for deleting a place by ID", slog.String("placeID", idStr))
//
//	id, err := strconv.ParseUint(idStr, 10, 64)
//	if err != nil {
//		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest, h.logger)
//		h.logger.Warn("Failed to parse place ID", slog.String("placeID", idStr), slog.String("error", err.Error()))
//		return
//	}
//	if err := h.uc.DeletePlace(r.Context(), uint(id)); err != nil {
//		if errors.Is(err, models.ErrNotFound) {
//			httpresponse.SendJSONResponse(w, nil, http.StatusNotFound, h.logger)
//			h.logger.Warn("Place not found", slog.String("placeID", idStr), slog.String("error", err.Error()))
//			return
//		}
//		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError, h.logger)
//		h.logger.Error("Failed to delete a place", slog.String("placeID", idStr), slog.String("error", err.Error()))
//		return
//	}
//	httpresponse.SendJSONResponse(w, "place successfully deleted", http.StatusOK, h.logger)
//
//	h.logger.DebugContext(logCtx, "Successfully updated place")
//}

// GetPlaceHandler godoc
// @Summary Retrieve an existing place
// @Description Get details of a place from the database by its id
// @Produce json
// @Param id body int true "ID of the place to retrieve"
// @Success 200 {object} models.GetPlace "Details of the requested place"
// @Failure 400 {object} httpresponses.ErrorResponse
// @Failure 500 {object} httpresponses.ErrorResponse
// @Router /attractions/{id} [get]
func (h *PlacesHandler) GetPlaceHandler(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for getting a place by ID", slog.String("placeID", idStr))

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest, h.logger)
		h.logger.Warn("Failed to parse place ID", slog.String("placeID", idStr), slog.String("error", err.Error()))
		return
	}
	place, err := h.client.GetPlace(r.Context(), &gen.GetPlaceRequest{Id: uint32(id)})
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			response := httpresponse.ErrorResponse{
				Message: "place not found",
			}
			httpresponse.SendJSONResponse(w, response, http.StatusNotFound, h.logger)
			h.logger.Warn("Place not found", slog.String("placeID", idStr), slog.String("error", err.Error()))
			return
		}
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError, h.logger)
		h.logger.Error("Failed to get a place", slog.String("placeID", idStr), slog.String("error", err.Error()))
		return
	}
	httpresponse.SendJSONResponse(w, place.Place, http.StatusOK, h.logger)

	h.logger.DebugContext(logCtx, "Successfully getting place")
}

// GetPlacesByNameHandler godoc
// @Summary Retrieve attractions by search string
// @Description Get a list of attractions from the database that match the provided search string
// @Produce json
// @Param searchString body string true "Name of the attractions to retrieve"
// @Success 200 {object} models.GetPlace "List of attractions matching the provided searchString"
// @Failure 400 {object} httpresponses.ErrorResponse
// @Failure 500 {object} httpresponses.ErrorResponse
// @Router /attractions/search/{placeName} [get]
func (h *PlacesHandler) SearchPlacesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self';")
	placeName := r.URL.Query().Get("placeName")
	placeName = template.HTMLEscapeString(placeName)

	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for searching attractions by place name", slog.String("placeName", placeName))

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest, h.logger)
		h.logger.Warn("Invalid offset parameter", slog.String("error", err.Error()))
		return
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest, h.logger)
		h.logger.Warn("Invalid limit parameter", slog.String("error", err.Error()))
		return
	}

	cityStr := r.URL.Query().Get("city")
	categoryStr := r.URL.Query().Get("category")

	var city, category int
	if cityStr != "" {
		city, err = strconv.Atoi(cityStr)
		if err != nil {
			h.logger.Warn("Invalid city parameter", slog.String("error", err.Error()))
			httpresponse.SendJSONResponse(w, httpresponse.ErrorResponse{
				Message: "Invalid city parameter",
			}, http.StatusBadRequest, h.logger)
			return
		}
	}

	if categoryStr != "" {
		category, err = strconv.Atoi(categoryStr)
		if err != nil {
			h.logger.Warn("Invalid category parameter", slog.String("error", err.Error()))
			httpresponse.SendJSONResponse(w, httpresponse.ErrorResponse{
				Message: "Invalid category parameter",
			}, http.StatusBadRequest, h.logger)
			return
		}
	}

	places, err := h.client.SearchPlaces(r.Context(), &gen.SearchPlacesRequest{Name: placeName, Category: int32(category), City: int32(city), Limit: int32(limit), Offset: int32(offset)})
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError, h.logger)
		h.logger.Error("Failed to search attractions", slog.String("placeName", placeName), slog.String("error", err.Error()))
		return
	}
	httpresponse.SendJSONResponse(w, places.Places, http.StatusOK, h.logger)

	h.logger.DebugContext(logCtx, "Successfully getting attractions by name")
}

func (h *PlacesHandler) GetPlacesByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	categoryName := r.URL.Query().Get("category")
	categoryName = template.HTMLEscapeString(categoryName)

	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for searching attractions by category", slog.String("categoryName", categoryName))

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest, h.logger)
		h.logger.Warn("Invalid offset parameter", slog.String("error", err.Error()))
		return
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest, h.logger)
		h.logger.Warn("Invalid limit parameter", slog.String("error", err.Error()))
		return
	}

	places, err := h.client.GetPlacesByCategory(r.Context(), &gen.GetPlacesByCategoryRequest{Category: categoryName, Limit: int32(limit), Offset: int32(offset)})
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError, h.logger)
		h.logger.Error("Error getting attractions by category",
			slog.Int("limit", limit),
			slog.Int("offset", offset),
			slog.String("categoryName", categoryName),
			slog.String("error", err.Error()))
		return
	}
	httpresponse.SendJSONResponse(w, places.Places, http.StatusOK, h.logger)
}
