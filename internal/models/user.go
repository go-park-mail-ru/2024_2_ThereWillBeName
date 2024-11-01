package models

import "time"

type User struct {
	ID         uint      `json:"id"`
	Login      string    `json:"login"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	AvatarPath string    `json:"avatar_path,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

type UserProfile struct {
	Login      string `json:"login"`
	AvatarPath string `json:"avatar_path,omitempty"`
	Email      string `json:"email,omitempty"`
}
