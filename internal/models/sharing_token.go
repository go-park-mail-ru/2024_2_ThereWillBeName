package models

import (
	"time"
)

type SharingToken struct {
	ID            uint      `json:"id"`
	TripID        uint      `json:"trip_id"`
	Token         string    `json:"token"`
	SharingOption string    `json:"sharing_option"`
	ExpiresAt     time.Time `json:"expires_at"`
	CreatedAt     time.Time `json:"created_at"`
}
