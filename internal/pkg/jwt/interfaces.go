package jwt

type JWTInterface interface {
	GenerateToken(userID uint, email, login string) (string, error)
	ParseToken(token string) (map[string]interface{}, error)
}
