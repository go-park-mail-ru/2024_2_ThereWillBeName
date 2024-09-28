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
	log.Println(user.Password)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	log.Printf(string(hashedPassword))
	log.Println(len(hashedPassword))
	user.Password = string(hashedPassword)
	return a.repo.CreateUser(ctx, user)
}

func (a *AuthUsecaseImpl) Login(ctx context.Context, email, password string) (string, error) {
	log.Println("Login function triggered")
	user, err := a.repo.GetUserByEmail(ctx, email)
	log.Printf("Hashed Password from DB: %s\n", user.Password)
	log.Printf("Password from Request: %s\n", password)

	if err != nil {
		log.Printf("Error retrieving user: %v\n", err)
		return "", err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	log.Print("Newly generated hash password:", string(hashedPassword))
	log.Println(string(hashedPassword) == (user.Password))

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Printf("Password mismatch: %v\n", err)
		return "", err
	} else {
		log.Println("Password match!")
	}

	return a.jwt.GenerateToken(uint(user.ID), user.Email)
}

func (a *AuthUsecaseImpl) Logout(ctx context.Context, token string) error {

	return nil
}
