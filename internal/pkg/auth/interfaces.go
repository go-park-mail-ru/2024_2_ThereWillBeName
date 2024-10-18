package auth

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go
type AuthUsecase interface {
	SignUp(ctx context.Context, user models.User) error
	Login(ctx context.Context, login, password string) (models.User, error)
}

type AuthRepo interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, id string) error
	GetUsers(ctx context.Context, count, offset int64) ([]models.User, error)
}
