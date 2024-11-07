package http

// import (
// 	"2024_2_ThereWillBeName/internal/models"
// 	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
// 	"2024_2_ThereWillBeName/internal/pkg/jwt/mocks"
// 	"2024_2_ThereWillBeName/internal/pkg/middleware"
// 	usermock "2024_2_ThereWillBeName/internal/pkg/user/mocks"
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"errors"
// 	"log/slog"
// 	"mime/multipart"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"strconv"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/gorilla/mux"
// 	"github.com/stretchr/testify/assert"
// )

// func TestSignUpHandler(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	opts := &slog.HandlerOptions{
// 		Level: slog.LevelDebug,
// 	}
// 	handl := slog.NewJSONHandler(os.Stdout, opts)
// 	logger := slog.New(handl)

// 	mockUsecase := usermock.NewMockUserUsecase(ctrl)
// 	mockJWT := mocks.NewMockJWTInterface(ctrl)
// 	handler := NewUserHandler(mockUsecase, mockJWT, logger)

// 	tests := []struct {
// 		name             string
// 		inputCredentials Credentials
// 		usecaseErr       error
// 		jwtErr           error
// 		expectedStatus   int
// 		expectedBody     httpresponse.ErrorResponse
// 	}{
// 		{
// 			name: "token generation failed",
// 			inputCredentials: Credentials{
// 				Login:    "tokenuser",
// 				Email:    "token@example.com",
// 				Password: "password123",
// 			},
// 			usecaseErr:     nil,
// 			jwtErr:         errors.New("token generation failed"),
// 			expectedStatus: http.StatusInternalServerError,
// 			expectedBody:   httpresponse.ErrorResponse{Message: "Token generation failed"},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mockUsecase.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(uint(1), tt.usecaseErr)

// 			mockJWT.EXPECT().
// 				GenerateToken(gomock.Any(), gomock.Any(), gomock.Any()).
// 				Return("mocked_token", tt.jwtErr).
// 				Times(1)

// 			reqBody, _ := json.Marshal(tt.inputCredentials)
// 			req := httptest.NewRequest("POST", "/signup", bytes.NewReader(reqBody))
// 			rec := httptest.NewRecorder()

// 			handler.SignUp(rec, req)

// 			assert.Equal(t, tt.expectedStatus, rec.Code)

// 			if tt.expectedStatus != http.StatusOK {
// 				var response httpresponse.ErrorResponse
// 				_ = json.NewDecoder(rec.Body).Decode(&response)
// 				assert.Equal(t, tt.expectedBody.Message, response.Message)
// 			}
// 		})
// 	}
// }

// func TestLogoutHandler_Success(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	opts := &slog.HandlerOptions{
// 		Level: slog.LevelDebug,
// 	}
// 	handl := slog.NewJSONHandler(os.Stdout, opts)
// 	logger := slog.New(handl)
// 	handler := &Handler{
// 		logger: logger,
// 	}

// 	req := httptest.NewRequest("POST", "/logout", nil)
// 	rec := httptest.NewRecorder()

// 	handler.Logout(rec, req)

// 	assert.Equal(t, http.StatusOK, rec.Code)

// 	cookies := rec.Result().Cookies()
// 	var tokenCookie *http.Cookie
// 	for _, cookie := range cookies {
// 		if cookie.Name == "token" {
// 			tokenCookie = cookie
// 			break
// 		}
// 	}

// 	assert.NotNil(t, tokenCookie, "Expected cookie 'token' to be set")
// 	assert.Equal(t, "", tokenCookie.Value, "Expected token cookie to be empty")
// 	assert.Equal(t, -1, tokenCookie.MaxAge, "Expected MaxAge to be -1 for token cookie removal")
// }

// func TestCurrentUserHandler_Success(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	opts := &slog.HandlerOptions{
// 		Level: slog.LevelDebug,
// 	}
// 	handl := slog.NewJSONHandler(os.Stdout, opts)
// 	logger := slog.New(handl)

// 	handler := &Handler{
// 		logger: logger,
// 	}

// 	ctx := context.WithValue(context.Background(), middleware.IdKey, uint(1))
// 	ctx = context.WithValue(ctx, middleware.LoginKey, "userlogin")
// 	ctx = context.WithValue(ctx, middleware.EmailKey, "user@example.com")

// 	req := httptest.NewRequest("GET", "/currentuser", nil).WithContext(ctx)
// 	rec := httptest.NewRecorder()

// 	handler.CurrentUser(rec, req)

// 	assert.Equal(t, http.StatusOK, rec.Code)

