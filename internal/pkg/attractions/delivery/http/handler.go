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
// @Failure 400 {object} httpresponses.Response "Bad request"
// @Failure 500 {object} httpresponses.Response "Internal Server Error"
// @Router /attractions [get]
func (h *PlacesHandler) GetPlacesHandler(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()
	h.logger.DebugContext(logCtx, "Handling request for searching attractions")

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		httpresponse.SendJSONResponse(logCtx, w, nil, http.StatusBadRequest, h.logger)
		h.logger.WarnContext(logCtx, "Invalid offset parameter", slog.String("error", err.Error()))
		return
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		httpresponse.SendJSONResponse(logCtx, w, nil, http.StatusBadRequest, h.logger)
		h.logger.WarnContext(logCtx, "Invalid limit parameter", slog.String("error", err.Error()))
		return
	}
	places, err := h.client.GetPlaces(r.Context(), &gen.GetPlacesRequest{Limit: int32(limit), Offset: int32(offset)})
	if err != nil {
		httpresponse.SendJSONResponse(logCtx, w, nil, http.StatusInternalServerError, h.logger)
		h.logger.ErrorContext(logCtx, "Couldn't get list of attractions",
			slog.Int("limit", limit),
			slog.Int("offset", offset),
			slog.String("error", err.Error()))
		return
	}

	h.logger.DebugContext(logCtx, "Successfully retrieved attractions", slog.Int("attractions_count", len(places.Places)))

	placeResponse := make(models.GetPLaceList, len(places.Places))
	for i, place := range places.Places {
		placeResponse[i] = models.GetPlace{
			ID:          int(place.Id),
			Name:        place.Name,
			ImagePath:   place.ImagePath,
			Description: place.Description,
			Rating:      int(place.Rating),
			Address:     place.Address,
			City:        place.City,
			PhoneNumber: place.PhoneNumber,
			Categories:  place.Categories,
			Latitude:    place.Latitude,
			Longitude:   place.Longitude,
		}
	}
	httpresponse.SendJSONResponse(logCtx, w, placeResponse, http.StatusOK, h.logger)
}

// PostPlaceHandler godoc
// @Summary Create a new place
// @Description Add a new place to the database
// @Accept json
// @Produce json
// @Param place body models.CreatePlace true "Place data"
// @Success 201 {object} httpresponses.Response "Place successfully created"
// @Failure 400 {object} httpresponses.Response
// @Failure 403 {object} httpresponses.Response "Token is missing"
// @Failure 403 {object} httpresponses.Response "Invalid token"
// @Failure 422 {object} httpresponses.Response
// @Failure 500 {object} httpresponses.Response
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
//// @Success 200 {object} httpresponses.Response "Place successfully updated"
//// @Failure 400 {object} httpresponses.Response
//// @Failure 403 {object} httpresponses.Response "Token is missing"
//// @Failure 403 {object} httpresponses.Response "Invalid token"
//// @Failure 422 {object} httpresponses.Response
//// @Failure 500 {object} httpresponses.Response
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
//// @Success 200 {object} httpresponses.Response "Place successfully deleted"
//// @Failure 400 {object} httpresponses.Response
//// @Failure 403 {object} httpresponses.Response "Token is missing"
//// @Failure 403 {object} httpresponses.Response "Invalid token"
//// @Failure 500 {object} httpresponses.Response
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
// @Failure 400 {object} httpresponses.Response
// @Failure 500 {object} httpresponses.Response
// @Router /attractions/{id} [get]
func (h *PlacesHandler) GetPlaceHandler(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()

	idStr := mux.Vars(r)["id"]

	logCtx = log.AppendCtx(logCtx, slog.String("placeID", idStr))

	h.logger.DebugContext(logCtx, "Handling request for getting a place by ID")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		httpresponse.SendJSONResponse(logCtx, w, nil, http.StatusBadRequest, h.logger)
		h.logger.WarnContext(logCtx, "Failed to parse place ID", slog.String("error", err.Error()))
		return
	}
	place, err := h.client.GetPlace(r.Context(), &gen.GetPlaceRequest{Id: uint32(id)})
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			response := httpresponse.Response{
				Message: "place not found",
			}
			httpresponse.SendJSONResponse(logCtx, w, response, http.StatusNotFound, h.logger)
			h.logger.ErrorContext(logCtx, "Place not found", slog.Any("error", err.Error()))
			return
		}
		httpresponse.SendJSONResponse(logCtx, w, nil, http.StatusInternalServerError, h.logger)
		h.logger.ErrorContext(logCtx, "Failed to get a place", slog.Any("error", err.Error()))
		return
	}

	responsePlace := models.GetPlace{
		ID:              int(place.Place.Id),
		Name:            place.Place.Name,
		ImagePath:       place.Place.ImagePath,
		Description:     place.Place.Description,
		Rating:          float32(place.Place.Rating),
		NumberOfReviews: int(place.Place.NumberOfReviews),
		Address:         place.Place.Address,
		City:            place.Place.City,
		PhoneNumber:     place.Place.PhoneNumber,
		Categories:      place.Place.Categories,
		Latitude:        float32(place.Place.Latitude),
		Longitude:       float32(place.Place.Longitude),
	}

	httpresponse.SendJSONResponse(logCtx, w, responsePlace, http.StatusOK, h.logger)

	h.logger.DebugContext(logCtx, "Successfully got place by ID")
}

