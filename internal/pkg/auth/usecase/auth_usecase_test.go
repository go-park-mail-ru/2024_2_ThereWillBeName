package usecase

import (
	"2024_2_ThereWillBeName/internal/models"
	mocks "2024_2_ThereWillBeName/internal/pkg/auth/mocks"
	"2024_2_ThereWillBeName/internal/pkg/jwt"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockAuthRepo(ctrl)
	mockJWT := jwt.NewJWT("secret_key")

	authUsecase := NewAuthUsecase(mockRepo, mockJWT)

	user := models.User{
		ID:    1,
		Login: "testuser",
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	assert.NoError(t, err)
	user.Password = string(hashedPassword)
	mockRepo.EXPECT().GetUserByLogin(gomock.Any(), user.Login).Return(user, nil).Times(1)

	token, err := authUsecase.Login(context.Background(), user.Login, "password123")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAuthRepo(ctrl)
	mockJWT := jwt.NewJWT("secret_key")

	authUsecase := NewAuthUsecase(mockRepo, mockJWT)

	user := models.User{
		Login:    "testuser",
		Password: "password123",
	}

	mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Do(func(ctx context.Context, u models.User) {
		assert.NotEqual(t, u.Password, "password123")
		assert.Equal(t, u.Login, user.Login)
	}).Return(nil).Times(1)

	err := authUsecase.SignUp(context.Background(), user)
	assert.NoError(t, err)
}
