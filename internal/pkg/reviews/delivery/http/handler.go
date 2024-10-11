package http

import (
	"2024_2_ThereWillBeName/internal/models"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"2024_2_ThereWillBeName/internal/pkg/reviews"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ReviewHandler struct {
	uc reviews.ReviewsUsecase
}

func NewReviewHandler(uc reviews.ReviewsUsecase) *ReviewHandler {
	return &ReviewHandler{uc}
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
	var review models.Review
	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	err = h.uc.CreateReview(context.Background(), review)
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Failed to create review",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
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
	var review models.Review
	vars := mux.Vars(r)
	reviewID, err := strconv.Atoi(vars["id"])
	if err != nil || reviewID < 0 {
		response := httpresponse.ErrorResponse{
			Message: "Invalid review ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Invalid review data",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	review.ID = uint(reviewID)
	err = h.uc.UpdateReview(context.Background(), review)
	if err != nil {
		if err.Error() == "review not found" {
			response := httpresponse.ErrorResponse{
				Message: "Review not found",
			}
			httpresponse.SendJSONResponse(w, response, http.StatusNotFound)
		} else {
			response := httpresponse.ErrorResponse{
				Message: "Failed to update review",
			}
			httpresponse.SendJSONResponse(w, response, http.StatusInternalServerError)
		}
		return
	}

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
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Invalid review ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	err = h.uc.DeleteReview(context.Background(), uint(id))
	if err != nil {
		if err.Error() == "review not found" {
			response := httpresponse.ErrorResponse{
				Message: "Review not found",
			}
			httpresponse.SendJSONResponse(w, response, http.StatusNotFound)
		} else {
			response := httpresponse.ErrorResponse{
				Message: "Failed to delete review",
			}
			httpresponse.SendJSONResponse(w, response, http.StatusInternalServerError)
		}
		return
	}

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
	placeID, err := strconv.ParseUint(placeIDStr, 10, 64)
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Invalid place ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	reviews, err := h.uc.GetReviewsByPlaceID(context.Background(), uint(placeID))
	if err != nil {
		if err.Error() == "no reviews found for the place" {
			response := httpresponse.ErrorResponse{
				Message: "No reviews found for the place",
			}
			httpresponse.SendJSONResponse(w, response, http.StatusNotFound)
		} else {
			response := httpresponse.ErrorResponse{
				Message: "Failed to retrieve reviews",
			}
			httpresponse.SendJSONResponse(w, response, http.StatusInternalServerError)
		}
		return
	}

	httpresponse.SendJSONResponse(w, reviews, http.StatusOK)
}

// GetReviewHandler godoc
// @Summary Retrieve a review by ID
// @Description Get review details by review ID
// @Produce json
// @Param id path int true "Review ID"
// @Success 200 {object} models.Review "Review details"
// @Failure 400 {object} httpresponses.ErrorResponse "Invalid review ID"
// @Failure 404 {object} httpresponses.ErrorResponse "Review not found"
// @Failure 500 {object} httpresponses.ErrorResponse "Failed to retrieve review"
// @Router /reviews/{id} [get]
func (h *ReviewHandler) GetReviewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reviewIDStr := vars["id"]
	reviewID, err := strconv.ParseUint(reviewIDStr, 10, 64)
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Invalid review ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	review, err := h.uc.GetReview(context.Background(), uint(reviewID))
	if err != nil {
		if err.Error() == "review not found" {
			response := httpresponse.ErrorResponse{
				Message: "Review not found",
			}
			httpresponse.SendJSONResponse(w, response, http.StatusNotFound)
		} else {
			response := httpresponse.ErrorResponse{
				Message: "Failed to retrieve review",
			}
			httpresponse.SendJSONResponse(w, response, http.StatusInternalServerError)
		}
		return
	}

	httpresponse.SendJSONResponse(w, review, http.StatusOK)
}
