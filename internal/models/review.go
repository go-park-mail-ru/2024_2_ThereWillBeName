package models

import "time"

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
