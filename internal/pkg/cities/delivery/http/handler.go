package http

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/cities"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

type CityHandler struct {
	ch cities.CityUsecase
}

func NewCityHandler(ch cities.CityUsecase) *CityHandler { return &CityHandler{ch} }

func (h *CityHandler) GetCitiesHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest)
		logger.Println(err.Error())
		return
	}
	cities, err := h.ch.GetCities(r.Context(), requestData.Limit, requestData.Offset)
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError)
		logger.Printf("Couldn't get list of cities: %v", err)
		return
	}
	logger.Printf("Got list of cities: %v", cities)
	httpresponse.SendJSONResponse(w, cities, http.StatusOK)
}

func (h *CityHandler) GetCityHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest)
		logger.Println(err.Error())
		return
	}
	city, err := h.ch.GetCity(r.Context(), requestData.ID)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			httpresponse.SendJSONResponse(w, nil, http.StatusNotFound)
			logger.Println(err.Error())
			return
		}
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError)
		logger.Printf("Couldn't get city: %v", err)
		return
	}
	httpresponse.SendJSONResponse(w, city, http.StatusOK)
}

func (h *CityHandler) PostCityHandler(w http.ResponseWriter, r *http.Request) {
	var city models.City
	if err := json.NewDecoder(r.Body).Decode(&city); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest)
		logger.Println(err.Error())
		return
	}
	logger.Println(city)
	if err := h.ch.CreateCity(r.Context(), city); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError)
		logger.Printf("Couldn't create city: %v", err)
		return
	}
	httpresponse.SendJSONResponse(w, "city successfully created", http.StatusCreated)
}

func (h *CityHandler) DeleteCityHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest)
		logger.Println(err.Error())
		return
	}
	if err := h.ch.DeleteCity(r.Context(), requestData.ID); err != nil {
		if errors.Is(err, models.ErrNotFound) {
			httpresponse.SendJSONResponse(w, nil, http.StatusNotFound)
			logger.Println(err.Error())
			return
		}
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError)
		logger.Printf("Couldn't delete city: %v", err)
		return
	}
	httpresponse.SendJSONResponse(w, "city successfully deleted", http.StatusNoContent)
}

func (h *CityHandler) PutCityHandler(w http.ResponseWriter, r *http.Request) {
	var city models.City
	if err := json.NewDecoder(r.Body).Decode(&city); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest)
		logger.Println(err.Error())
		return
	}
	if err := h.ch.UpdateCity(r.Context(), city); err != nil {
		if errors.Is(err, models.ErrNotFound) {
			httpresponse.SendJSONResponse(w, nil, http.StatusNotFound)
			logger.Println(err.Error())
			return
		}
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError)
		logger.Printf("Couldn't update city: %v", err)
		return
	}
	httpresponse.SendJSONResponse(w, "city successfully updated", http.StatusOK)
}
