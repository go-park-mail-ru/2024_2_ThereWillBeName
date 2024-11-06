package middleware

import (
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"2024_2_ThereWillBeName/internal/pkg/jwt"
	"context"
	"log/slog"
	"net/http"
)

type contextKey string

const (
	IdKey    contextKey = "userID"
	LoginKey contextKey = "login"
	EmailKey contextKey = "email"
)

func MiddlewareAuth(jwtService *jwt.JWT, next http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("token")
		if err != nil {
			response := httpresponse.ErrorResponse{
				Message: "Cookie not found",
			}

			logger.Error("Cookie not found", slog.Any("error", err.Error()))
			httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized, logger)
			return
		}

		claims, err := jwtService.ParseToken(cookie.Value)
		if err != nil {
			response := httpresponse.ErrorResponse{
				Message: "Invalid token",
			}

			if logger != nil {
				logger.Error("Invalid token", slog.Any("error", err))
			}
			httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized, logger)
			return
		}
		userID := uint(claims["id"].(float64))
		login := claims["login"].(string)
		email := claims["email"].(string)
		if logger != nil {
			logger.Info("Token parsed", slog.String("login", login), slog.String("email", email))
		}
		ctx := context.WithValue(r.Context(), IdKey, userID)
		ctx = context.WithValue(ctx, LoginKey, login)
		ctx = context.WithValue(ctx, EmailKey, email)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
