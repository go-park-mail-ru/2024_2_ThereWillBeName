package httpresponses

import (
	"encoding/json"
	"net/http"
)

func SendJSONResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		http.Error(w, "Нет данных для отправки", http.StatusNoContent)
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Не удалось преобразовать в json", http.StatusInternalServerError)
	}
}
