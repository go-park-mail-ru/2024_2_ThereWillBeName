package models

import "time"

type Trip struct {
	ID          uint      `json:"id" db:"id"`
	UserID      uint      `json:"user_id" db:"user_id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Destination string    `json:"destination" db:"destination"`
	StartDate   string    `json:"start_date" db:"start_date"`
	EndDate     string    `json:"end_date" db:"end_date"`
	Private     bool      `json:"private" db:"private"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
