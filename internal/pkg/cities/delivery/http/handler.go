package http

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/cities"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	log "2024_2_ThereWillBeName/internal/pkg/logger"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CitiesHandler struct {
	uc     cities.CitiesUsecase
	logger *slog.Logger
}

func NewCitiesHandler(uc cities.CitiesUsecase, logger *slog.Logger) *CitiesHandler {
	return &CitiesHandler{uc, logger}
}

func ErrorCheck(err error, action string, logger *slog.Logger, ctx context.Context) (httpresponse.ErrorResponse, int) {
	if errors.Is(err, models.ErrNotFound) {
		logContext := log.AppendCtx(ctx, slog.String("action", action))
		logger.ErrorContext(logContext, fmt.Sprintf("Error during %s operation", action), slog.Any("error", err.Error()))

		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		return response, http.StatusNotFound
	}
	logContext := log.AppendCtx(ctx, slog.String("action", action))
	logger.ErrorContext(logContext, fmt.Sprintf("Failed to %s cities", action), slog.Any("error", err.Error()))

	response := httpresponse.ErrorResponse{
		Message: fmt.Sprintf("Failed to %s cities", action),
	}
	return response, http.StatusInternalServerError
}

// SearchCitiesByNameHandler godoc
// @Summary Retrieve cities by name
// @Description Get cities details by city name
// @Produce json
// @Success 200 {array} models.City "Cities details"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid query"
// @Failure 403 {object} httpresponses.ErrorResponse "CSRF token missing"
// @Failure 403 {object} httpresponses.ErrorResponse "Invalid CSRF token"
// @Failure 404 {object} httpresponses.ErrorResponse "Cities not found"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to retrieve cities"
// @Router /cities/search [get]
func (h *CitiesHandler) SearchCitiesByNameHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for searching places by name", slog.String("name", query))

	if query == "" {
		h.logger.Error("Search query parameter is empty")
		response := httpresponse.ErrorResponse{
			Message: "Invalid query",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}

	cities, err := h.uc.SearchCitiesByName(context.Background(), query)
	if err != nil {
		logCtx := log.AppendCtx(context.Background(), slog.String("name", query))
		response, status := ErrorCheck(err, "retrieve", h.logger, logCtx)
		httpresponse.SendJSONResponse(w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully retrieved cities")
	httpresponse.SendJSONResponse(w, cities, http.StatusOK, h.logger)
}

// SearchCityByIDHandler godoc
// @Summary Retrieve a city by ID
// @Description Get city details by city ID
// @Produce json
// @Param id path int true "City ID"
// @Success 200 {object} models.City "City details"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid city ID"
// @Failure 403 {object} httpresponses.ErrorResponse "CSRF token missing"
// @Failure 403 {object} httpresponses.ErrorResponse "Invalid CSRF token"
// @Failure 404 {object} httpresponses.ErrorResponse "City not found"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to retrieve cities"
// @Router /cities/{id} [get]
func (h *CitiesHandler) SearchCityByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cityIDStr := vars["id"]

	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for searching cities by ID", slog.String("ID", cityIDStr))

	if cityIDStr == "" {
		h.logger.Error("Search ID parameter is empty")

		response := httpresponse.ErrorResponse{
			Message: "Invalid city ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}

	cityID, err := strconv.ParseUint(cityIDStr, 10, 32)
	if err != nil {
		h.logger.Error("Invalid city ID")
		response := httpresponse.ErrorResponse{
			Message: "Invalid city ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}

	city, err := h.uc.SearchCityByID(context.Background(), uint(cityID))
	if err != nil {
		logCtx := log.AppendCtx(context.Background(), slog.Uint64("ID", (cityID)))
		response, status := ErrorCheck(err, "retrieve", h.logger, logCtx)
		httpresponse.SendJSONResponse(w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully retrieved cities")

	httpresponse.SendJSONResponse(w, city, http.StatusOK, h.logger)
}
