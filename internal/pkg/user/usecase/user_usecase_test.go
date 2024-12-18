package usecase

// import (
// 	"2024_2_ThereWillBeName/internal/models"
// 	mock "2024_2_ThereWillBeName/internal/pkg/user/mocks"
// 	"context"
// 	"errors"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"
// 	"golang.org/x/crypto/bcrypt"
// )

// func TestUserUsecaseImpl_SignUp(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	storagePath := "/storage"
// 	mockRepo := mock.NewMockUserRepo(ctrl)
// 	usecase := NewUserUsecase(mockRepo, storagePath)

// 	user := models.User{
// 		Login:    "newuser",
// 		Password: "plainpassword",
// 	}

// 	expectedUserID := uint(1)

// 	tests := []struct {
// 		name           string
// 		mockBehavior   func()
// 		user           models.User
// 		expectedUserID uint
// 		expectedErr    string
// 	}{
// 		{
// 			name: "Success",
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.AssignableToTypeOf(models.User{})).DoAndReturn(func(ctx context.Context, user models.User) (uint, error) {
// 					assert.NotEqual(t, user.Password, "plainpassword")
// 					err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("plainpassword"))
// 					assert.NoError(t, err)
// 					return expectedUserID, nil
// 				})
// 			},
// 			user:           user,
// 			expectedUserID: expectedUserID,
// 			expectedErr:    "",
// 		},
// 		{
// 			name: "Repository Error",
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.AssignableToTypeOf(models.User{})).Return(uint(0), errors.New("internal error"))
// 			},
// 			user:           user,
// 			expectedUserID: 0,
// 			expectedErr:    "internal error",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.mockBehavior()

// 			userID, err := usecase.SignUp(context.Background(), tt.user)

// 			if tt.expectedErr != "" {
// 				assert.Error(t, err)
// 				assert.Equal(t, tt.expectedErr, err.Error())
// 				assert.Equal(t, tt.expectedUserID, userID)
// 			} else {
// 				assert.NoError(t, err)
// 				assert.Equal(t, tt.expectedUserID, userID)
// 			}
// 		})
// 	}
// }

// func TestUserUsecaseImpl_Login(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	storagePath := "/storage"
// 	mockRepo := mock.NewMockUserRepo(ctrl)
// 	usecase := NewUserUsecase(mockRepo, storagePath)

// 	userEmail := "testuser@example.com"
// 	plainPassword := "password123"
// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)

// 	user := models.User{
// 		Email:    userEmail,
// 		Password: string(hashedPassword),
// 	}

// 	tests := []struct {
// 		name         string
// 		mockBehavior func()
// 		email        string
// 		password     string
// 		expectedErr  string
// 		expectedUser models.User
// 	}{
// 		{
// 			name: "Success - Correct password",
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().GetUserByEmail(gomock.Any(), userEmail).Return(user, nil)
// 			},
// 			email:        userEmail,
// 			password:     plainPassword,
// 			expectedErr:  "",
// 			expectedUser: user,
// 		},
// 		{
// 			name: "Error - User not found",
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().GetUserByEmail(gomock.Any(), userEmail).Return(models.User{}, models.ErrNotFound)
// 			},
// 			email:        userEmail,
// 			password:     plainPassword,
// 			expectedErr:  "not found",
// 			expectedUser: models.User{},
// 		},
// 		{
// 			name: "Error - Incorrect password",
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().GetUserByEmail(gomock.Any(), userEmail).Return(user, nil)
// 			},
// 			email:        userEmail,
// 			password:     "wrongpassword",
// 			expectedErr:  "crypto/bcrypt: hashedPassword is not the hash of the given password",
// 			expectedUser: models.User{},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.mockBehavior()

// 			user, err := usecase.Login(context.Background(), tt.email, tt.password)

// 			if tt.expectedErr != "" {
// 				assert.Error(t, err)
// 				assert.Contains(t, err.Error(), tt.expectedErr)
// 			} else {
// 				assert.NoError(t, err)
// 				assert.Equal(t, tt.expectedUser, user)
// 			}
// 		})
// 	}
// }
