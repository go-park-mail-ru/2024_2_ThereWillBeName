package usecase

import (
    "context"
    "2024_2_ThereWillBeName/internal/models"
    "2024_2_ThereWillBeName/internal/pkg/auth"
    "2024_2_ThereWillBeName/internal/pkg/jwt" 
    "golang.org/x/crypto/bcrypt"
)

type AuthUsecaseImpl struct {
    repo auth.AuthRepo
    jwt  *jwt.JWT
}

func NewAuthUsecase(repo auth.AuthRepo, jwt *jwt.JWT) *AuthUsecaseImpl {
    return &AuthUsecaseImpl{
        repo: repo,
        jwt: jwt,
    }
}

func (a *AuthUsecaseImpl) SignUp(ctx context.Context, user models.User) error {
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    user.Password = string(hashedPassword)
    return a.repo.CreateUser(ctx, user)
}

func (a *AuthUsecaseImpl) Login(ctx context.Context, email, password string) (string, error) {
    user, err := a.repo.GetUserByEmail(ctx, email)
    if err != nil {
        return "", err
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", err
    }

    return a.jwt.GenerateToken(uint(user.ID), user.Email)
}

func (a *AuthUsecaseImpl) Logout(ctx context.Context, token string) error {

    return nil
}
