//go:generate easyjson .

package httpresponses

import (
	"context"
	"github.com/mailru/easyjson"
	"log/slog"
	"net/http"
)

//easyjson:json
type Response struct {
	Message string `json:"message"`
}

func SendJSONResponse(logCtx context.Context, w http.ResponseWriter, data easyjson.Marshaler, status int, logger *slog.Logger) {
	w.WriteHeader(status)

	if data == nil {
		logger.WarnContext(logCtx, "Received nil data for JSON response")
		http.Error(w, "No data to send", http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, _, err := easyjson.MarshalToHTTPResponseWriter(data, w); err != nil {
		logger.ErrorContext(logCtx, "Failed to encode response to JSON", slog.Any("error", err.Error()))
		http.Error(w, "Failed to convert to json", http.StatusInternalServerError)
	}
}