// GetPlacesByNameHandler godoc
// @Summary Retrieve attractions by search string
// @Description Get a list of attractions from the database that match the provided search string
// @Produce json
// @Param searchString body string true "Name of the attractions to retrieve"
// @Success 200 {object} models.GetPlace "List of attractions matching the provided searchString"
// @Failure 400 {object} httpresponses.Response
// @Failure 500 {object} httpresponses.Response
// @Router /attractions/search/{placeName} [get]
func (h *PlacesHandler) SearchPlacesHandler(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()

	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self';")
	placeName := r.URL.Query().Get("placeName")
	placeName = template.HTMLEscapeString(placeName)

	logCtx = log.AppendCtx(logCtx, slog.String("placename", placeName))
	h.logger.DebugContext(logCtx, "Handling request for searching attractions by place name")

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		httpresponse.SendJSONResponse(logCtx, w, nil, http.StatusBadRequest, h.logger)
		h.logger.WarnContext(logCtx, "Invalid offset parameter", slog.Int("offset", offset), slog.String("error", err.Error()))
		return
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		httpresponse.SendJSONResponse(logCtx, w, nil, http.StatusBadRequest, h.logger)
		h.logger.Warn("Invalid limit parameter", slog.Int("limit", limit), slog.String("error", err.Error()))
		return
	}

	cityStr := r.URL.Query().Get("city")
	categoryStr := r.URL.Query().Get("category")
	filterStr := r.URL.Query().Get("filter")

	var city, category, filterType int
	if cityStr != "" {
		city, err = strconv.Atoi(cityStr)
		if err != nil {
			h.logger.WarnContext(logCtx, "Invalid city parameter", slog.String("error", err.Error()))
			httpresponse.SendJSONResponse(logCtx, w, httpresponse.Response{
				Message: "Invalid city parameter",
			}, http.StatusBadRequest, h.logger)
			return
		}
	}

	if categoryStr != "" {
		category, err = strconv.Atoi(categoryStr)
		if err != nil {
			h.logger.WarnContext(logCtx, "Invalid category parameter", slog.String("error", err.Error()))
			httpresponse.SendJSONResponse(logCtx, w, httpresponse.Response{
				Message: "Invalid category parameter",
			}, http.StatusBadRequest, h.logger)
			return
		}
	}

	if filterStr != "" {
		filterType, err = strconv.Atoi(filterStr)
		if err != nil {
			h.logger.Warn("Invalid filter parameter", slog.String("error", err.Error()))
			httpresponse.SendJSONResponse(logCtx, w, httpresponse.ErrorResponse{
				Message: "Invalid filter parameter",
			}, http.StatusBadRequest, h.logger)
			return
		}
	}

	places, err := h.client.SearchPlaces(r.Context(), &gen.SearchPlacesRequest{Name: placeName, Category: int32(category), City: int32(city), FilterType: int32(filterType), Limit: int32(limit), Offset: int32(offset)})
	if err != nil {
		httpresponse.SendJSONResponse(logCtx, w, nil, http.StatusInternalServerError, h.logger)
		h.logger.ErrorContext(logCtx, "Failed to search attractions", slog.String("error", err.Error()))
		return
	}

	placeResponse := make(models.GetPLaceList, len(places.Places))
	for i, place := range places.Places {
		responsePlaces[i] = models.GetPlace{
			ID:              int(place.Id),
			Name:            place.Name,
			ImagePath:       place.ImagePath,
			Description:     place.Description,
			Rating:          float32(place.Rating),
			NumberOfReviews: int(place.NumberOfReviews),
			Address:         place.Address,
			City:            place.City,
			PhoneNumber:     place.PhoneNumber,
			Categories:      place.Categories,
		}
	}

	httpresponse.SendJSONResponse(logCtx, w, responsePlaces, http.StatusOK, h.logger)
	h.logger.DebugContext(logCtx, "Successfully getting attractions by name", slog.Int("places_count", len(places.Places)))

}

func (h *PlacesHandler) GetPlacesByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()

	categoryName := r.URL.Query().Get("category")
	categoryName = template.HTMLEscapeString(categoryName)

	logCtx = log.AppendCtx(logCtx, slog.String("category_name", categoryName))

	h.logger.DebugContext(logCtx, "Handling request for searching attractions by category")

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		httpresponse.SendJSONResponse(logCtx, w, nil, http.StatusBadRequest, h.logger)
		h.logger.WarnContext(logCtx, "Invalid offset parameter", slog.Int("offset", offset), slog.String("error", err.Error()))
		return
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		httpresponse.SendJSONResponse(logCtx, w, nil, http.StatusBadRequest, h.logger)
		h.logger.WarnContext(logCtx, "Invalid limit parameter", slog.Int("limit", limit), slog.String("error", err.Error()))
		return
	}

	places, err := h.client.GetPlacesByCategory(r.Context(), &gen.GetPlacesByCategoryRequest{Category: categoryName, Limit: int32(limit), Offset: int32(offset)})
	if err != nil {
		httpresponse.SendJSONResponse(logCtx, w, nil, http.StatusInternalServerError, h.logger)
		h.logger.ErrorContext(logCtx, "Error getting attractions by category",
			slog.Int("limit", limit),
			slog.Int("offset", offset),
			slog.String("error", err.Error()))
		return
	}
	placeResponse := make(models.GetPLaceList, len(places.Places))
	for i, place := range places.Places {
		responsePlaces[i] = models.GetPlace{
			ID:              int(place.Id),
			Name:            place.Name,
			ImagePath:       place.ImagePath,
			Description:     place.Description,
			Rating:          float32(place.Rating),
			NumberOfReviews: int(place.NumberOfReviews),
			Address:         place.Address,
			City:            place.City,
			PhoneNumber:     place.PhoneNumber,
			Categories:      place.Categories,
		}
	}
	h.logger.DebugContext(logCtx, "Successfully got attractions by category name")

	httpresponse.SendJSONResponse(logCtx, w, placeResponse, http.StatusOK, h.logger)
}
