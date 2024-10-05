package httpresponses

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func SendJSONResponse(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)

	if data == nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Не удалось преобразовать в json", http.StatusInternalServerError)
	}
}
