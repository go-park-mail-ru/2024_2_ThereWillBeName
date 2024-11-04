package usecase

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/jwt"
	"2024_2_ThereWillBeName/internal/pkg/user"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseImpl struct {
	repo user.UserRepo
}

func NewUserUsecase(repo user.UserRepo, jwt *jwt.JWT) *UserUsecaseImpl {
	return &UserUsecaseImpl{
		repo: repo,
	}
}

func saveAvatarFile(avatarFile multipart.File, path string) error {
	outFile, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create avatar file: %w", models.ErrInternal)
	}
	defer outFile.Close()

	if _, err := io.Copy(outFile, avatarFile); err != nil {
		return fmt.Errorf("failed to write avatar content: %w", models.ErrInternal)
	}

	return nil
}

func (a *UserUsecaseImpl) SignUp(ctx context.Context, user models.User) (uint, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	return a.repo.CreateUser(ctx, user)
}

func (a *UserUsecaseImpl) Login(ctx context.Context, email, password string) (models.User, error) {
	user, err := a.repo.GetUserByEmail(ctx, email)

	if err != nil {
		log.Printf("Error retrieving user: %v\n", err)
		return models.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Printf("Password mismatch: %v\n", err)
		return models.User{}, err
	} else {
		log.Println("Password match!")
	}

	return user, nil
}

func (a *UserUsecaseImpl) UploadAvatar(ctx context.Context, userID uint, avatarFile multipart.File, header *multipart.FileHeader) (string, error) {
	avatarPath, err := a.repo.GetAvatarPathByUserId(ctx, userID)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return "", fmt.Errorf("failed to fetch avatar path: %w", models.ErrNotFound)
		}
		return "", fmt.Errorf("internal error: %w", models.ErrInternal)
	}

	storagePath := os.Getenv("AVATAR_STORAGE_PATH")
	fileExt := filepath.Ext(header.Filename)

	var realAvatarPath string
	if avatarPath != "" {
		//если путь к аватару есть в таблице у юзера -> изменяем только расширение у файла
		realAvatarPath = avatarPath[:len(avatarPath)-len(filepath.Ext(avatarPath))] + fileExt
	} else {
		realAvatarPath = filepath.Join(storagePath, fmt.Sprintf("user_%d_avatar%s", userID, fileExt))
	}

	if err := saveAvatarFile(avatarFile, realAvatarPath); err != nil {
		return err.Error(), err
	}

	if avatarPath == "" || realAvatarPath != avatarPath {
		if err := a.repo.UpdateAvatarPathByUserId(ctx, userID, realAvatarPath); err != nil {
			return "", fmt.Errorf("failed to update avatar path in database: %w", models.ErrInternal)
		}
	}

	return realAvatarPath, nil
}

func (a *UserUsecaseImpl) GetProfile(ctx context.Context, userID, requesterID uint) (models.UserProfile, error) {
	user, err := a.repo.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return models.UserProfile{}, fmt.Errorf("invalid request: %w", models.ErrNotFound)
		}
		return models.UserProfile{}, fmt.Errorf("internal error: %w", models.ErrInternal)
	}

	if requesterID != userID {
		user.Email = ""
	}

	return models.UserProfile{
		Login:      user.Login,
		AvatarPath: user.AvatarPath,
		Email:      user.Email,
	}, nil
}