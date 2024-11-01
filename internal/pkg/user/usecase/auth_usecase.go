package usecase

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/jwt"
	"2024_2_ThereWillBeName/internal/pkg/user"
	"context"
	"log"

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

func (a *UserUsecaseImpl) SignUp(ctx context.Context, user models.User) error {
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
