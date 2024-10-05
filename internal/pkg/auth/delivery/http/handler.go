package http

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/auth"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"2024_2_ThereWillBeName/internal/pkg/jwt"
	"2024_2_ThereWillBeName/internal/pkg/middleware"

	"context"
	"encoding/json"
	"net/http"
)

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Handler struct {
	usecase auth.AuthUsecase
	jwt     *jwt.JWT
}

func NewAuthHandler(usecase auth.AuthUsecase, jwt *jwt.JWT) *Handler {
	return &Handler{
		usecase: usecase,
		jwt:     jwt,
	}
}

// SignUp godoc
// @Summary Sign up a new user
// @Description Create a new user with login and password
// @Accept json
// @Produce json
// @Param credentials body Credentials true "User credentials"
// @Success 201 {object} models.User "User created successfully"
// @Failure 400 {object} httpresponses.ErrorResponse "Bad Request"
// @Failure 500 {object} httpresponses.ErrorResponse "Internal Server Error"
// @Router /signup [post]
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		response := httpresponse.ErrorResponse{
			Message: err.Error(),
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	user := models.User{
		Login:    credentials.Login,
		Password: credentials.Password,
	}

	if err := h.usecase.SignUp(context.Background(), user); err != nil {
		response := httpresponse.ErrorResponse{
			Message: err.Error(),
		}
		httpresponse.SendJSONResponse(w, response, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Login godoc
// @Summary Login a user
// @Description Authenticate a user and return a token
// @Accept json
// @Produce json
// @Param credentials body Credentials true "User credentials"
// @Success 200 {string} string "Token"
// @Failure 400 {object} httpresponses.ErrorResponse "Bad Request"
// @Failure 401 {object} httpresponses.ErrorResponse "Unauthorized"
// @Router /login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		response := httpresponse.ErrorResponse{
			Message: err.Error(),
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	token, err := h.usecase.Login(context.Background(), credentials.Login, credentials.Password)
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: err.Error(),
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	})

	w.WriteHeader(http.StatusOK)
}

// Logout godoc
// @Summary Logout a user
// @Description Log out the user by clearing the authentication token
// @Produce json
// @Success 200 {string} string "Logged out successfully"
// @Router /logout [post]
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	})

	w.WriteHeader(http.StatusOK)
}

// CurrentUser godoc
// @Summary Get the current user
// @Description Retrieve the current authenticated user information
// @Produce json
// @Success 200 {object} models.User "Current user"
// @Failure 401 {object} httpresponses.ErrorResponse "Unauthorized"
// @Failure 500 {object} httpresponses.ErrorResponse "Internal Server Error"
// @Router /users/me [get]
func (h *Handler) CurrentUser(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {
		response := httpresponse.ErrorResponse{
			Message: "Пользователь не авторизирован",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized)
		return
	}

	login, ok := r.Context().Value(middleware.LoginKey).(string)
	if !ok {
		response := httpresponse.ErrorResponse{
			Message: "Пользователь не авторизирован",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized)
		return
	}

	response := models.User{
		ID:    userID,
		Login: login,
	}

	httpresponse.SendJSONResponse(w, response, http.StatusOK)
}
