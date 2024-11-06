package http

// import (
// 	"2024_2_ThereWillBeName/internal/models"
// 	"errors"

// 	// "go.uber.org/mock/gomock"

// 	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
// 	mock "2024_2_ThereWillBeName/internal/pkg/user/mocks"
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestSignUp(t *testing.T) {
// 	// Create a mock for the user usecase
// 	userUsecase := new(mock.MockUserUsecase)
// 	jwtService := new(mock.MockJWT)

// 	handler := NewUserHandler(userUsecase, jwtService)

// 	tests := []struct {
// 		name           string
// 		credentials    Credentials
// 		mockSignUpID   int64
// 		mockSignUpErr  error
// 		mockToken      string
// 		expectedStatus int
// 		expectedBody   interface{}
// 	}{
// 		{
// 			name: "Successful signup",
// 			credentials: Credentials{
// 				Login:    "testuser",
// 				Email:    "test@example.com",
// 				Password: "securepassword",
// 			},
// 			mockSignUpID:   1,
// 			mockSignUpErr:  nil,
// 			mockToken:      "some.jwt.token",
// 			expectedStatus: http.StatusOK,
// 			expectedBody: models.User{
// 				ID:    1,
// 				Login: "testuser",
// 				Email: "test@example.com",
// 			},
// 		},
// 		{
// 			name: "User already exists",
// 			credentials: Credentials{
// 				Login:    "existinguser",
// 				Email:    "existing@example.com",
// 				Password: "securepassword",
// 			},
// 			mockSignUpID:   0,
// 			mockSignUpErr:  models.ErrAlreadyExists,
// 			expectedStatus: http.StatusConflict,
// 			expectedBody: httpresponse.ErrorResponse{
// 				Message: "user already exists",
// 			},
// 		},
// 		{
// 			name: "Failed registration",
// 			credentials: Credentials{
// 				Login:    "failuser",
// 				Email:    "fail@example.com",
// 				Password: "securepassword",
// 			},
// 			mockSignUpID:   0,
// 			mockSignUpErr:  errors.New("registration failed"),
// 			expectedStatus: http.StatusInternalServerError,
// 			expectedBody: httpresponse.ErrorResponse{
// 				Message: "Registration failed",
// 			},
// 		},
// 		{
// 			name: "Token generation failed",
// 			credentials: Credentials{
// 				Login:    "tokenfailuser",
// 				Email:    "tokenfail@example.com",
// 				Password: "securepassword",
// 			},
// 			mockSignUpID:   1,
// 			mockSignUpErr:  nil,
// 			mockToken:      "",
// 			expectedStatus: http.StatusInternalServerError,
// 			expectedBody: httpresponse.ErrorResponse{
// 				Message: "Token generation failed",
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Prepare the mock responses
// 			userUsecase.On("SignUp", mock.Anything, mock.AnythingOfType("models.User")).
// 				Return(tt.mockSignUpID, tt.mockSignUpErr)

// 			if tt.mockToken != "" {
// 				jwtService.On("GenerateToken", tt.mockSignUpID, tt.credentials.Email, tt.credentials.Login).
// 					Return(tt.mockToken, nil)
// 			} else {
// 				jwtService.On("GenerateToken", mock.Anything, mock.Anything, mock.Anything).
// 					Return("", errors.New("token generation failed"))
// 			}

// 			// Prepare the request
// 			body, _ := json.Marshal(tt.credentials)
// 			req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
// 			w := httptest.NewRecorder()

// 			// Call the handler
// 			handler.SignUp(w, req)

// 			// Check the response
// 			res := w.Result()
// 			assert.Equal(t, tt.expectedStatus, res.StatusCode)

// 			var responseBody interface{}
// 			if tt.expectedStatus == http.StatusOK {
// 				var userResponse models.User
// 				json.NewDecoder(res.Body).Decode(&userResponse)
// 				responseBody = userResponse
// 			} else {
// 				var errorResponse httpresponse.ErrorResponse
// 				json.NewDecoder(res.Body).Decode(&errorResponse)
// 				responseBody = errorResponse
// 			}

// 			assert.Equal(t, tt.expectedBody, responseBody)
// 			// Assert that the mock expectations were met
// 			userUsecase.AssertExpectations(t)
// 			jwtService.AssertExpectations(t)
// 		})
// 	}
// }

// // func TestSignUp_InvalidJSON(t *testing.T) {
// // 	ctrl := gomock.NewController(t)
// // 	defer ctrl.Finish()

// // 	mockUsecase := mock.NewMockAuthUsecase(ctrl)
// // 	handler := NewAuthHandler(mockUsecase, nil)

// // 	req, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/signup", bytes.NewBuffer([]byte(`invalid json`)))
// // 	req.Header.Set("Content-Type", "application/json")

// // 	rr := httptest.NewRecorder()

// // 	handler.SignUp(rr, req)

// // 	assert.Equal(t, http.StatusBadRequest, rr.Code)
// // }

// // func TestSignUp_CreateUserError(t *testing.T) {
// // 	ctrl := gomock.NewController(t)
// // 	defer ctrl.Finish()

// // 	mockUsecase := mock.NewMockAuthUsecase(ctrl)
// // 	handler := NewAuthHandler(mockUsecase, nil)
// // 	mockUsecase.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(errors.New("internal error"))

// // 	body, _ := json.Marshal(map[string]string{
// // 		"login":    "testuser",
// // 		"password": "testpass",
// // 	})
// // 	req, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/signup", bytes.NewBuffer(body))
// // 	rr := httptest.NewRecorder()

// // 	handler.SignUp(rr, req)

// // 	assert.Equal(t, http.StatusInternalServerError, rr.Code)
// // }

// // func TestLogin_Success(t *testing.T) {
// // 	ctrl := gomock.NewController(t)
// // 	defer ctrl.Finish()

// // 	mockUsecase := mock.NewMockAuthUsecase(ctrl)

// // 	mockUsecase.EXPECT().Login(gomock.Any(), "testuser", "testpass").Return("jwt_token", nil)

// // 	handler := NewAuthHandler(mockUsecase, nil)

// // 	body, _ := json.Marshal(map[string]string{
// // 		"login":    "testuser",
// // 		"password": "testpass",
// // 	})
// // 	req, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/login", bytes.NewBuffer(body))
// // 	rr := httptest.NewRecorder()

// // 	handler.Login(rr, req)

// // 	assert.Equal(t, http.StatusOK, rr.Code)

// // 	response := rr.Result()
// // 	defer response.Body.Close()

// // 	cookie := response.Cookies()[0]
// // 	assert.Equal(t, "token", cookie.Name)
// // 	assert.Equal(t, "jwt_token", cookie.Value)
// // }

// // func TestLogin_InvalidCredentials(t *testing.T) {
// // 	ctrl := gomock.NewController(t)
// // 	defer ctrl.Finish()

// // 	mockUsecase := mock.NewMockAuthUsecase(ctrl)

// // 	mockUsecase.EXPECT().Login(gomock.Any(), "testuser", "wrongpass").Return("", errors.New("unauthorized"))

// // 	handler := NewAuthHandler(mockUsecase, nil)

// // 	body, _ := json.Marshal(map[string]string{
// // 		"login":    "testuser",
// // 		"password": "wrongpass",
// // 	})
// // 	req, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/login", bytes.NewBuffer(body))
// // 	rr := httptest.NewRecorder()

// // 	handler.Login(rr, req)

// // 	assert.Equal(t, http.StatusUnauthorized, rr.Code)
// // }

// // func TestCurrentUser_Success(t *testing.T) {
// // 	ctrl := gomock.NewController(t)
// // 	defer ctrl.Finish()

// // 	mockUsecase := mock.NewMockAuthUsecase(ctrl)
// // 	handler := NewAuthHandler(mockUsecase, nil)

// // 	ctx := context.WithValue(context.Background(), middleware.IdKey, uint(1))
// // 	ctx = context.WithValue(ctx, middleware.LoginKey, "testuser")

// // 	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/users/me", nil)
// // 	rr := httptest.NewRecorder()

// // 	handler.CurrentUser(rr, req)

// // 	assert.Equal(t, http.StatusOK, rr.Code)

// // 	var user models.User
// // 	if err := json.NewDecoder(rr.Body).Decode(&user); err != nil {
// // 		t.Fatal(err)
// // 	}

// // 	assert.Equal(t, uint(1), user.ID)
// // 	assert.Equal(t, "testuser", user.Login)
// // }

// // func TestCurrentUser_NoUserID(t *testing.T) {
// // 	ctrl := gomock.NewController(t)
// // 	defer ctrl.Finish()

// // 	mockUsecase := mock.NewMockAuthUsecase(ctrl)
// // 	handler := NewAuthHandler(mockUsecase, nil)

// // 	ctx := context.WithValue(context.Background(), middleware.LoginKey, "testuser")
// // 	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/users/me", nil)
// // 	rr := httptest.NewRecorder()

// // 	handler.CurrentUser(rr, req)

// // 	assert.Equal(t, http.StatusUnauthorized, rr.Code)
// // 	assert.Equal(t, "{\"message\":\"User is not authorized\"}\n", rr.Body.String())
// // }

// // func TestCurrentUser_NoLogin(t *testing.T) {
// // 	ctrl := gomock.NewController(t)
// // 	defer ctrl.Finish()

// // 	mockUsecase := mock.NewMockAuthUsecase(ctrl)
// // 	handler := NewAuthHandler(mockUsecase, nil)

// // 	ctx := context.WithValue(context.Background(), middleware.IdKey, uint(1))
// // 	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/users/me", nil)
// // 	rr := httptest.NewRecorder()

// // 	handler.CurrentUser(rr, req)

// // 	assert.Equal(t, http.StatusUnauthorized, rr.Code)
// // 	assert.Equal(t, "{\"message\":\"User is not authorized\"}\n", rr.Body.String())
// // }

// // func TestLogout_Success(t *testing.T) {
// // 	ctrl := gomock.NewController(t)
// // 	defer ctrl.Finish()

// // 	mockUsecase := mock.NewMockAuthUsecase(ctrl)
// // 	handler := NewAuthHandler(mockUsecase, nil)

// // 	req, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/logout", nil)
// // 	rr := httptest.NewRecorder()

// // 	handler.Logout(rr, req)

// // 	assert.Equal(t, http.StatusOK, rr.Code)
// // }
