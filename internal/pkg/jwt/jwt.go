package jwt

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	secret []byte
	logger *slog.Logger
}

func NewJWT(secret string, logger *slog.Logger) *JWT {
	return &JWT{
		secret: []byte(secret),
	}
}

func (j *JWT) GenerateToken(userID uint, email, login string) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"email":  email,
		"login":  login,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}

func (j *JWT) ParseToken(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.secret, nil
	})

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		j.logger.Debug("from Generate token", "JWT_claims", claims)
		return claims, nil
	}

	return nil, err
}
