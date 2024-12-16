package models

import (
	"2024_2_ThereWillBeName/internal/validator"
	"time"
)

type User struct {
	ID         uint      `json:"id"`
	Login      string    `json:"login"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	AvatarPath string    `json:"avatar_path"`
	CreatedAt  time.Time `json:"created_at"`
}

type UserProfile struct {
	Login      string `json:"login"`
	AvatarPath string `json:"avatar_path"`
	Email      string `json:"email"`
}

func ValidateUser(v *validator.Validator, user *User) {
	// v.Check(user.Login != "", "login", "must be provided")
	v.Check(user.Password != "", "password", "must be provided")
	v.Check(user.Email != "", "email", "must be provided")
	v.Matches(user.Email, validator.EmailRX)
}
