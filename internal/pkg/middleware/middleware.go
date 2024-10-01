package middleware

import (
	"2024_2_ThereWillBeName/internal/pkg/jwt"
	"context"
	"log"
	"net/http"
)

type contextKey string

const (
	IdKey    contextKey = "userID"
	LoginKey contextKey = "login"
)

func MiddlewareAuth(jwtService *jwt.JWT, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Cookie not found", http.StatusUnauthorized)
			return
		}

		claims, err := jwtService.ParseToken(cookie.Value)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		userID := uint(claims["id"].(float64))
		login := claims["login"].(string)
		log.Println("middleware: ", userID, IdKey)
		ctx := context.WithValue(r.Context(), IdKey, userID)
		ctx = context.WithValue(ctx, LoginKey, login)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
