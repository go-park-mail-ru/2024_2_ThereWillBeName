package middleware

import (
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type HashToken struct {
	Secret []byte
}

func NewHMACHashToken(secret string) (*HashToken, error) {
	return &HashToken{Secret: []byte(secret)}, nil
}

func (tk *HashToken) Create(userUUID uuid.UUID, tokenExpTime int64) (string, error) {
	h := hmac.New(sha256.New, tk.Secret)
	data := fmt.Sprintf("%s:%d", userUUID.String(), tokenExpTime)
	h.Write([]byte(data))
	token := hex.EncodeToString(h.Sum(nil)) + ":" + strconv.FormatInt(tokenExpTime, 10)
	return token, nil
}

func (tk *HashToken) Check(userUUID uuid.UUID, inputToken string) (bool, error) {
	tokenData := strings.Split(inputToken, ":")
	if len(tokenData) != 2 {
		return false, fmt.Errorf("invalid token format")
	}
	tokenExpStr := strings.TrimSpace(tokenData[1])
	tokenExp, err := strconv.ParseInt(tokenExpStr, 10, 64)
	if err != nil {
		return false, fmt.Errorf("invalid token expiration")
	}

	if tokenExp < time.Now().Unix() {
		return false, fmt.Errorf("token expired")
	}

	h := hmac.New(sha256.New, tk.Secret)
	data := fmt.Sprintf("%s:%d", userUUID.String(), tokenExp)
	h.Write([]byte(data))
	expectedMAC := h.Sum(nil)
	messageMAC, err := hex.DecodeString(tokenData[0])
	if err != nil {
		return false, fmt.Errorf("failed to decode token")
	}

	return hmac.Equal(messageMAC, expectedMAC), nil
}

func CSRFMiddleware(tk *HashToken, next http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		csrfToken := r.Header.Get("X-CSRF-Token")
		userUUIDStr := r.Header.Get("X-User-UUID")

		if userUUIDStr == "" || csrfToken == "" {
			response := httpresponse.ErrorResponse{
				Message: "CSRF token or UUID missing",
			}
			if logger != nil {
				logger.Error("CSRF token or UUID missing")
			}
			httpresponse.SendJSONResponse(w, response, http.StatusForbidden, logger)
			return
		}

		userUUID, err := uuid.Parse(userUUIDStr)
		if err != nil {
			response := httpresponse.ErrorResponse{
				Message: "Invalid UUID format",
			}
			if logger != nil {
				logger.Error("Invalid UUID format", slog.Any("error", err.Error()))
			}
			httpresponse.SendJSONResponse(w, response, http.StatusForbidden, logger)
			return
		}

		valid, err := tk.Check(userUUID, csrfToken)
		if err != nil || !valid {
			response := httpresponse.ErrorResponse{
				Message: "CSRF token validation failed",
			}
			if logger != nil {
				logger.Error("CSRF token validation failed", slog.Any("error", err.Error()))
			}
			httpresponse.SendJSONResponse(w, response, http.StatusForbidden, logger)
			return
		}

		next.ServeHTTP(w, r)
	})
}
