package http

import (
	"2024_2_ThereWillBeName/internal/models"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	log "2024_2_ThereWillBeName/internal/pkg/logger"
	"2024_2_ThereWillBeName/internal/pkg/middleware"
	"2024_2_ThereWillBeName/internal/pkg/reviews/delivery/grpc/gen"
	"2024_2_ThereWillBeName/internal/validator"
	"context"
	"html/template"
	"log/slog"

	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ReviewHandler struct {
	client gen.ReviewsClient
	logger *slog.Logger
}

func NewReviewHandler(client gen.ReviewsClient, logger *slog.Logger) *ReviewHandler {
	return &ReviewHandler{client, logger}
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
// @Failure 403 {object} httpresponses.ErrorResponse "Token is missing"
// @Failure 403 {object} httpresponses.ErrorResponse "Invalid token"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to create review"
// @Router /reviews [post]
func (h *ReviewHandler) CreateReviewHandler(w http.ResponseWriter, r *http.Request) {
	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for creating review")

	userID, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.Warn("Failed to retrieve user ID from context")

		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized, h.logger)
		return
	}

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
	v := validator.New()
	if models.ValidateReview(v, &review); !v.Valid() {
		httpresponse.SendJSONResponse(w, nil, http.StatusUnprocessableEntity, h.logger)
		return
	}

	review.ReviewText = template.HTMLEscapeString(review.ReviewText)

	if review.Rating < 1 || review.Rating > 5 {
		h.logger.Warn("Invalid rating",
			slog.String("rating", strconv.Itoa(review.Rating)))

		response := httpresponse.ErrorResponse{
			Message: "Invalid rating",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}

	review.UserID = userID

	reviewRequest := &gen.Review{
		Id:         uint32(review.ID),
		UserId:     uint32(review.UserID),
		PlaceId:    uint32(review.PlaceID),
		Rating:     int32(review.Rating),
		ReviewText: review.ReviewText,
	}

	createdReview, err := h.client.CreateReview(context.Background(), &gen.CreateReviewRequest{Review: reviewRequest})
	if err != nil {
		response, status := ErrorCheck(err, "create", h.logger, context.Background())
		httpresponse.SendJSONResponse(w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully created a review")

	httpresponse.SendJSONResponse(w, createdReview.Review, http.StatusCreated, h.logger)
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
// @Failure 403 {object} httpresponses.ErrorResponse "Token is missing"
// @Failure 403 {object} httpresponses.ErrorResponse "Invalid token"
// @Failure 404 {object} httpresponses.ErrorResponse "Review not found"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to update review"
// @Router /reviews/{id} [put]
func (h *ReviewHandler) UpdateReviewHandler(w http.ResponseWriter, r *http.Request) {
	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for updating a review")

	_, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.Warn("Failed to retrieve user ID from context")

		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized, h.logger)
		return
	}

	var review models.Review
	vars := mux.Vars(r)
	reviewID, err := strconv.Atoi(vars["reviewID"])
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

	v := validator.New()
	if models.ValidateReview(v, &review); !v.Valid() {
		httpresponse.SendJSONResponse(w, nil, http.StatusUnprocessableEntity, h.logger)
		return
	}

	review.ReviewText = template.HTMLEscapeString(review.ReviewText)

	review.ID = uint(reviewID)

	reviewRequest := &gen.Review{
		ReviewText: review.ReviewText,
		UserId:     uint32(review.UserID),
		PlaceId:    uint32(review.PlaceID),
		Rating:     int32(review.Rating),
		Id:         uint32(review.ID),
	}
	res, err := h.client.UpdateReview(r.Context(), &gen.UpdateReviewRequest{Review: reviewRequest})
	if err != nil {
		logCtx := log.AppendCtx(context.Background(), slog.Int("reviewID", reviewID))
		response, status := ErrorCheck(err, "update", h.logger, logCtx)
		httpresponse.SendJSONResponse(w, response, status, h.logger)
		return
	}
	h.logger.DebugContext(logCtx, "Successfully updated a review")

	httpresponse.SendJSONResponse(w, res.Success, http.StatusOK, h.logger)
}

// DeleteReviewHandler godoc
// @Summary Delete a review
// @Description Delete a review by review ID
// @Produce json
// @Param id path int true "Review ID"
// @Success 204 "Review deleted successfully"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid review ID"
// @Failure 403 {object} httpresponses.ErrorResponse "Token is missing"
// @Failure 403 {object} httpresponses.ErrorResponse "Invalid token"
// @Failure 404 {object} httpresponses.ErrorResponse "Review not found"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to delete review"
// @Router /reviews/{id} [delete]
func (h *ReviewHandler) DeleteReviewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["reviewID"]

	_, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.Warn("Failed to retrieve user ID from context")

		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized, h.logger)
		return
	}

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

	_, err = h.client.DeleteReview(r.Context(), &gen.DeleteReviewRequest{Id: uint32(id)})
	if err != nil {
		logCtx := log.AppendCtx(context.Background(), slog.String("reviewID", idStr))
		response, status := ErrorCheck(err, "delete", h.logger, logCtx)
		httpresponse.SendJSONResponse(w, response, status, h.logger)
		return
	}
	h.logger.DebugContext(logCtx, "Successfully deleted a review")

	response := map[string]string{
		"message": "Review deleted successfully",
	}

	httpresponse.SendJSONResponse(w, response, http.StatusOK, h.logger)
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
// @Router /attractions/{placeID}/reviews [get]
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
	reviews, err := h.client.GetReviewsByPlaceID(r.Context(), &gen.GetReviewsByPlaceIDRequest{PlaceId: uint32(placeID), Limit: int32(limit), Offset: int32(offset)})
	if err != nil {
		logCtx := log.AppendCtx(context.Background(), slog.String("placeID", placeIDStr))
		response, status := ErrorCheck(err, "retrieve", h.logger, logCtx)
		httpresponse.SendJSONResponse(w, response, status, h.logger)
		return
	}
	h.logger.DebugContext(logCtx, "Successfully got reviews by place ID")

	httpresponse.SendJSONResponse(w, reviews.Reviews, http.StatusOK, h.logger)
}

