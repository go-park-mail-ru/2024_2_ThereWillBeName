package delivery

import (
	"2024_2_ThereWillBeName/internal/models"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"2024_2_ThereWillBeName/internal/pkg/places"
	"encoding/json"
	"net/http"
)

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
		http.Error(w, "Не удалось получить список достопримечательностей", http.StatusInternalServerError)
		return
	}
	httpresponse.SendJSONResponse(w, places, http.StatusOK)
}

func (h *PlacesHandler) PostPlaceHandler(w http.ResponseWriter, r *http.Request) {
	var place models.Place
	if err := json.NewDecoder(r.Body).Decode(&place); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.uc.CreatePlace(r.Context(), place); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	httpresponse.SendJSONResponse(w, "place succesfully created", http.StatusCreated)
}

func (h *PlacesHandler) PutPlaceHandler(w http.ResponseWriter, r *http.Request) {
	var place models.Place
	if err := json.NewDecoder(r.Body).Decode(&place); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.uc.UpdatePlace(r.Context(), place); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httpresponse.SendJSONResponse(w, "place successfully updated", http.StatusOK)
}

func (h *PlacesHandler) DeletePlaceHandler(w http.ResponseWriter, r *http.Request) {
	var name string
	if err := json.NewDecoder(r.Body).Decode(&name); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.uc.DeletePlace(r.Context(), name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httpresponse.SendJSONResponse(w, "place successfully deleted", http.StatusOK)
}

func (h *PlacesHandler) GetPlaceHandler(w http.ResponseWriter, r *http.Request) {
	var name string
	if err := json.NewDecoder(r.Body).Decode(&name); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	place, err := h.uc.ReadPlace(r.Context(), name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httpresponse.SendJSONResponse(w, place, http.StatusOK)
}
