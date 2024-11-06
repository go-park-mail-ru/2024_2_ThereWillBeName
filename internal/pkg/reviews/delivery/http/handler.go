package http

import (
	"2024_2_ThereWillBeName/internal/models"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"2024_2_ThereWillBeName/internal/pkg/reviews"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
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

func ErrorCheck(err error, action string) (httpresponse.ErrorResponse, int) {
	if errors.Is(err, models.ErrNotFound) {
		log.Printf("%s error: %s", action, err)
		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		return response, http.StatusNotFound
	}
	log.Printf("%s error: %s", action, err)
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
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self';")

	var review models.Review
	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		log.Printf("create error: %s", err)
		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	review.ReviewText = template.HTMLEscapeString(review.ReviewText)

	err = h.uc.CreateReview(context.Background(), review)
	if err != nil {
		response, status := ErrorCheck(err, "create")
		httpresponse.SendJSONResponse(w, response, status)
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
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self';")

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
		log.Printf("update error: %s", err)
		response := httpresponse.ErrorResponse{
			Message: "Invalid review data",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	review.ReviewText = template.HTMLEscapeString(review.ReviewText)

	review.ID = uint(reviewID)
	err = h.uc.UpdateReview(context.Background(), review)
	if err != nil {
		response, status := ErrorCheck(err, "update")
		httpresponse.SendJSONResponse(w, response, status)
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
		log.Printf("delete error: %s", err)
		response := httpresponse.ErrorResponse{
			Message: "Invalid review ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	err = h.uc.DeleteReview(context.Background(), uint(id))
	if err != nil {
		response, status := ErrorCheck(err, "delete")
		httpresponse.SendJSONResponse(w, response, status)
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
		log.Printf("retrieve error: %s", err)
		response := httpresponse.ErrorResponse{
			Message: "Invalid place ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}
	pageStr := r.URL.Query().Get("page")
	page := 1
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			log.Printf("retrieve error: %s", err)
			response := httpresponse.ErrorResponse{
				Message: "Invalid page number",
			}
			httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
			return
		}
	}
	limit := 10
	offset := limit * (page - 1)
	reviews, err := h.uc.GetReviewsByPlaceID(context.Background(), uint(placeID), limit, offset)
	if err != nil {
		response, status := ErrorCheck(err, "retrieve")
		httpresponse.SendJSONResponse(w, response, status)
		return
	}

	httpresponse.SendJSONResponse(w, reviews, http.StatusOK)
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
	reviewID, err := strconv.ParseUint(reviewIDStr, 10, 64)
	if err != nil {
		log.Printf("retrieve error: %s", err)
		response := httpresponse.ErrorResponse{
			Message: "Invalid review ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	review, err := h.uc.GetReview(context.Background(), uint(reviewID))
	if err != nil {
		response, status := ErrorCheck(err, "retrieve")
		httpresponse.SendJSONResponse(w, response, status)
		return
	}

	httpresponse.SendJSONResponse(w, review, http.StatusOK)
}