// 	var response models.User
// 	_ = json.NewDecoder(rec.Body).Decode(&response)
// 	assert.Equal(t, uint(1), response.ID)
// 	assert.Equal(t, "userlogin", response.Login)
// 	assert.Equal(t, "user@example.com", response.Email)
// }

// func TestCurrentUserHandler_Unauthorized(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	opts := &slog.HandlerOptions{
// 		Level: slog.LevelDebug,
// 	}
// 	handl := slog.NewJSONHandler(os.Stdout, opts)
// 	logger := slog.New(handl)

// 	handler := &Handler{
// 		logger: logger,
// 	}

// 	ctx := context.Background()

// 	req := httptest.NewRequest("GET", "/currentuser", nil).WithContext(ctx)
// 	rec := httptest.NewRecorder()

// 	handler.CurrentUser(rec, req)

// 	assert.Equal(t, http.StatusUnauthorized, rec.Code)

// 	var response httpresponse.ErrorResponse
// 	_ = json.NewDecoder(rec.Body).Decode(&response)
// 	assert.Equal(t, "User is not authorized", response.Message)
// }

// func TestUploadAvatar(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	opts := &slog.HandlerOptions{
// 		Level: slog.LevelDebug,
// 	}
// 	handl := slog.NewJSONHandler(os.Stdout, opts)
// 	logger := slog.New(handl)

// 	mockUsecase := usermock.NewMockUserUsecase(ctrl)
// 	mockJWT := mocks.NewMockJWTInterface(ctrl)
// 	handler := NewUserHandler(mockUsecase, mockJWT, logger)

// 	tests := []struct {
// 		name           string
// 		userID         uint
// 		authUserID     uint
// 		usecaseErr     error
// 		uploadSuccess  bool
// 		expectedStatus int
// 		expectedBody   httpresponse.ErrorResponse
// 	}{
// 		{
// 			name:           "successful avatar upload",
// 			userID:         3,
// 			authUserID:     3,
// 			usecaseErr:     nil,
// 			uploadSuccess:  true,
// 			expectedStatus: http.StatusOK,
// 			expectedBody:   httpresponse.ErrorResponse{},
// 		},
// 		{
// 			name:           "unauthorized user",
// 			userID:         1,
// 			authUserID:     2,
// 			usecaseErr:     nil,
// 			uploadSuccess:  false,
// 			expectedStatus: http.StatusUnauthorized,
// 			expectedBody:   httpresponse.ErrorResponse{Message: "User is not authorized to upload avatar for this ID"},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if tt.uploadSuccess && tt.userID == tt.authUserID {
// 				mockUsecase.EXPECT().UploadAvatar(gomock.Any(), tt.authUserID, gomock.Any(), gomock.Any()).
// 					Return("mocked/avatar/path", tt.usecaseErr).Times(1)
// 			} else if tt.userID == tt.authUserID {
// 				mockUsecase.EXPECT().UploadAvatar(gomock.Any(), tt.authUserID, gomock.Any(), gomock.Any()).
// 					Return("", tt.usecaseErr).Times(1)
// 			}

// 			body := new(bytes.Buffer)
// 			writer := multipart.NewWriter(body)
// 			part, _ := writer.CreateFormFile("avatar", "avatar.png")
// 			_, err := part.Write([]byte("dummy avatar content"))
// 			if err != nil {
// 				t.Errorf("failed to write to part: %v", err)
// 			}
// 			_ = writer.Close()

// 			req := httptest.NewRequest(http.MethodPut, "/users/"+strconv.Itoa(int(tt.userID))+"/avatars", body)
// 			req.Header.Set("Content-Type", writer.FormDataContentType())

// 			ctx := context.WithValue(req.Context(), middleware.IdKey, uint(tt.authUserID))
// 			req = req.WithContext(ctx)

// 			router := mux.NewRouter()
// 			router.HandleFunc("/users/{userID}/avatars", handler.UploadAvatar).Methods(http.MethodPut)

// 			rec := httptest.NewRecorder()
// 			router.ServeHTTP(rec, req)

// 			assert.Equal(t, tt.expectedStatus, rec.Code)

// 			if tt.expectedStatus != http.StatusOK {
// 				var response httpresponse.ErrorResponse
// 				_ = json.NewDecoder(rec.Body).Decode(&response)
// 				assert.Equal(t, tt.expectedBody.Message, response.Message)
// 			} else {
// 				var response map[string]string
// 				_ = json.NewDecoder(rec.Body).Decode(&response)
// 				assert.Equal(t, "Avatar uploaded successfully", response["message"])
// 			}
// 		})
// 	}
// }
