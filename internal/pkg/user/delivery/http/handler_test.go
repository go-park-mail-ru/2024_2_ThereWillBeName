package http

import (
	"2024_2_ThereWillBeName/internal/models"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"2024_2_ThereWillBeName/internal/pkg/jwt/mocks"
	"2024_2_ThereWillBeName/internal/pkg/middleware"
	usermock "2024_2_ThereWillBeName/internal/pkg/user/mocks"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSignUpHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	handl := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handl)

	mockUsecase := usermock.NewMockUserUsecase(ctrl)
	mockJWT := mocks.NewMockJWTInterface(ctrl)
	handler := NewUserHandler(mockUsecase, mockJWT, logger)

	tests := []struct {
		name             string
		inputCredentials Credentials
		usecaseErr       error
		jwtErr           error
		expectedStatus   int
		expectedBody     httpresponse.ErrorResponse
	}{
		{
			name: "token generation failed",
			inputCredentials: Credentials{
				Login:    "tokenuser",
				Email:    "token@example.com",
				Password: "password123",
			},
			usecaseErr:     nil,
			jwtErr:         errors.New("token generation failed"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   httpresponse.ErrorResponse{Message: "Token generation failed"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(uint(1), tt.usecaseErr)

			// Мокаем вызов GenerateToken
			mockJWT.EXPECT().
				GenerateToken(gomock.Any(), gomock.Any(), gomock.Any()).
				Return("mocked_token", tt.jwtErr).
				Times(1)

			reqBody, _ := json.Marshal(tt.inputCredentials)
			req := httptest.NewRequest("POST", "/signup", bytes.NewReader(reqBody))
			rec := httptest.NewRecorder()

			handler.SignUp(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)

			if tt.expectedStatus != http.StatusOK {
				var response httpresponse.ErrorResponse
				_ = json.NewDecoder(rec.Body).Decode(&response)
				assert.Equal(t, tt.expectedBody.Message, response.Message)
			}
		})
	}
}

func TestLogoutHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	handl := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handl)
	// Мокаем необходимые зависимости (например, логгер)
	handler := &Handler{
		logger: logger,
	}

	// Мокаем контекст запроса
	req := httptest.NewRequest("POST", "/logout", nil)
	rec := httptest.NewRecorder()

	// Выполняем запрос
	handler.Logout(rec, req)

	// Проверяем, что статус ответа 200 OK
	assert.Equal(t, http.StatusOK, rec.Code)

	// Проверяем, что cookie "token" имеет пустое значение и MaxAge = -1 (удаление cookie)
	cookies := rec.Result().Cookies()
	var tokenCookie *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == "token" {
			tokenCookie = cookie
			break
		}
	}

	assert.NotNil(t, tokenCookie, "Expected cookie 'token' to be set")
	assert.Equal(t, "", tokenCookie.Value, "Expected token cookie to be empty")
	assert.Equal(t, -1, tokenCookie.MaxAge, "Expected MaxAge to be -1 for token cookie removal")
}

func TestCurrentUserHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	handl := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handl)

	// Мокаем необходимые зависимости
	handler := &Handler{
		logger: logger,
	}

	// Создаем контекст с нужными данными для успешного запроса
	ctx := context.WithValue(context.Background(), middleware.IdKey, uint(1))
	ctx = context.WithValue(ctx, middleware.LoginKey, "userlogin")
	ctx = context.WithValue(ctx, middleware.EmailKey, "user@example.com")

	// Создаем запрос и записываем результат
	req := httptest.NewRequest("GET", "/currentuser", nil).WithContext(ctx)
	rec := httptest.NewRecorder()

	// Выполняем запрос
	handler.CurrentUser(rec, req)

	// Проверяем, что статус ответа 200 OK
	assert.Equal(t, http.StatusOK, rec.Code)

	// Проверяем, что в ответе возвращены правильные данные пользователя
	var response models.User
	_ = json.NewDecoder(rec.Body).Decode(&response)
	assert.Equal(t, uint(1), response.ID)
	assert.Equal(t, "userlogin", response.Login)
	assert.Equal(t, "user@example.com", response.Email)
}

func TestCurrentUserHandler_Unauthorized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	handl := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handl)

	// Мокаем необходимые зависимости
	handler := &Handler{
		logger: logger,
	}

	// Создаем контекст без нужных данных
	ctx := context.Background()

	// Создаем запрос и записываем результат
	req := httptest.NewRequest("GET", "/currentuser", nil).WithContext(ctx)
	rec := httptest.NewRecorder()

	// Выполняем запрос
	handler.CurrentUser(rec, req)

	// Проверяем, что статус ответа 401 Unauthorized
	assert.Equal(t, http.StatusUnauthorized, rec.Code)

	// Проверяем, что в ответе сообщение об ошибке
	var response httpresponse.ErrorResponse
	_ = json.NewDecoder(rec.Body).Decode(&response)
	assert.Equal(t, "User is not authorized", response.Message)
}

func TestUploadAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	handl := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handl)

	mockUsecase := usermock.NewMockUserUsecase(ctrl)
	mockJWT := mocks.NewMockJWTInterface(ctrl)
	handler := NewUserHandler(mockUsecase, mockJWT, logger)

	// Тестовые данные
	tests := []struct {
		name           string
		userID         string
		authUserID     uint
		usecaseErr     error
		uploadSuccess  bool
		expectedStatus int
		expectedBody   httpresponse.ErrorResponse
	}{
		{
			name:           "successful avatar upload",
			userID:         "1",
			authUserID:     1,
			usecaseErr:     nil,
			uploadSuccess:  true,
			expectedStatus: http.StatusOK,
			expectedBody:   httpresponse.ErrorResponse{}, // Пустой, так как успешный запрос
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.uploadSuccess {
				mockUsecase.EXPECT().UploadAvatar(gomock.Any(), tt.authUserID, gomock.Any(), gomock.Any()).Return("mocked/avatar/path", tt.usecaseErr).Times(1)
			} else {
				mockUsecase.EXPECT().UploadAvatar(gomock.Any(), tt.authUserID, gomock.Any(), gomock.Any()).Return("", tt.usecaseErr).Times(1)
			}

			if tt.uploadSuccess {
				mockUsecase.EXPECT().UploadAvatar(gomock.Any(), tt.authUserID, gomock.Any(), gomock.Any()).Return("mocked/avatar/path", tt.usecaseErr).Times(1)
			} else {
				mockUsecase.EXPECT().UploadAvatar(gomock.Any(), tt.authUserID, gomock.Any(), gomock.Any()).Return("", tt.usecaseErr).Times(1)
			}

			// Создаем тело запроса с multipart-данными
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)
			part, _ := writer.CreateFormFile("avatar", "avatar.png")
			part.Write([]byte("dummy avatar content")) // имитация содержания файла
			_ = writer.Close()

			// Создаем запрос
			req := httptest.NewRequest("POST", "/upload"+tt.userID, body)
			req.Header.Set("Content-Type", writer.FormDataContentType())
			ctx := context.WithValue(req.Context(), middleware.IdKey, tt.authUserID)
			req = req.WithContext(ctx)

			// Ответ
			rec := httptest.NewRecorder()

			// Выполняем запрос
			handler.UploadAvatar(rec, req)

			// Проверяем статус ответа
			assert.Equal(t, tt.expectedStatus, rec.Code)

			// Если это не успешный ответ, проверяем тело ответа
			if tt.expectedStatus != http.StatusOK {
				var response httpresponse.ErrorResponse
				_ = json.NewDecoder(rec.Body).Decode(&response)
				assert.Equal(t, tt.expectedBody.Message, response.Message)
			}
		})
	}
}
