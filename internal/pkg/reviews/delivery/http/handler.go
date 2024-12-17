package http

import (
	"2024_2_ThereWillBeName/internal/models"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	log "2024_2_ThereWillBeName/internal/pkg/logger"
	"2024_2_ThereWillBeName/internal/pkg/middleware"
	"2024_2_ThereWillBeName/internal/pkg/reviews/delivery/grpc/gen"
	"2024_2_ThereWillBeName/internal/validator"
	"context"
	"errors"
	"fmt"
	"github.com/mailru/easyjson"
	"html/template"
	"log/slog"
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

func ErrorCheck(err error, action string, logger *slog.Logger, ctx context.Context) (httpresponse.Response, int) {
	logContext := log.AppendCtx(ctx, slog.String("action", action))
	logContext = log.AppendCtx(logContext, slog.Any("error", err.Error()))
	if errors.Is(err, models.ErrNotFound) {
		logger.WarnContext(logContext, fmt.Sprintf("Error during %s operation", action))
		response := httpresponse.Response{
			Message: "Invalid request",
		}
		return response, http.StatusNotFound
	}
	logger.ErrorContext(logContext, fmt.Sprintf("Failed to %s reviews", action))
	response := httpresponse.Response{
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
// @Failure 400 {object} httpresponses.Response "Invalid request"
// @Failure 403 {object} httpresponses.Response "Token is missing"
// @Failure 403 {object} httpresponses.Response "Invalid token"
// @Failure 500 {object} httpresponses.Response "Failed to create review"
// @Router /reviews [post]
func (h *ReviewHandler) CreateReviewHandler(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()
	h.logger.DebugContext(logCtx, "Handling request for creating a review")

	userID, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.WarnContext(logCtx, "Failed to retrieve user ID from context")

		response := httpresponse.Response{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	var review models.Review
	err := easyjson.UnmarshalFromReader(r.Body, &review)
	if err != nil {
		h.logger.WarnContext(logCtx, "Failed to decode review data",
			slog.Any("error", err.Error()),
			slog.String("review_data", fmt.Sprintf("%+v", review)))

		response := httpresponse.Response{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}
	v := validator.New()
	if models.ValidateReview(v, &review); !v.Valid() {
		h.logger.WarnContext(logCtx, "Review data is not valid")
		httpresponse.SendJSONResponse(logCtx, w, nil, http.StatusUnprocessableEntity, h.logger)
		return
	}

	review.ReviewText = template.HTMLEscapeString(review.ReviewText)

	if review.Rating < 1 || review.Rating > 5 {
		h.logger.WarnContext(logCtx, "Invalid rating",
			slog.String("rating", strconv.Itoa(review.Rating)))

		response := httpresponse.Response{
			Message: "Invalid rating",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
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

	h.logger.DebugContext(logCtx, "Review request details", slog.Any("reviewRequest", reviewRequest))

	createdReview, err := h.client.CreateReview(logCtx, &gen.CreateReviewRequest{Review: reviewRequest})
	if err != nil {
		response, status := ErrorCheck(err, "create", h.logger, logCtx)
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully created a review",
		slog.Int("review_id", int(createdReview.Review.Id)))

	reviewResponce := models.GetReview{
		ID:         uint(createdReview.Review.Id),
		UserLogin:  createdReview.Review.UserLogin,
		AvatarPath: createdReview.Review.AvatarPath,
		Rating:     int(createdReview.Review.Rating),
		ReviewText: createdReview.Review.ReviewText,
	}
	httpresponse.SendJSONResponse(logCtx, w, reviewResponce, http.StatusCreated, h.logger)
}

// UpdateReviewHandler godoc
// @Summary Update an existing review
// @Description Update review details by review ID
// @Accept json
// @Produce json
// @Param id path int true "Review ID"
// @Param review body models.Review true "Updated review details"
// @Success 200 {object} models.Review "Review updated successfully"
// @Failure 400 {object} httpresponses.Response "Invalid review ID"
// @Failure 403 {object} httpresponses.Response "Token is missing"
// @Failure 403 {object} httpresponses.Response "Invalid token"
// @Failure 404 {object} httpresponses.Response "Review not found"
// @Failure 500 {object} httpresponses.Response "Failed to update review"
// @Router /reviews/{id} [put]
func (h *ReviewHandler) UpdateReviewHandler(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()

	_, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.WarnContext(logCtx, "Failed to retrieve user ID from context")

		response := httpresponse.Response{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	var review models.Review
	vars := mux.Vars(r)
	reviewID, err := strconv.Atoi(vars["reviewID"])

	logCtx = log.AppendCtx(logCtx, slog.Int("review_id", reviewID))
	h.logger.DebugContext(logCtx, "Handling request for updating a review")

	if err != nil || reviewID < 0 {
		response := httpresponse.Response{
			Message: "Invalid review ID",
		}
		h.logger.WarnContext(logCtx, "Failed to parse place ID", slog.Any("error", err.Error()))

		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}
	err = easyjson.UnmarshalFromReader(r.Body, &review)
	if err != nil {
		h.logger.WarnContext(logCtx, "Failed to decode review data", slog.String("review_data", fmt.Sprintf("%+v", review)), slog.String("error", err.Error()))
		response := httpresponse.Response{
			Message: "Invalid review data",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	v := validator.New()
	if models.ValidateReview(v, &review); !v.Valid() {
		h.logger.WarnContext(logCtx, "Review data is not valid")
		httpresponse.SendJSONResponse(logCtx, w, nil, http.StatusUnprocessableEntity, h.logger)
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

	h.logger.DebugContext(logCtx, "Review request details", slog.Any("reviewRequest", reviewRequest))

	_, err = h.client.UpdateReview(r.Context(), &gen.UpdateReviewRequest{Review: reviewRequest})
	if err != nil {
		response, status := ErrorCheck(err, "update", h.logger, logCtx)
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully updated a review")

	httpresponse.SendJSONResponse(logCtx, w, httpresponse.Response{"Successfully updated a review"}, http.StatusOK, h.logger)
}

// DeleteReviewHandler godoc
// @Summary Delete a review
// @Description Delete a review by review ID
// @Produce json
// @Param id path int true "Review ID"
// @Success 204 "Review deleted successfully"
// @Failure 400 {object} httpresponses.Response "Invalid review ID"
// @Failure 403 {object} httpresponses.Response "Token is missing"
// @Failure 403 {object} httpresponses.Response "Invalid token"
// @Failure 404 {object} httpresponses.Response "Review not found"
// @Failure 500 {object} httpresponses.Response "Failed to delete review"
// @Router /reviews/{id} [delete]
func (h *ReviewHandler) DeleteReviewHandler(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()

	vars := mux.Vars(r)
	idStr := vars["reviewID"]

	_, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.WarnContext(logCtx, "Failed to retrieve user ID from context")

		response := httpresponse.Response{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	logCtx = log.AppendCtx(logCtx, slog.String("reviewID", idStr))
	h.logger.DebugContext(logCtx, "Handling request for deleting a review")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.WarnContext(logCtx, "Failed to parse review ID", slog.Any("error", err.Error()))
		response := httpresponse.Response{
			Message: "Invalid review ID",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	_, err = h.client.DeleteReview(r.Context(), &gen.DeleteReviewRequest{Id: uint32(id)})
	if err != nil {
		response, status := ErrorCheck(err, "delete", h.logger, logCtx)
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully deleted a review")

	httpresponse.SendJSONResponse(logCtx, w, httpresponse.Response{"Review deleted successfully"}, http.StatusOK, h.logger)
}

// GetReviewsByPlaceIDHandler godoc
// @Summary Retrieve reviews by place ID
// @Description Get all reviews for a specific place
// @Produce json
// @Param placeID path int true "Place ID"
// @Success 200 {array} models.Review "List of reviews"
// @Failure 400 {object} httpresponses.Response "Invalid place ID"
// @Failure 404 {object} httpresponses.Response "No reviews found for the place"
// @Failure 500 {object} httpresponses.Response "Failed to retrieve reviews"
// @Router /attractions/{placeID}/reviews [get]
func (h *ReviewHandler) GetReviewsByPlaceIDHandler(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()

	vars := mux.Vars(r)
	placeIDStr := vars["placeID"]

	logCtx = log.AppendCtx(logCtx, slog.String("place_id", placeIDStr))

	h.logger.DebugContext(logCtx, "Handling request for getting reviews by place ID")

	placeID, err := strconv.ParseUint(placeIDStr, 10, 64)
	if err != nil {
		h.logger.WarnContext(logCtx, "Failed to parse place ID", slog.Any("error", err.Error()))
		response := httpresponse.Response{
			Message: "Invalid place ID",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}
	pageStr := r.URL.Query().Get("page")
	page := 1
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			h.logger.WarnContext(logCtx, "Invalid page number", slog.Any("error", err.Error()))
			response := httpresponse.Response{
				Message: "Invalid page number",
			}
			httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
			return
		}
	}
	limit := 10
	offset := limit * (page - 1)
	reviews, err := h.client.GetReviewsByPlaceID(r.Context(), &gen.GetReviewsByPlaceIDRequest{PlaceId: uint32(placeID), Limit: int32(limit), Offset: int32(offset)})
	if err != nil {
		response, status := ErrorCheck(err, "retrieve", h.logger, logCtx)
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully got reviews by place ID", slog.Int("reviews_count", len(reviews.Reviews)))

	reviewsResponse := make(models.GetReviewList, len(reviews.Reviews))
	for i, review := range reviews.Reviews {
		reviewsResponse[i] = models.GetReview{
			ID:         uint(review.Id),
			UserLogin:  review.UserLogin,
			AvatarPath: review.AvatarPath,
			Rating:     int(review.Rating),
			ReviewText: review.ReviewText,
		}
	}

	httpresponse.SendJSONResponse(logCtx, w, reviewsResponse, http.StatusOK, h.logger)
}

// GetReviewsByUserIDHandler godoc
// @Summary Retrieve reviews by user ID
// @Description Get all reviews for an user
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {array} models.GetReviewByUserID "List of reviews"
// @Failure 400 {object} httpresponses.Response "Invalid user ID"
// @Failure 404 {object} httpresponses.Response "No reviews found for the user"
// @Failure 500 {object} httpresponses.Response "Failed to retrieve reviews"
// @Router /users/{userID}/reviews [get]
func (h *ReviewHandler) GetReviewsByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()

	userID, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.WarnContext(logCtx, "Failed to retrieve user ID from context")

		response := httpresponse.Response{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	logCtx = log.AppendCtx(logCtx, slog.Int("user_id", int(userID)))
	h.logger.DebugContext(logCtx, "Handling request for getting reviews by user ID")

	pageStr := r.URL.Query().Get("page")
	page := 1
	var err error
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			h.logger.WarnContext(logCtx, "Invalid page number", slog.Any("error", err.Error()))
			response := httpresponse.Response{
				Message: "Invalid page number",
			}
			httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
			return
		}
	}
	limit := 10
	offset := limit * (page - 1)
	reviews, err := h.client.GetReviewsByUserID(r.Context(), &gen.GetReviewsByUserIDRequest{UserId: uint32(userID), Limit: int32(limit), Offset: int32(offset)})
	if err != nil {
		response, status := ErrorCheck(err, "retrieve", h.logger, logCtx)
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully got reviews by user ID", slog.Int("reviews_count", len(reviews.Reviews)))

	reviewsResponse := make(models.GetReviewByUserIDList, len(reviews.Reviews))
	for i, review := range reviews.Reviews {
		reviewsResponse[i] = models.GetReviewByUserID{
			ID:         uint(review.Id),
			PlaceName:  review.PlaceName,
			Rating:     int(review.Rating),
			ReviewText: review.ReviewText,
		}
	}
	httpresponse.SendJSONResponse(logCtx, w, reviewsResponse, http.StatusOK, h.logger)
}

// GetReviewHandler godoc
// @Summary Retrieve a review by ID
// @Description Get review details by review ID
// @Produce json
// @Param id path int true "Review ID"
// @Success 200 {object} models.GetReview "Review details"
// @Failure 400 {object} httpresponses.Response "Invalid review ID"
// @Failure 404 {object} httpresponses.Response "Review not found"
// @Failure 500 {object} httpresponses.Response "Failed to retrieve review"
// @Router /reviews/{id} [get]
func (h *ReviewHandler) GetReviewHandler(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()

	vars := mux.Vars(r)
	reviewIDStr := vars["reviewID"]

	logCtx = log.AppendCtx(logCtx, slog.String("reviewID", reviewIDStr))
	h.logger.DebugContext(logCtx, "Handling request for getting review by ID")

	reviewID, err := strconv.ParseUint(reviewIDStr, 10, 64)
	if err != nil {
		h.logger.WarnContext(logCtx, "Failed to parse review ID", slog.Any("error", err.Error()))
		response := httpresponse.Response{
			Message: "Invalid review ID",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	review, err := h.client.GetReview(r.Context(), &gen.GetReviewRequest{Id: uint32(reviewID)})
	if err != nil {
		response, status := ErrorCheck(err, "retrieve", h.logger, logCtx)
		httpresponse.SendJSONResponse(logCtx, w, response, status, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Successfully got review by ID")

	reviewResponce := models.GetReview{
		ID:         uint(review.Review.Id),
		UserLogin:  review.Review.UserLogin,
		AvatarPath: review.Review.AvatarPath,
		Rating:     int(review.Review.Rating),
		ReviewText: review.Review.ReviewText,
	}

	httpresponse.SendJSONResponse(logCtx, w, reviewResponce, http.StatusOK, h.logger)
}
