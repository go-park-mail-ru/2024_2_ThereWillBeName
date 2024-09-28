package http

import (
    "context"
    "encoding/json"
    "net/http"
    "2024_2_ThereWillBeName/internal/models"
    "2024_2_ThereWillBeName/internal/pkg/auth"
    "2024_2_ThereWillBeName/internal/pkg/jwt"
)

type Handler struct {
    usecase auth.AuthUsecase
    jwt     *jwt.JWT
}

func NewHandler(usecase auth.AuthUsecase, jwt *jwt.JWT) *Handler {
    return &Handler{
        usecase: usecase,
        jwt: jwt,
    }
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := h.usecase.SignUp(context.Background(), user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
    var credentials struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    token, err := h.usecase.Login(context.Background(), credentials.Email, credentials.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    http.SetCookie(w, &http.Cookie{
        Name:     "token",
        Value:    token,
        Path:     "/",
        HttpOnly: true,
        Secure:   true,
    })

    w.WriteHeader(http.StatusOK)
}

func (h *Handler) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("token")
        if err != nil {
            http.Error(w, "Cookie not found", http.StatusUnauthorized)
            return
        }

        _, err = h.jwt.ParseToken(cookie.Value)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        //логика для работы с claims

        next.ServeHTTP(w, r)
    })
}
