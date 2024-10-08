package delivery

import (
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"2024_2_ThereWillBeName/internal/pkg/places"
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
func (h *PlacesHandler) GetPlaceHandler(w http.ResponseWriter, r *http.Request) {
	places, err := h.uc.GetPlaces(r.Context())
	if err != nil {
		http.Error(w, "Не удалось получить список достопримечательностей", http.StatusInternalServerError)
		return
	}
	httpresponse.SendJSONResponse(w, places, http.StatusOK)
}
