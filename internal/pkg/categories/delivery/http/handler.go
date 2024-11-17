package http

import (
	"2024_2_ThereWillBeName/internal/pkg/categories"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	log "2024_2_ThereWillBeName/internal/pkg/logger"
	"log/slog"
	"net/http"
	"strconv"
)

type CategoriesHandler struct {
	uc     categories.CategoriesUsecase
	logger *slog.Logger
}

func NewCategoriesHandler(uc categories.CategoriesUsecase, logger *slog.Logger) *CategoriesHandler {
	return &CategoriesHandler{uc: uc, logger: logger}
}

func (h *CategoriesHandler) GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for getting categories")

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest, h.logger)
		h.logger.Warn("Invalid offset parameter", slog.String("error", err.Error()))
		return
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest, h.logger)
		h.logger.Warn("Invalid limit parameter", slog.String("error", err.Error()))
		return
	}

	categories, err := h.uc.GetCategories(logCtx, limit, offset)
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError, h.logger)
		h.logger.Error("Error getting categories",
			slog.Int("limit", limit),
			slog.Int("offset", offset),
			slog.String("error", err.Error()))
		return
	}
	h.logger.DebugContext(logCtx, "Successfully retrieved categories", slog.Any("categories", categories),
		slog.Int("limit", limit),
		slog.Int("offset", offset))
	httpresponse.SendJSONResponse(w, categories, http.StatusOK, h.logger)
}
