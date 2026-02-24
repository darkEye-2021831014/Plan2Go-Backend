package user

import (
	"net/http"
	"plan2go-backend/config"
	"plan2go-backend/util"
)

func (h *Handler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	tokenString := authHeader
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		tokenString = authHeader[7:]
	}

	cnf := config.GetConfig()
	claims, err := util.VerifyToken(tokenString, cnf.Jwt_SecretKey)
	if err != nil {
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}

	email := claims.Email
	if email == "" {
		http.Error(w, "Token missing email", http.StatusUnauthorized)
		return
	}

	foundUser, err := h.userRepo.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if foundUser == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	util.SendData(w, foundUser, http.StatusOK)
}
