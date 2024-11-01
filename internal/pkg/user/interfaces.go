package user

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"mime/multipart"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go
type UserUsecase interface {
	SignUp(ctx context.Context, user models.User) error
	Login(ctx context.Context, login, password string) (models.User, error)
	UpdateAvatar(ctx context.Context, userID string, avatarFile multipart.File, fileHeader *multipart.FileHeader) error
	GetProfile(ctx context.Context, userID, requesterID string) (models.User, error)
}

type UserRepo interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, id string) error
	GetUsers(ctx context.Context, count, offset int64) ([]models.User, error)
	UpdateAvatarPath(ctx context.Context, userID string, avatarPath string) error
	GetUserByID(ctx context.Context, userID string) (models.User, error)
}
