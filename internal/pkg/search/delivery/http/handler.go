package http

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"2024_2_ThereWillBeName/internal/pkg/search"
	"errors"
	"log/slog"
	"net/http"
	"net/url"
)

type SearchHandler struct {
	uc     search.SearchUsecase
	logger *slog.Logger
}

func NewSearchHandler(uc search.SearchUsecase, logger *slog.Logger) *SearchHandler {
	return &SearchHandler{uc, logger}

}

func (h *SearchHandler) SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		httpresponses.SendJSONResponse(w, nil, http.StatusBadRequest, h.logger)
		return
	}

	decodedQuery, err := url.QueryUnescape(query)
	if err != nil {
		httpresponses.SendJSONResponse(w, nil, http.StatusBadRequest, h.logger)
		return
	}

	results, err := h.uc.Search(r.Context(), decodedQuery)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			httpresponses.SendJSONResponse(w, nil, http.StatusNotFound, h.logger)
		}
		httpresponses.SendJSONResponse(w, nil, http.StatusInternalServerError, h.logger)
		return
	}

	httpresponses.SendJSONResponse(w, results, http.StatusOK, h.logger)
}
