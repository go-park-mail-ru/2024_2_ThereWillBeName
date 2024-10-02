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

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := models.User{
		Login:    credentials.Login,
		Password: credentials.Password,
	}

	if err := h.usecase.SignUp(context.Background(), user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.usecase.Login(context.Background(), credentials.Login, credentials.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
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

func (h *Handler) CurrentUser(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {
		http.Error(w, "Пользователь не авторизирован", http.StatusUnauthorized)
		return
	}

	login, ok := r.Context().Value(middleware.LoginKey).(string)
	if !ok {
		http.Error(w, "Пользователь не авторизирован", http.StatusUnauthorized)
		return
	}

	response := models.User{
		ID:    userID,
		Login: login,
	}

	httpresponse.SendJSONResponse(w, response, http.StatusOK)
}
