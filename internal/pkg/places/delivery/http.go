package delivery

import (
	"2024_2_ThereWillBeName/internal/models"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"2024_2_ThereWillBeName/internal/pkg/places"
	"encoding/json"
	"log"
	"net/http"
	"os"
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
// @Success 200 {array} models.Place "List of places"
// @Failure 500 {string} string
// @Router /places [get]
func (h *PlacesHandler) GetPlacesHandler(w http.ResponseWriter, r *http.Request) {
	places, err := h.uc.GetPlaces(r.Context())
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError)
		logger.Printf("Не удалось получить список достопримечательностей: %v", err)
		return
	}
	httpresponse.SendJSONResponse(w, places, http.StatusOK)
}

// PostPlaceHandler godoc
// @Summary Create a new place
// @Description Add a new place to the database
// @Accept json
// @Produce json
// @Param place body models.Place true "Place data"
// @Success 201 {string} string "Place successfully created"
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /create [post]
func (h *PlacesHandler) PostPlaceHandler(w http.ResponseWriter, r *http.Request) {
	var place models.Place
	if err := json.NewDecoder(r.Body).Decode(&place); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest)
		logger.Printf(err.Error())
		return
	}
	if err := h.uc.CreatePlace(r.Context(), place); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError)
		logger.Printf(err.Error())
		return
	}
	httpresponse.SendJSONResponse(w, "place succesfully created", http.StatusCreated)
}

// PutPlaceHandler godoc
// @Summary Update an existing place
// @Description Update the details of an existing place in the database
// @Accept json
// @Produce json
// @Param place body models.Place true "Updated place data"
// @Success 200 {string} string "Place successfully updated"
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /update/{id} [put]
func (h *PlacesHandler) PutPlaceHandler(w http.ResponseWriter, r *http.Request) {
	var place models.Place
	if err := json.NewDecoder(r.Body).Decode(&place); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest)
		logger.Printf(err.Error())
		return
	}
	if err := h.uc.UpdatePlace(r.Context(), place); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError)
		logger.Printf(err.Error())
		return
	}
	httpresponse.SendJSONResponse(w, "place successfully updated", http.StatusOK)
}

// DeletePlaceHandler godoc
// @Summary Delete an existing place
// @Description Remove a place from the database by its name
// @Accept json
// @Produce json
// @Param name body string true "Name of the place to be deleted"
// @Success 200 {string} string "Place successfully deleted"
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /delete/{id} [delete]
func (h *PlacesHandler) DeletePlaceHandler(w http.ResponseWriter, r *http.Request) {
	var name string
	if err := json.NewDecoder(r.Body).Decode(&name); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest)
		logger.Printf(err.Error())
		return
	}
	if err := h.uc.DeletePlace(r.Context(), name); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError)
		logger.Printf(err.Error())
		return
	}
	httpresponse.SendJSONResponse(w, "place successfully deleted", http.StatusOK)
}

// GetPlaceHandler godoc
// @Summary Retrieve an existing place
// @Description Get details of a place from the database by its name
// @Accept json
// @Produce json
// @Param name body string true "Name of the place to retrieve"
// @Success 200 {object} models.Place "Details of the requested place"
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /place/{id} [get]
func (h *PlacesHandler) GetPlaceHandler(w http.ResponseWriter, r *http.Request) {
	var name string
	if err := json.NewDecoder(r.Body).Decode(&name); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest)
		logger.Printf(err.Error())
		return
	}
	place, err := h.uc.GetPlace(r.Context(), name)
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError)
		logger.Printf(err.Error())
		return
	}
	httpresponse.SendJSONResponse(w, place, http.StatusOK)
}

// GetPlacesByNameHandler godoc
// @Summary Retrieve places by search string
// @Description Get a list of places from the database that match the provided search string
// @Accept json
// @Produce json
// @Param searchString body string true "Name of the places to retrieve"
// @Success 200 {array} models.Place "List of places matching the provided searchString"
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /search [get]
func (h *PlacesHandler) GetPlacesBySearchHandler(w http.ResponseWriter, r *http.Request) {
	var searchString string
	if err := json.NewDecoder(r.Body).Decode(&searchString); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest)
		logger.Printf(err.Error())
		return
	}
	places, err := h.uc.GetPlacesBySearch(r.Context(), searchString)
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError)
		logger.Printf(err.Error())
		return
	}
	httpresponse.SendJSONResponse(w, places, http.StatusOK)
}
