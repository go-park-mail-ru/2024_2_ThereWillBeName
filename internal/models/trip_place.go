package models

import "time"

type TripPlace struct {
	ID        uint      `json:"id"`
	TripID    uint      `json:"trip_id"`
	PlaceID   int       `json:"place_id"`
	CreatedAt time.Time `json:"created_at"`
}
