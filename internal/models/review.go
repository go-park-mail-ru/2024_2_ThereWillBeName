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
	Username   string `json:"username"`
	AvatarPath string `json:"avatar_path"`
	Rating     int    `json:"rating"`
	ReviewText string `json:"review_text"`
}
