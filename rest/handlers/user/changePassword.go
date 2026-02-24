package user

import (
	"encoding/json"
	"net/http"
	"plan2go-backend/config"
	"plan2go-backend/util"
)

type changePasswordReq struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func (h *Handler) ChangePassword(w http.ResponseWriter, r *http.Request) {

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

	var req changePasswordReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.OldPassword == "" || req.NewPassword == "" {
		http.Error(w, "Old and new password required", http.StatusBadRequest)
		return
	}

	storedPassword, err := h.userRepo.GetUserPassword(email)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if storedPassword == "" {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if !util.CheckPasswordHash(req.OldPassword, storedPassword) {
		http.Error(w, "Old password incorrect", http.StatusUnauthorized)
		return
	}

	newHashedPassword, err := util.HashPassword(req.NewPassword)
	if err != nil {
		http.Error(w, "Error hashing new password", http.StatusInternalServerError)
		return
	}

	_, err = h.userRepo.UpdatePassword(email, newHashedPassword)
	if err != nil {
		http.Error(w, "Failed to update password", http.StatusInternalServerError)
		return
	}

	util.SendData(w, map[string]string{
		"message": "Password changed successfully",
	}, http.StatusOK)
}
