package models

import "time"

type Trip struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CityID      uint      `json:"city_id"`
	StartDate   string    `json:"start_date"`
	EndDate     string    `json:"end_date"`
	Private     bool      `json:"private"`
	CreatedAt   time.Time `json:"created_at"`
}
