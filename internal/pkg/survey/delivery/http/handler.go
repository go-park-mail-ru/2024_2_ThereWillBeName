package http

import (
	"2024_2_ThereWillBeName/internal/models"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	log "2024_2_ThereWillBeName/internal/pkg/logger"
	"2024_2_ThereWillBeName/internal/pkg/middleware"
	surveysGen "2024_2_ThereWillBeName/internal/pkg/survey/delivery/grpc/gen"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type SurveyHandler struct {
	client surveysGen.SurveyServiceClient
	logger *slog.Logger
}

func NewSurveyHandler(client surveysGen.SurveyServiceClient, logger *slog.Logger) *SurveyHandler {
	return &SurveyHandler{
		client: client,
		logger: logger,
	}
}

func ErrorCheck(err error, action string, logger *slog.Logger, ctx context.Context) (httpresponse.Response, int) {
	if errors.Is(err, models.ErrNotFound) {

		logContext := log.AppendCtx(ctx, slog.String("action", action))
		logger.ErrorContext(logContext, fmt.Sprintf("Error during %s operation", action), slog.Any("error", err.Error()))

		response := httpresponse.Response{
			Message: "Invalid request",
		}
		return response, http.StatusNotFound
	}
	logContext := log.AppendCtx(ctx, slog.String("action", action))
	logger.ErrorContext(logContext, fmt.Sprintf("Failed to %s survey", action), slog.Any("error", err.Error()))
	response := httpresponse.Response{
		Message: fmt.Sprintf("Failed to %s survey", action),
	}
	return response, http.StatusInternalServerError
}

func (h *SurveyHandler) GetSurveyById(w http.ResponseWriter, r *http.Request) {
	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for getting a survey")

	_, ok := r.Context().Value(middleware.IdKey).(uint)

	if !ok {
		response := httpresponse.Response{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	vars := mux.Vars(r)
	surveyIdStr := vars["id"]

	surveyId, err := strconv.ParseUint(surveyIdStr, 10, 64)
	if err != nil {
		h.logger.Warn("Failed to parse survey ID", slog.String("surveyID", surveyIdStr), slog.String("error", err.Error()))
		response := httpresponse.Response{
			Message: "Invalid survey ID",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	survey, err := h.client.GetSurveyById(r.Context(), &surveysGen.GetSurveyByIdRequest{Id: uint32(surveyId)})
	if err != nil {
		logCtx := log.AppendCtx(r.Context(), slog.String("surveyId", surveyIdStr))
		response, status := ErrorCheck(err, "retrieve", h.logger, logCtx)
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully got survey by ID")

	surveyResponse := models.Survey{
		Id:         uint(survey.Survey.Id),
		SurveyText: survey.Survey.SurveyText,
		MaxRating:  int(survey.Survey.MaxRating),
	}

	httpresponse.SendJSONResponse(logCtx, w, surveyResponse, http.StatusOK, h.logger)
}

func (h *SurveyHandler) CreateSurveyResponse(w http.ResponseWriter, r *http.Request) {
	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request to  create survey response")

	_, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {
		response := httpresponse.Response{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	var surveyResponse models.SurveyResponse
	err := json.NewDecoder(r.Body).Decode(&surveyResponse)
	if err != nil {
		h.logger.Warn("Failed to decode survey response", slog.String("error", err.Error()))
		response := httpresponse.Response{
			Message: "Invalid request body",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	_, err = h.client.CreateSurvey(r.Context(), &surveysGen.CreateSurveyRequest{ServeyResponce: &surveysGen.SurveyResponce{
		SurveyId: uint32(surveyResponse.SurveyId),
		UserId:   uint32(surveyResponse.UserId),
		Rating:   uint32(surveyResponse.Rating),
	}})
	if err != nil {
		logCtx := log.AppendCtx(r.Context(), slog.String("surveyID", fmt.Sprint(surveyResponse.SurveyId)))
		response, status := ErrorCheck(err, "submit", h.logger, logCtx)
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully submitted survey response")

	httpresponse.SendJSONResponse(logCtx, w, nil, http.StatusOK, h.logger)
}

func (h *SurveyHandler) GetSurveyStatsBySurveyId(w http.ResponseWriter, r *http.Request) {
	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for survey statistics by survey ID")

	_, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {
		response := httpresponse.Response{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	vars := mux.Vars(r)
	surveyIdStr := vars["id"]

	surveyId, err := strconv.ParseUint(surveyIdStr, 10, 64)
	if err != nil {
		h.logger.Warn("Failed to parse survey ID", slog.String("surveyID", surveyIdStr), slog.String("error", err.Error()))
		response := httpresponse.Response{
			Message: "Invalid survey ID",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	stats, err := h.client.GetSurveyStatsBySurveyId(r.Context(), &surveysGen.GetSurveyStatsBySurveyIdRequest{
		Id: uint32(surveyId),
	})
	if err != nil {
		logCtx := log.AppendCtx(r.Context(), slog.String("surveyId", surveyIdStr))
		response, status := ErrorCheck(err, "retrieve survey statistics", h.logger, logCtx)
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully retrieved survey statistics by ID")

	surveyStats := models.SurveyStatsBySurvey{
		SurveyId:   uint(stats.SurveyStatsBySurvey.ServeyId),
		SurveyText: stats.SurveyStatsBySurvey.ServeyText,
		AvgRating:  float64(stats.SurveyStatsBySurvey.AvgRating),
	}
	httpresponse.SendJSONResponse(logCtx, w, surveyStats, http.StatusOK, h.logger)
}

func (h *SurveyHandler) GetSurveyStatsByUserId(w http.ResponseWriter, r *http.Request) {
	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for survey statistics by survey ID")

	_, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {
		response := httpresponse.Response{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	vars := mux.Vars(r)
	userIdStr := vars["id"]

	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		h.logger.Warn("Failed to parse user ID", slog.String("userID", userIdStr), slog.String("error", err.Error()))
		response := httpresponse.Response{
			Message: "Invalid user ID",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	stats, err := h.client.GetSurveyStatsByUserId(r.Context(), &surveysGen.GetSurveyStatsByUserIdRequest{
		UserId: uint32(userId),
	})
	if err != nil {
		logCtx := log.AppendCtx(r.Context(), slog.String("userId", userIdStr))
		response, status := ErrorCheck(err, "retrieve survey statistics by user", h.logger, logCtx)
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully retrieved survey statistics by user ID")

	userSurveyStatsList := make(models.UserSurveyStatsList, len(stats.UserServeyStats))
	for i, stat := range stats.UserServeyStats {
		userSurveyStatsList[i] = models.UserSurveyStats{
			SurveyId:   uint(stat.ServeyId),
			SurveyText: stat.ServeyText,
			Answered:   stat.Answered,
		}
	}
	httpresponse.SendJSONResponse(logCtx, w, userSurveyStatsList, http.StatusOK, h.logger)

}
