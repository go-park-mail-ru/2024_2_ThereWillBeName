package httpresponses

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func SendJSONResponse(logCtx context.Context, w http.ResponseWriter, data interface{}, status int, logger *slog.Logger) {
	w.WriteHeader(status)

	if data == nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.ErrorContext(logCtx, "Failed to encode response to JSON", slog.Any("error", err.Error()))

		http.Error(w, "Failed to convert to json", http.StatusInternalServerError)
	}
}
