package delivery

import (
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

func (h *PlacesHandler) GetPlaceHandler(w http.ResponseWriter, r *http.Request) {
	places, err := h.uc.GetPlaces(r.Context())
	if err != nil {
		http.Error(w, "Не удалось получить список достопримечательностей", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(places)
	if err != nil {
		http.Error(w, "Не удалось преобразовать в json", http.StatusInternalServerError)
		return
	}
}
