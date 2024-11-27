package http

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/httpresponses"
	log "2024_2_ThereWillBeName/internal/pkg/logger"
	searchGen "2024_2_ThereWillBeName/internal/pkg/search/delivery/grpc/gen"
	"errors"
	"log/slog"
	"net/http"
	"net/url"
)

type SearchHandler struct {
	client searchGen.SearchClient
	logger *slog.Logger
}

func NewSearchHandler(client searchGen.SearchClient, logger *slog.Logger) *SearchHandler {
	return &SearchHandler{client, logger}

}

func (h *SearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for global searching places and cities")

	query := r.URL.Query().Get("query")
	if query == "" {
		h.logger.Warn("Query parameter can't be empty")
		httpresponses.SendJSONResponse(w, nil, http.StatusBadRequest, h.logger)
		return
	}

	decodedQuery, err := url.QueryUnescape(query)
	if err != nil {
		h.logger.Warn("Error decoding query", slog.String("error", err.Error()))
		httpresponses.SendJSONResponse(w, nil, http.StatusBadRequest, h.logger)
		return
	}

	results, err := h.client.Search(r.Context(), &searchGen.SearchRequest{DecodedQuery: decodedQuery})
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			h.logger.WarnContext(logCtx, "No results found for query", slog.String("decodedQuery", decodedQuery))

			httpresponses.SendJSONResponse(w, nil, http.StatusNotFound, h.logger)
		}

		h.logger.ErrorContext(logCtx, "Failed to search", slog.String("decodedQuery", decodedQuery), slog.String("error", err.Error()))

		httpresponses.SendJSONResponse(w, nil, http.StatusInternalServerError, h.logger)
		return
	}
	h.logger.InfoContext(logCtx, "Search completed successfully", slog.Int("resultCount", len(results.SearchResult)))

	httpresponses.SendJSONResponse(w, results.SearchResult, http.StatusOK, h.logger)
}
