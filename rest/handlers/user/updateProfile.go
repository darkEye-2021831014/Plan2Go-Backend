package user

import (
	"encoding/json"
	"net/http"
	"plan2go-backend/config"
	"plan2go-backend/util"
	"strings"
)

type updateProfileReq struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

func (h *Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	tokenString := authHeader
	if len(authHeader) > 7 && strings.HasPrefix(authHeader, "Bearer ") {
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

	var req updateProfileReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.FirstName == "" && req.LastName == "" && req.Phone == "" {
		http.Error(w, "Nothing to update", http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}

	updatedUser, err := h.userRepo.UpdateUserProfile(user)
	if err != nil {
		http.Error(w, "Failed to update profile", http.StatusInternalServerError)
		return
	}

	util.SendData(w, map[string]interface{}{
		"message": "Profile updated successfully",
		"user":    updatedUser,
	}, http.StatusOK)
}
