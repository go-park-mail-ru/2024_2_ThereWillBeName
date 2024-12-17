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
	logCtx := r.Context()
	h.logger.DebugContext(logCtx, "Handling request for global searching places and cities")

	query := r.URL.Query().Get("query")
	if query == "" {
		h.logger.WarnContext(logCtx, "Query parameter can't be empty")
		httpresponses.SendJSONResponse(logCtx, w, nil, http.StatusBadRequest, h.logger)
		return
	}

	decodedQuery, err := url.QueryUnescape(query)
	if err != nil {
		h.logger.WarnContext(logCtx, "Error decoding query", slog.String("error", err.Error()))
		httpresponses.SendJSONResponse(logCtx, w, nil, http.StatusBadRequest, h.logger)
		return
	}

	logCtx = log.AppendCtx(logCtx, slog.String("decoded_query", decodedQuery))

	results, err := h.client.Search(r.Context(), &searchGen.SearchRequest{DecodedQuery: decodedQuery})
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			h.logger.WarnContext(logCtx, "No results found for query")

			httpresponses.SendJSONResponse(logCtx, w, nil, http.StatusNotFound, h.logger)
		}

		h.logger.ErrorContext(logCtx, "Failed to search", slog.String("error", err.Error()))

		httpresponses.SendJSONResponse(logCtx, w, nil, http.StatusInternalServerError, h.logger)
		return
	}
	h.logger.DebugContext(logCtx, "Search completed successfully", slog.Int("resultCount", len(results.SearchResult)))

	searchResponse := make(models.SearchResultList, len(results.SearchResult))
	for i, res := range results.SearchResult {
		searchResponse[i] = models.SearchResult{
			Id:   uint(res.Id),
			Name: res.Name,
			Type: res.Type,
		}
	}
	httpresponses.SendJSONResponse(logCtx, w, searchResponse, http.StatusOK, h.logger)
}
