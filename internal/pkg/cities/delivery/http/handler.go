package http

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/cities/delivery/grpc/gen"
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
	client gen.CitiesClient
	logger *slog.Logger
}

func NewCitiesHandler(client gen.CitiesClient, logger *slog.Logger) *CitiesHandler {
	return &CitiesHandler{client, logger}
}

func ErrorCheck(err error, action string, logger *slog.Logger, ctx context.Context) (httpresponse.ErrorResponse, int) {
	logContext := log.AppendCtx(ctx, slog.String("action", action))
	logContext = log.AppendCtx(logContext, slog.Any("error", err.Error()))

	if errors.Is(err, models.ErrNotFound) {
		logger.ErrorContext(logContext, fmt.Sprintf("Error during %s operation", action))

		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		return response, http.StatusNotFound
	}
	logger.ErrorContext(logContext, fmt.Sprintf("Failed to %s cities", action))

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
// @Failure 403 {object} httpresponses.ErrorResponse "Token is missing"
// @Failure 403 {object} httpresponses.ErrorResponse "Invalid token"
// @Failure 404 {object} httpresponses.ErrorResponse "Cities not found"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to retrieve cities"
// @Router /cities/search [get]
func (h *CitiesHandler) SearchCitiesByNameHandler(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()

	query := r.URL.Query().Get("q")
	logCtx = log.AppendCtx(logCtx, slog.String("name", query))
	h.logger.DebugContext(logCtx, "Handling request for searching cities by name")

	if query == "" {
		h.logger.WarnContext(logCtx, "Search query parameter is empty")
		response := httpresponse.ErrorResponse{
			Message: "Invalid query",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	cities, err := h.client.SearchCitiesByName(context.Background(), &gen.SearchCitiesByNameRequest{Query: query})
	if err != nil {
		response, status := ErrorCheck(err, "retrieve", h.logger, logCtx)
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully retrieved cities", slog.Any("cities", cities.Cities))
	httpresponse.SendJSONResponse(logCtx, w, cities.Cities, http.StatusOK, h.logger)
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
	logCtx := r.Context()

	vars := mux.Vars(r)
	cityIDStr := vars["id"]

	logCtx = log.AppendCtx(logCtx, slog.String("city_id", cityIDStr))

	h.logger.DebugContext(logCtx, "Handling request for searching cities by ID")

	if cityIDStr == "" {
		h.logger.WarnContext(logCtx, "Search ID parameter is empty")

		response := httpresponse.ErrorResponse{
			Message: "Invalid city ID",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	cityID, err := strconv.ParseUint(cityIDStr, 10, 32)
	if err != nil {
		h.logger.WarnContext(logCtx, "Invalid city ID", slog.Any("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid city ID",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	city, err := h.client.SearchCityByID(context.Background(), &gen.SearchCityByIDRequest{Id: uint32(cityID)})
	if err != nil {
		response, status := ErrorCheck(err, "retrieve", h.logger, logCtx)
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully retrieved city by ID", slog.Any("city", city.City))

	httpresponse.SendJSONResponse(logCtx, w, city.City, http.StatusOK, h.logger)
}
