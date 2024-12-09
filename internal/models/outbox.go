package models

import "time"

type OutboxRecord struct {
	ID          int        `json:"id"`
	EventType   string     `json:"event_type"`
	Payload     string     `json:"payload"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	ProcessedAt *time.Time `json:"processed_at,omitempty"`
}
