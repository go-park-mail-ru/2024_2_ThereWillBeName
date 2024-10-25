package http

import (
	"2024_2_ThereWillBeName/internal/models"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"2024_2_ThereWillBeName/internal/pkg/places"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

type PlacesHandler struct {
	uc places.PlaceUsecase
}

func NewPlacesHandler(uc places.PlaceUsecase) *PlacesHandler {
	return &PlacesHandler{uc}
}

// GetPlaceHandler godoc
// @Summary Get a list of places
// @Description Retrieve a list of places from the database
// @Produce json
// @Success 200 {array} models.GetPlace "List of places"
// @Failure 400 {object} httpresponses.ErrorResponse "Bad request"
// @Failure 500 {object} httpresponses.ErrorResponse "Internal Server Error"
// @Router /places [get]
func (h *PlacesHandler) GetPlacesHandler(w http.ResponseWriter, r *http.Request) {
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest)
		logger.Println(err.Error())
		return
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest)
		logger.Println(err.Error())
		return
	}
	places, err := h.uc.GetPlaces(r.Context(), limit, offset)
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError)
		logger.Printf("Couldn't get list of places: %v", err)
		return
	}
	httpresponse.SendJSONResponse(w, places, http.StatusOK)
}

// PostPlaceHandler godoc
// @Summary Create a new place
// @Description Add a new place to the database
// @Accept json
// @Produce json
// @Param place body models.CreatePlace true "Place data"
// @Success 201 {object} httpresponses.ErrorResponse "Place successfully created"
// @Failure 400 {object} httpresponses.ErrorResponse
// @Failure 500 {object} httpresponses.ErrorResponse
// @Router /places [post]
func (h *PlacesHandler) PostPlaceHandler(w http.ResponseWriter, r *http.Request) {
	var place models.CreatePlace
	if err := json.NewDecoder(r.Body).Decode(&place); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest)
		logger.Println(err.Error())
		return
	}
	if err := h.uc.CreatePlace(r.Context(), place); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError)
		logger.Println(err.Error())
		return
	}
	httpresponse.SendJSONResponse(w, "place succesfully created", http.StatusCreated)
}

// PutPlaceHandler godoc
// @Summary Update an existing place
// @Description Update the details of an existing place in the database
// @Accept json
// @Produce json
// @Param place body models.UpdatePlace true "Updated place data"
// @Success 200 {object} httpresponses.ErrorResponse "Place successfully updated"
// @Failure 400 {object} httpresponses.ErrorResponse
// @Failure 500 {object} httpresponses.ErrorResponse
// @Router /places/{id} [put]
func (h *PlacesHandler) PutPlaceHandler(w http.ResponseWriter, r *http.Request) {
	var place models.UpdatePlace
	if err := json.NewDecoder(r.Body).Decode(&place); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest)
		logger.Println(err.Error())
		return
	}
	if err := h.uc.UpdatePlace(r.Context(), place); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError)
		logger.Println(err.Error())
		return
	}
	httpresponse.SendJSONResponse(w, "place successfully updated", http.StatusOK)
}

// DeletePlaceHandler godoc
// @Summary Delete an existing place
// @Description Remove a place from the database by its name
// @Produce json
// @Param name body string true "Name of the place to be deleted"
// @Success 200 {object} httpresponses.ErrorResponse "Place successfully deleted"
// @Failure 400 {object} httpresponses.ErrorResponse
// @Failure 500 {object} httpresponses.ErrorResponse
// @Router /places/{id} [delete]
func (h *PlacesHandler) DeletePlaceHandler(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest)
		logger.Println(err.Error())
		return
	}
	if err := h.uc.DeletePlace(r.Context(), uint(id)); err != nil {
		if errors.Is(err, models.ErrNotFound) {
			httpresponse.SendJSONResponse(w, nil, http.StatusNotFound)
			logger.Println(err.Error())
			return
		}
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError)
		logger.Println(err.Error())
		return
	}
	httpresponse.SendJSONResponse(w, "place successfully deleted", http.StatusOK)
}

// GetPlaceHandler godoc
// @Summary Retrieve an existing place
// @Description Get details of a place from the database by its id
// @Produce json
// @Param id body int true "ID of the place to retrieve"
// @Success 200 {object} models.GetPlace "Details of the requested place"
// @Failure 400 {object} httpresponses.ErrorResponse
// @Failure 500 {object} httpresponses.ErrorResponse
// @Router /places/{id} [get]
func (h *PlacesHandler) GetPlaceHandler(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest)
		logger.Println(err.Error())
		return
	}
	place, err := h.uc.GetPlace(r.Context(), uint(id))
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			response := httpresponse.ErrorResponse{
				Message: "place not found",
			}
			httpresponse.SendJSONResponse(w, response, http.StatusNotFound)
			logger.Println(err.Error())
			return
		}
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError)
		logger.Println(err.Error())
		return
	}
	httpresponse.SendJSONResponse(w, place, http.StatusOK)
}

// GetPlacesByNameHandler godoc
// @Summary Retrieve places by search string
// @Description Get a list of places from the database that match the provided search string
// @Produce json
// @Param searchString body string true "Name of the places to retrieve"
// @Success 200 {object} models.GetPlace "List of places matching the provided searchString"
// @Failure 400 {object} httpresponses.ErrorResponse
// @Failure 500 {object} httpresponses.ErrorResponse
// @Router /places/search/{placeName} [get]
func (h *PlacesHandler) SearchPlacesHandler(w http.ResponseWriter, r *http.Request) {
	placeName := mux.Vars(r)["placeName"]
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest)
		logger.Println(err.Error())
		return
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest)
		logger.Println(err.Error())
		return
	}
	places, err := h.uc.SearchPlaces(r.Context(), placeName, limit, offset)
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError)
		logger.Println(err.Error())
		return
	}
	httpresponse.SendJSONResponse(w, places, http.StatusOK)
}
