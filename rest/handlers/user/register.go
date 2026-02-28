package user

import (
	"encoding/json"
	"net/http"
	"plan2go-backend/config"
	"plan2go-backend/repo"
	"plan2go-backend/util"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user repo.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if user.Email == "" || user.Password == "" || user.FirstName == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword
	user.IsVerified = false

	// Create user in DB
	createdUser, err := h.userRepo.CreateUser(user)
	if err != nil {
		http.Error(w, "User already exists or DB error", http.StatusConflict)
		return
	}

	// Generate OTP
	otp := util.GenerateOTP()

	// Save OTP in DB
	err = h.emailRepo.SaveOTP(createdUser.Email, otp)
	if err != nil {
		http.Error(w, "Failed to generate OTP", http.StatusInternalServerError)
		return
	}

	// Send OTP email
	// err = util.SendOTPEmail(createdUser.Email, otp)
	// if err != nil {
	// 	fmt.Println("Warning: failed to send OTP email:", err)
	// }

	token, err := util.GenerateToken(config.GetConfig().Jwt_SecretKey, createdUser.Email)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	util.SendData(w, map[string]interface{}{
		"success":     true,
		"user":        createdUser,
		"token":       token,
		"requiresOTP": true,
		"message":     "User created successfully. OTP sent to email.",
		"otp":         otp,
	}, http.StatusCreated)
}
