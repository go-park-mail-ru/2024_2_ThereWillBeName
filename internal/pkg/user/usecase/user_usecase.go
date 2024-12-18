package usecase

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/user"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseImpl struct {
	repo        user.UserRepo
	storagePath string
}

func NewUserUsecase(repo user.UserRepo, storagePath string) *UserUsecaseImpl {
	return &UserUsecaseImpl{
		repo:        repo,
		storagePath: storagePath,
	}
}

func saveAvatarData(avatarData []byte, path string) error {
	outFile, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create avatar file: %w", models.ErrInternal)
	}
	defer outFile.Close()

	if _, err := outFile.Write(avatarData); err != nil {
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

func (a *UserUsecaseImpl) UploadAvatar(ctx context.Context, userID uint, avatarData []byte, avatarFileName string) (string, error) {
	avatarPath, err := a.repo.GetAvatarPathByUserId(ctx, userID)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return "", fmt.Errorf("failed to fetch avatar path: %w", models.ErrNotFound)
		}
		return "", fmt.Errorf("internal error: %w", models.ErrInternal)
	}

	realAvatarPath := filepath.Join(a.storagePath, avatarFileName)

	if avatarPath != "" && avatarPath != avatarFileName {
		oldAvatarPath := filepath.Join(a.storagePath, avatarPath)
		if err := os.Remove(oldAvatarPath); err != nil {
			return "", fmt.Errorf("failed to delete old avatar file: %w", err)
		}
	}

	if err := saveAvatarData(avatarData, realAvatarPath); err != nil {
		return "", fmt.Errorf("failed to save avatar file: %w", err)
	}

	if avatarPath == "" || avatarPath != avatarFileName {
		if err := a.repo.UpdateAvatarPathByUserId(ctx, userID, avatarFileName); err != nil {
			return "", fmt.Errorf("failed to update avatar path in database: %w", models.ErrInternal)
		}
	}

	return avatarFileName, nil
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

func (a *UserUsecaseImpl) UpdatePassword(ctx context.Context, userData models.User, newPassword string) error {
	user, err := a.repo.GetUserByEmail(ctx, userData.Email)
	if err != nil {
		log.Printf("UpdatePassword: Error retrieving user: %v\n", err)
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password)); err != nil {
		log.Printf("UpdatePassword: Password mismatch: %v\n", err)
		return models.ErrMismatch
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)

	return a.repo.UpdatePassword(ctx, user.ID, string(hashedPassword))
}

func (a *UserUsecaseImpl) UpdateProfile(ctx context.Context, userID uint, login, email string) error {
	return a.repo.UpdateProfile(ctx, userID, login, email)
}

func (a *UserUsecaseImpl) GetAchievements(ctx context.Context, userID uint) ([]models.Achievement, error) {
	achievements, err := a.repo.GetAchievements(ctx, userID)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return nil, fmt.Errorf("invalid request: %w", models.ErrNotFound)
		}
		return nil, fmt.Errorf("internal error: %w", models.ErrInternal)
	}
	return achievements, nil
}
