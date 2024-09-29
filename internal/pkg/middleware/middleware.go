package middleware

import (
	"2024_2_ThereWillBeName/internal/pkg/jwt"
	"context"
	"net/http"
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

		userID := claims["id"].(uint)
		login := claims["login"].(string)
		ctx := context.WithValue(r.Context(), "userID", userID)
		ctx = context.WithValue(ctx, "login", login)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
