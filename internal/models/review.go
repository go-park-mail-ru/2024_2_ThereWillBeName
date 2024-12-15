package models

import (
	"2024_2_ThereWillBeName/internal/validator"
	"time"
)

type Review struct {
	ID         uint      `json:"id"`
	UserID     uint      `json:"user_id"`
	PlaceID    uint      `json:"place_id"`
	Rating     int       `json:"rating"`
	ReviewText string    `json:"review_text"`
	CreatedAt  time.Time `json:"created_at"`
}

type GetReview struct {
	ID         uint   `json:"id"`
	UserLogin  string `json:"user_login"`
	AvatarPath string `json:"avatar_path"`
	Rating     int    `json:"rating"`
	ReviewText string `json:"review_text"`
}

type GetReviewByUserID struct {
	ID         uint   `json:"id"`
	PlaceName  string `json:"place_name"`
	Rating     int    `json:"rating"`
	ReviewText string `json:"review_text"`
}

//easyjson:json
type ReviewList []Review

//easyjson:json
type GetReviewList []GetReview

//easyjson:json
type GetReviewByUserIDList []GetReviewByUserID

func ValidateReview(v *validator.Validator, review *Review) {
	v.Check(review.ReviewText != "", "reviewText", "must be provided")
	v.Check(len(review.ReviewText) <= 255, "reviewText", "must not be more than 255 symbols")
	v.Check(review.Rating != 0, "rating", "must be provided")
	v.Check(review.PlaceID != 0, "place id", "must be provided")
	v.Check(review.UserID != 0, "user id", "must be provided")
}
