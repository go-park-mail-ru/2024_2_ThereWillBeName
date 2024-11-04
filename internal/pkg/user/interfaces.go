package user

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"mime/multipart"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go
type UserUsecase interface {
	SignUp(ctx context.Context, user models.User) (uint, error)
	Login(ctx context.Context, login, password string) (models.User, error)
	UploadAvatar(ctx context.Context, userID uint, avatarFile multipart.File, header *multipart.FileHeader) (string, error)
	GetProfile(ctx context.Context, userID, requesterID uint) (models.UserProfile, error)
}

type UserRepo interface {
	CreateUser(ctx context.Context, user models.User) (uint, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, id string) error
	GetUsers(ctx context.Context, count, offset int64) ([]models.User, error)
	UpdateAvatarPathByUserId(ctx context.Context, userID uint, avatarPath string) error
	GetAvatarPathByUserId(ctx context.Context, userID uint) (string, error)
	GetUserByID(ctx context.Context, userID uint) (models.UserProfile, error)
}
