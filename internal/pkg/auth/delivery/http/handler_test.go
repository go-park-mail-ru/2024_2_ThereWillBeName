package http

import (
	"2024_2_ThereWillBeName/internal/models"
	mock "2024_2_ThereWillBeName/internal/pkg/auth/mocks"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock.NewMockAuthUsecase(ctrl)

	handler := NewAuthHandler(mockUsecase, nil)

	user := models.User{
		Login:    "testuser",
		Password: "testpassword",
	}

	body, _ := json.Marshal(map[string]string{
		"login":    user.Login,
		"password": user.Password,
	})

	mockUsecase.EXPECT().SignUp(context.Background(), user).Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler.SignUp(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestSignUp_InvalidJSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock.NewMockAuthUsecase(ctrl)
	handler := NewAuthHandler(mockUsecase, nil)

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/signup", bytes.NewBuffer([]byte(`invalid json`)))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler.SignUp(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestSignUp_CreateUserError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock.NewMockAuthUsecase(ctrl)
	handler := NewAuthHandler(mockUsecase, nil)
	mockUsecase.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(errors.New("internal error"))

	body, _ := json.Marshal(map[string]string{
		"login":    "testuser",
		"password": "testpass",
	})
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/signup", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	handler.SignUp(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestLogin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock.NewMockAuthUsecase(ctrl)

	mockUsecase.EXPECT().Login(gomock.Any(), "testuser", "testpass").Return("jwt_token", nil)

	handler := NewAuthHandler(mockUsecase, nil)

	body, _ := json.Marshal(map[string]string{
		"login":    "testuser",
		"password": "testpass",
	})
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/login", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	handler.Login(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	response := rr.Result()
	defer response.Body.Close()

	cookie := rr.Result().Cookies()[0]
	assert.Equal(t, "token", cookie.Name)
	assert.Equal(t, "jwt_token", cookie.Value)
}

func TestLogin_InvalidCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock.NewMockAuthUsecase(ctrl)

	mockUsecase.EXPECT().Login(gomock.Any(), "testuser", "wrongpass").Return("", errors.New("unauthorized"))

	handler := NewAuthHandler(mockUsecase, nil)

	body, _ := json.Marshal(map[string]string{
		"login":    "testuser",
		"password": "wrongpass",
	})
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/login", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	handler.Login(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}
