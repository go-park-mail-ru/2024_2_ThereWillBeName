package models

import "time"

type User struct {
	ID        uint      `json:"id" db:"id"`
	Login     string    `json:"login" db:"login"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
