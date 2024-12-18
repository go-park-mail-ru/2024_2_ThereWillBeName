package http

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/categories/delivery/grpc/gen"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	log "2024_2_ThereWillBeName/internal/pkg/logger"

	"log/slog"
	"net/http"
	"strconv"
)

type CategoriesHandler struct {
	client gen.CategoriesClient
	logger *slog.Logger
}

func NewCategoriesHandler(client gen.CategoriesClient, logger *slog.Logger) *CategoriesHandler {
	return &CategoriesHandler{client, logger}
}

func (h *CategoriesHandler) GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()
	h.logger.DebugContext(logCtx, "Handling request for getting categories")

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		httpresponse.SendJSONResponse(logCtx, w, nil, http.StatusBadRequest, h.logger)
		h.logger.WarnContext(logCtx, "Invalid offset parameter", slog.Int("offset", offset), slog.String("error", err.Error()))
		return
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		httpresponse.SendJSONResponse(logCtx, w, nil, http.StatusBadRequest, h.logger)
		h.logger.WarnContext(logCtx, "Invalid limit parameter", slog.Int("limit", limit), slog.String("error", err.Error()))
		return
	}

	logCtx = log.AppendCtx(logCtx, slog.Int("limit", limit))
	logCtx = log.AppendCtx(logCtx, slog.Int("offset", offset))

	categories, err := h.client.GetCategories(logCtx, &gen.GetCategoriesRequest{Limit: int32(limit), Offset: int32(offset)})
	if err != nil {
		httpresponse.SendJSONResponse(logCtx, w, nil, http.StatusInternalServerError, h.logger)
		h.logger.ErrorContext(logCtx, "Error getting categories", slog.String("error", err.Error()))
		return
	}
	categoryResponse := make(models.CategoryList, len(categories.Categories))
	for i, category := range categories.Categories {
		categoryResponse[i] = models.Category{
			ID:   int(category.Id),
			Name: category.Name,
		}
	}
	h.logger.DebugContext(logCtx, "Successfully retrieved categories", slog.Any("categories", categories.Categories))
	httpresponse.SendJSONResponse(logCtx, w, categoryResponse, http.StatusOK, h.logger)
}
