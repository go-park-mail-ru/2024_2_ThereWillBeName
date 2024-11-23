package http

import (
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"2024_2_ThereWillBeName/internal/pkg/middleware"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type SurveyHandler struct {
	client gen.SurveyClient
	logger *slog.Logger
}

func NewSurveyHandler(client gen.SurveyClient, logger *slog.Logger) *SurveyHandler {
	return &SurveyHandler{
		client: client,
		logger: logger,
	}
}

func (h *SurveyHandler) GetSurveyById(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(middleware.IdKey).(uint)

	if !ok {
		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized, h.logger)
		return
	}

	vars := mux.Vars(r)
	surveyIdStr := vars["id"]

	surveyID, err := strconv.ParseUint(surveyIdStr, 10, 64)
	if err != nil {
		h.logger.Warn("Failed to parse survey ID", slog.String("surveyID", surveyIdStr), slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid survey ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}

	survey, err := h.client.GetSurveyById(r.Context(), &surveysGen.uint(surveyID))
	if err != nil {
		logCtx := log.AppendCtx(r.Context(), slog.String("surveyID", surveyIDStr))
		response, status := ErrorCheck(err, "retrieve", h.logger, logCtx)
		httpresponse.SendJSONResponse(w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully got survey by ID")

	// Send survey details as response
	httpresponse.SendJSONResponse(w, survey, http.StatusOK, h.logger)

}

func (h *SurveyHandler) CreateSurveyResponse(w http.ResponseWriter, r *http.Request) {

}

func (h *SurveyHandler) GetSurveyStatsBySurveyId(w http.ResponseWriter, r *http.Request) {

}

func (h *SurveyHandler) GetSurveyStatsByUserId(w http.ResponseWriter, r *http.Request) {

}
