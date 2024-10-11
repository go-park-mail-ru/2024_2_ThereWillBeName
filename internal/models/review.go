package models

import "time"

type Review struct {
	ID         uint      `json:"id" db:"id"`
	UserID     uint      `json:"user_id" db:"user_id"`
	PlaceID    uint      `json:"place_id" db:"place_id"`
	Rating     int       `json:"rating" db:"rating"`
	ReviewText string    `json:"review_text" db:"review_text"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}
