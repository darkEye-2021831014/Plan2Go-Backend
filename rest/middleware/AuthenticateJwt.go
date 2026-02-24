package middleware

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

type claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}


func (m *ConfigMiddleware) AuthenticateJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// 1. Get Authorization header
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		// 2. Remove "Bearer " prefix
		if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
			tokenStr = tokenStr[7:]
		}

		// 3. Parse token
		claims := &claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.cnf.Jwt_SecretKey), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// 4. Set user info in header (string conversion fixed)
		r.Header.Set("userID", strconv.Itoa(claims.UserID))
		r.Header.Set("username", claims.Username)

		// 5. Call next handler
		next.ServeHTTP(w, r)
	})
}
