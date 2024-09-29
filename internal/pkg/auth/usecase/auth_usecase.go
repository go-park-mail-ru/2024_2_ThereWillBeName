package usecase

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/auth"
	"2024_2_ThereWillBeName/internal/pkg/jwt"
	"context"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecaseImpl struct {
	repo auth.AuthRepo
	jwt  *jwt.JWT
}

func NewAuthUsecase(repo auth.AuthRepo, jwt *jwt.JWT) *AuthUsecaseImpl {
	return &AuthUsecaseImpl{
		repo: repo,
		jwt:  jwt,
	}
}

func (a *AuthUsecaseImpl) SignUp(ctx context.Context, user models.User) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	return a.repo.CreateUser(ctx, user)
}

func (a *AuthUsecaseImpl) Login(ctx context.Context, login, password string) (string, error) {
	user, err := a.repo.GetUserByLogin(ctx, login)

	if err != nil {
		log.Printf("Error retrieving user: %v\n", err)
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Printf("Password mismatch: %v\n", err)
		return "", err
	} else {
		log.Println("Password match!")
	}

	return a.jwt.GenerateToken(uint(user.ID), user.Login)
}

func (a *AuthUsecaseImpl) Logout(ctx context.Context, token string) error {

	return nil
}
