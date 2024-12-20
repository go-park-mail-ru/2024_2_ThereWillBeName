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

type UserResponseWithToken struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

type ResponseWithId struct {
	ID      uint   `json:"id"`
	Message string `json:"message"`
}

type Response struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserResponse struct {
	ID      uint32      `json:"id"`
	Profile UserProfile `json:"profile"`
}

type Achievement struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	IconPath string `json:"icon_path"`
}

//easyjson:json
type AchievementList []Achievement

func ValidateUser(v *validator.Validator, user *User) {
	// v.Check(user.Login != "", "login", "must be provided")
	v.Check(user.Password != "", "password", "must be provided")
	v.Check(user.Email != "", "email", "must be provided")
	v.Matches(user.Email, validator.EmailRX)
}
