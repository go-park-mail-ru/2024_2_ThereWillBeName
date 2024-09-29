package auth

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
)

type AuthUsecase interface {
	SignUp(ctx context.Context, user models.User) error
	Login(ctx context.Context, login, password string) (string, error) // Возвращает JWT токен
	Logout(ctx context.Context, token string) error
}

type AuthRepo interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByLogin(ctx context.Context, login string) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, id string) error
	GetUsers(ctx context.Context, count, offset int64) ([]models.User, error)
}
