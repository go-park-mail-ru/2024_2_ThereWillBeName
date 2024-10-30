package middleware

import (
	//"crypto/rand"
	"encoding/hex"
	//"net/http"
	"crypto/hmac"
	"crypto/sha256"
)

type HashToken struct {
	Secret []byte
}

func NewHMACHashToken(secret string) (*HashToken, error) {
	return &HashToken{Secret: []byte(secret)}, nil
}

func (tk *HashToken) Create(s *Session, tokenExpTime int64) (string, error) {
	h := hmac.New(sha256.New, []byte(tk.Secret))
	data := fmt.Sprintf("%s:%d:%d", s.ID, s.UserID, tokenExpTime)
	h.Write([]byte(data))
	token := hex.EncodeToString(h.Sum(nil)) + ":" + strconv.FormatInt(tokenExpTime, 10)
	return token, nil
}
