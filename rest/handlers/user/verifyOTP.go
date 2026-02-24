package user

import (
	"encoding/json"
	"net/http"
	"plan2go-backend/util"
)

// Request payload
type VerifyOTPRequest struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

func (h *Handler) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var req VerifyOTPRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.SendData(w, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
		}, http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.OTP == "" {
		util.SendData(w, map[string]interface{}{
			"success": false,
			"error":   "Email and OTP are required",
		}, http.StatusBadRequest)
		return
	}

	valid, err := h.emailRepo.VerifyOTP(req.Email, req.OTP)
	if err != nil || !valid {
		util.SendData(w, map[string]interface{}{
			"success": false,
			"error":   "Invalid OTP",
		}, http.StatusUnauthorized)
		return
	}

	if !valid {
		util.SendData(w, map[string]interface{}{
			"success": false,
			"error":   "Invalid or expired OTP",
		}, http.StatusUnauthorized)
		return
	}

	user, err := h.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		util.SendData(w, map[string]interface{}{
			"success": false,
			"error":   "Database error while fetching user",
		}, http.StatusInternalServerError)
		return
	}
	if user == nil {
		util.SendData(w, map[string]interface{}{
			"success": false,
			"error":   "User not found",
		}, http.StatusNotFound)
		return
	}

	updatedUser, err := h.userRepo.UpdateUserVerification(user.Email, true)
	if err != nil {
		util.SendData(w, map[string]interface{}{
			"success": false,
			"error":   "Failed to update user verification",
		}, http.StatusInternalServerError)
		return
	}

	_ = h.emailRepo.DeleteOTP(req.Email)

	util.SendData(w, map[string]interface{}{
		"success": true,
		"message": "User verified successfully",
		"user":    updatedUser,
	}, http.StatusOK)
}
