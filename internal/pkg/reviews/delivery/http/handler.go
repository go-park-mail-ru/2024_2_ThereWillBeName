package http

import (
	"2024_2_ThereWillBeName/internal/models"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	log "2024_2_ThereWillBeName/internal/pkg/logger"
	"2024_2_ThereWillBeName/internal/pkg/reviews"
	"context"
	"log/slog"

	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ReviewHandler struct {
	uc     reviews.ReviewsUsecase
	logger *slog.Logger
}

func NewReviewHandler(uc reviews.ReviewsUsecase, logger *slog.Logger) *ReviewHandler {
	return &ReviewHandler{uc, logger}
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
		Message: fmt.Sprintf("Failed to %s review", action),
	}
	return response, http.StatusInternalServerError
}

// CreateReviewHandler godoc
// @Summary Create a new review
// @Description Create a new review for a place
// @Accept json
// @Produce json
// @Param review body models.Review true "Review details"
// @Success 201 {object} models.Review "Review created successfully"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid request"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to create review"
// @Router /reviews [post]
func (h *ReviewHandler) CreateReviewHandler(w http.ResponseWriter, r *http.Request) {
	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for creating review")

	var review models.Review
	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		h.logger.Warn("Failed to decode review data",
			slog.String("error", err.Error()),
			slog.String("review_data", fmt.Sprintf("%+v", review)))

		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}

	if review.Rating < 1 || review.Rating > 5 {
		h.logger.Warn("Invalid rating",
			slog.String("rating", strconv.Itoa(review.Rating)))

		response := httpresponse.ErrorResponse{
			Message: "Invalid rating",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}

	createdReview, err := h.uc.CreateReview(context.Background(), review)
	if err != nil {
		response, status := ErrorCheck(err, "create", h.logger, context.Background())
		httpresponse.SendJSONResponse(w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully created a review")

	httpresponse.SendJSONResponse(w, createdReview, http.StatusCreated, h.logger)
}

// UpdateReviewHandler godoc
// @Summary Update an existing review
// @Description Update review details by review ID
// @Accept json
// @Produce json
// @Param id path int true "Review ID"
// @Param review body models.Review true "Updated review details"
// @Success 200 {object} models.Review "Review updated successfully"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid review ID"
// @Failure 404 {object} httpresponses.ErrorResponse "Review not found"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to update review"
// @Router /reviews/{id} [put]
func (h *ReviewHandler) UpdateReviewHandler(w http.ResponseWriter, r *http.Request) {
	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for updating a review")

	var review models.Review
	vars := mux.Vars(r)
	reviewID, err := strconv.Atoi(vars["id"])
	if err != nil || reviewID < 0 {
		response := httpresponse.ErrorResponse{
			Message: "Invalid review ID",
		}
		h.logger.Warn("Failed to parse place ID", slog.Int("reviewID", reviewID), slog.String("error", err.Error()))

		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		h.logger.Warn("Failed to decode review data", slog.String("review_data", fmt.Sprintf("%+v", review)), slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid review data",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}

	review.ID = uint(reviewID)
	err = h.uc.UpdateReview(context.Background(), review)
	if err != nil {
		logCtx := log.AppendCtx(context.Background(), slog.Int("reviewID", reviewID))
		response, status := ErrorCheck(err, "update", h.logger, logCtx)
		httpresponse.SendJSONResponse(w, response, status, h.logger)
		return
	}
	h.logger.DebugContext(logCtx, "Successfully updated a review")

	w.WriteHeader(http.StatusOK)
}

// DeleteReviewHandler godoc
// @Summary Delete a review
// @Description Delete a review by review ID
// @Produce json
// @Param id path int true "Review ID"
// @Success 204 "Review deleted successfully"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid review ID"
// @Failure 404 {object} httpresponses.ErrorResponse "Review not found"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to delete review"
// @Router /reviews/{id} [delete]
func (h *ReviewHandler) DeleteReviewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for deleting a review", slog.String("reviewID", idStr))

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.Warn("Failed to parse review ID", slog.String("reviewID", idStr), slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid review ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}

	err = h.uc.DeleteReview(context.Background(), uint(id))
	if err != nil {
		logCtx := log.AppendCtx(context.Background(), slog.String("reviewID", idStr))
		response, status := ErrorCheck(err, "delete", h.logger, logCtx)
		httpresponse.SendJSONResponse(w, response, status, h.logger)
		return
	}
	h.logger.DebugContext(logCtx, "Successfully deleted a review")

	w.WriteHeader(http.StatusNoContent)
}

// GetReviewsByPlaceIDHandler godoc
// @Summary Retrieve reviews by place ID
// @Description Get all reviews for a specific place
// @Produce json
// @Param placeID path int true "Place ID"
// @Success 200 {array} models.Review "List of reviews"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid place ID"
// @Failure 404 {object} httpresponses.ErrorResponse "No reviews found for the place"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to retrieve reviews"
// @Router /places/{placeID}/reviews [get]
func (h *ReviewHandler) GetReviewsByPlaceIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	placeIDStr := vars["placeID"]

	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for getting reviews by place ID", slog.String("placeID", placeIDStr))

	placeID, err := strconv.ParseUint(placeIDStr, 10, 64)
	if err != nil {
		h.logger.Warn("Failed to parse place ID", slog.String("placeID", placeIDStr), slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid place ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}
	pageStr := r.URL.Query().Get("page")
	page := 1
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			response := httpresponse.ErrorResponse{
				Message: "Invalid page number",
			}
			httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
			return
		}
	}
	limit := 10
	offset := limit * (page - 1)
	reviews, err := h.uc.GetReviewsByPlaceID(context.Background(), uint(placeID), limit, offset)
	if err != nil {
		logCtx := log.AppendCtx(context.Background(), slog.String("placeID", placeIDStr))
		response, status := ErrorCheck(err, "retrieve", h.logger, logCtx)
		httpresponse.SendJSONResponse(w, response, status, h.logger)
		return
	}
	h.logger.DebugContext(logCtx, "Successfully got reviews by place ID")

	httpresponse.SendJSONResponse(w, reviews, http.StatusOK, h.logger)
}

// GetReviewsByUserIDHandler godoc
// @Summary Retrieve reviews by user ID
// @Description Get all reviews for an user
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {array} models.GetReviewByUserId "List of reviews"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid user ID"
// @Failure 404 {object} httpresponses.ErrorResponse "No reviews found for the user"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to retrieve reviews"
// @Router /users/{userID}/reviews [get]
func (h *ReviewHandler) GetReviewsByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["userID"]

	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for getting reviews by user ID", slog.String("userID", userIDStr))

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		h.logger.Warn("Failed to parse user ID", slog.String("userID", userIDStr), slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid user ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}
	pageStr := r.URL.Query().Get("page")
	page := 1
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			response := httpresponse.ErrorResponse{
				Message: "Invalid page number",
			}
			httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
			return
		}
	}
	limit := 10
	offset := limit * (page - 1)
	reviews, err := h.uc.GetReviewsByUserID(context.Background(), uint(userID), limit, offset)
	if err != nil {
		logCtx := log.AppendCtx(context.Background(), slog.String("userID", userIDStr))
		response, status := ErrorCheck(err, "retrieve", h.logger, logCtx)
		httpresponse.SendJSONResponse(w, response, status, h.logger)
		return
	}
	h.logger.DebugContext(logCtx, "Successfully got reviews by user ID")

	httpresponse.SendJSONResponse(w, reviews, http.StatusOK, h.logger)
}

// GetReviewHandler godoc
// @Summary Retrieve a review by ID
// @Description Get review details by review ID
// @Produce json
// @Param id path int true "Review ID"
// @Success 200 {object} models.GetReview "Review details"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid review ID"
// @Failure 404 {object} httpresponses.ErrorResponse "Review not found"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to retrieve review"
// @Router /reviews/{id} [get]
func (h *ReviewHandler) GetReviewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reviewIDStr := vars["reviewID"]

	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for getting review by ID", slog.String("reviewID", reviewIDStr))

	reviewID, err := strconv.ParseUint(reviewIDStr, 10, 64)
	if err != nil {
		h.logger.Warn("Failed to parse review ID", slog.String("reviewID", reviewIDStr), slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid review ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}

	review, err := h.uc.GetReview(context.Background(), uint(reviewID))
	if err != nil {
		logCtx := log.AppendCtx(context.Background(), slog.String("reviewID", reviewIDStr))
		response, status := ErrorCheck(err, "retrieve", h.logger, logCtx)
		httpresponse.SendJSONResponse(w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully got review by ID")

	httpresponse.SendJSONResponse(w, review, http.StatusOK, h.logger)
}