// GetReviewsByUserIDHandler godoc
// @Summary Retrieve reviews by user ID
// @Description Get all reviews for an user
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {array} models.GetReviewByUserID "List of reviews"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid user ID"
// @Failure 404 {object} httpresponses.ErrorResponse "No reviews found for the user"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to retrieve reviews"
// @Router /users/{userID}/reviews [get]
func (h *ReviewHandler) GetReviewsByUserIDHandler(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.Warn("Failed to retrieve user ID from context")

		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized, h.logger)
		return
	}

	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for getting reviews by user ID", slog.Int("userID", int(userID)))

	pageStr := r.URL.Query().Get("page")
	page := 1
	var err error
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
	reviews, err := h.client.GetReviewsByUserID(r.Context(), &gen.GetReviewsByUserIDRequest{UserId: uint32(userID), Limit: int32(limit), Offset: int32(offset)})
	if err != nil {
		logCtx := log.AppendCtx(context.Background(), slog.Int("userID", int(userID)))
		response, status := ErrorCheck(err, "retrieve", h.logger, logCtx)
		httpresponse.SendJSONResponse(w, response, status, h.logger)
		return
	}
	h.logger.DebugContext(logCtx, "Successfully got reviews by user ID")

	httpresponse.SendJSONResponse(w, reviews.Reviews, http.StatusOK, h.logger)
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

	review, err := h.client.GetReview(r.Context(), &gen.GetReviewRequest{Id: uint32(reviewID)})
	if err != nil {
		logCtx := log.AppendCtx(context.Background(), slog.String("reviewID", reviewIDStr))
		response, status := ErrorCheck(err, "retrieve", h.logger, logCtx)
		httpresponse.SendJSONResponse(w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully got review by ID")

	httpresponse.SendJSONResponse(w, review.Review, http.StatusOK, h.logger)
}
