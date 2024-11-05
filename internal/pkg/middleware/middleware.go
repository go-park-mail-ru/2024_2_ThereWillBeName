package middleware

import (
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"2024_2_ThereWillBeName/internal/pkg/jwt"
	"context"
	"log"
	"net/http"
)

type contextKey string

const (
	IdKey    contextKey = "userID"
	LoginKey contextKey = "login"
	EmailKey contextKey = "email"
)

func MiddlewareAuth(jwtService *jwt.JWT, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			response := httpresponse.ErrorResponse{
				Message: "Cookie not found",
			}
			httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized)
			return
		}

		claims, err := jwtService.ParseToken(cookie.Value)
		if err != nil {
			response := httpresponse.ErrorResponse{
				Message: "Invalid token",
			}
			httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized)
			return
		}
		userID := uint(claims["id"].(float64))
		login := claims["login"].(string)
		email := claims["email"].(string)
		log.Println("middleware: ", userID, IdKey)
		ctx := context.WithValue(r.Context(), IdKey, userID)
		ctx = context.WithValue(ctx, LoginKey, login)
		ctx = context.WithValue(ctx, EmailKey, email)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
