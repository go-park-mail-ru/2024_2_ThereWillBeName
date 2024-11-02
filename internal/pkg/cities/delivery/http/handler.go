package http

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/cities"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CitiesHandler struct {
	uc cities.CitiesUsecase
}

func NewCitiesHandler(uc cities.CitiesUsecase) *CitiesHandler {
	return &CitiesHandler{uc}
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
		Message: fmt.Sprintf("Failed to %s cities", action),
	}
	return response, http.StatusInternalServerError
}

// SearchCitiesByNameHandler godoc
// @Summary Retrieve cities by name
// @Description Get cities details by city name
// @Produce json
// @Success 200 {array} models.City "Cities details"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid query"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to retrieve cities"
// @Router /cities/search [get]
func (h *CitiesHandler) SearchCitiesByNameHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		response := httpresponse.ErrorResponse{
			Message: "Invalid query",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	cities, err := h.uc.SearchCitiesByName(context.Background(), query)
	if err != nil {
		response, status := ErrorCheck(err, "retieve")
		httpresponse.SendJSONResponse(w, response, status)
		return
	}

	httpresponse.SendJSONResponse(w, cities, http.StatusOK)
}

// SearchCityByIDHandler godoc
// @Summary Retrieve a city by ID
// @Description Get city details by city ID
// @Produce json
// @Param id path int true "City ID"
// @Success 200 {object} models.City "City details"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid city ID"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to retrieve cities"
// @Router /cities/{id} [get]
func (h *CitiesHandler) SearchCityByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cityIDStr := vars["id"]
	if cityIDStr == "" {
		response := httpresponse.ErrorResponse{
			Message: "Invalid city ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	cityID, err := strconv.ParseUint(cityIDStr, 10, 32)
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Invalid city ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	city, err := h.uc.SearchCityByID(context.Background(), uint(cityID))
	if err != nil {
		response, status := ErrorCheck(err, "retrieve")
		httpresponse.SendJSONResponse(w, response, status)
		return
	}

	httpresponse.SendJSONResponse(w, city, http.StatusOK)
}
