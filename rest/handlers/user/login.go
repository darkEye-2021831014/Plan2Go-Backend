package user

import (
	"encoding/json"
	"net/http"
	"plan2go-backend/config"
	"plan2go-backend/util"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.SendData(w, map[string]interface{}{
			"success": false,
			"error":   "Invalid request",
		}, http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		util.SendData(w, map[string]interface{}{
			"success": false,
			"error":   "Database error",
		}, http.StatusInternalServerError)
		return
	}

	if user == nil {
		util.SendData(w, map[string]interface{}{
			"success": false,
			"error":   "User not found",
		}, http.StatusUnauthorized)
		return
	}

	// Check if user is verified
	if !user.IsVerified {
		// Delete unverified user
		if err := h.userRepo.DeleteUserByEmail(user.Email); err != nil {
			util.SendData(w, map[string]interface{}{
				"success": false,
				"error":   "Failed to delete unverified user",
			}, http.StatusInternalServerError)
			return
		}

		util.SendData(w, map[string]interface{}{
			"success": false,
			"error":   "Email not verified. Your registration has been removed. Please register again.",
		}, http.StatusUnauthorized)
		return
	}

	// Check password
	if !util.CheckPasswordHash(req.Password, user.Password) {
		util.SendData(w, map[string]interface{}{
			"success": false,
			"error":   "Wrong password",
		}, http.StatusUnauthorized)
		return
	}

	cnf := config.GetConfig()
	token, _ := util.GenerateToken(cnf.Jwt_SecretKey, req.Email)

	util.SendData(w, map[string]interface{}{
		"success": true,
		"token":   token,
		"user":    user,
	}, http.StatusOK)
}
