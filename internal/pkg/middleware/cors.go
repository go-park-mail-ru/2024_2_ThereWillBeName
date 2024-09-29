package middleware

import "net/http"

type CORSMiddleware struct {
	AllowedOrigins []string
}

func NewCORSMiddleware(allowedOrigins []string) *CORSMiddleware {
	return &CORSMiddleware{
		AllowedOrigins: allowedOrigins,
	}
}
func (c *CORSMiddleware) CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "POST,PUT,DELETE,GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if len(c.AllowedOrigins) > 0 {
			w.Header().Set("Access-Control-Allow-Origin", c.AllowedOrigins[0])
		}
	},
	)
}
