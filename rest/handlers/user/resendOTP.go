package user

import (
	"net/http"
	"plan2go-backend/config"
	"plan2go-backend/util"
)

func (h *Handler) ResendOTP(w http.ResponseWriter, r *http.Request) {
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

	otp, err := h.emailRepo.FetchOTP(email)
	if err != nil {
		util.SendData(w, map[string]interface{}{
			"success": false,
			"error":   "No OTP found for this email",
		}, http.StatusNotFound)
		return
	}

	err = util.SendOTPEmail(email, otp)
	if err != nil {
		util.SendData(w, map[string]interface{}{
			"success": false,
			"error":   "Failed to send OTP email",
		}, http.StatusInternalServerError)
		return
	}

	util.SendData(w, map[string]interface{}{
		"success": true,
		"message": "OTP sent successfully",
	}, http.StatusOK)
}
