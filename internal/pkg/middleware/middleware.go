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

func MiddlewareAuth(jwtService jwt.JWTInterface, next http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-CSRF-Token")

		if token == "" {
			response := httpresponse.ErrorResponse{
				Message: "Token is missing",
			}
			if logger != nil {
				logger.Error("Token is missing")
			}
			httpresponse.SendJSONResponse(w, response, http.StatusForbidden, logger)
			return
		}
		claims, err := jwtService.ParseToken(token)
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
		userID := uint(claims["userID"].(float64))
		login := claims["login"].(string)
		email := claims["email"].(string)
		if logger != nil {
			logger.Info("Token parsed", slog.Int("userID", int(userID)), slog.String("login", login), slog.String("email", email))
		}
		ctx := context.WithValue(r.Context(), IdKey, userID)
		ctx = context.WithValue(ctx, LoginKey, login)
		ctx = context.WithValue(ctx, EmailKey, email)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
