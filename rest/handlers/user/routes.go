package user

import (
	"net/http"
	"plan2go-backend/rest/middleware"
)

// Helper to check HTTP method
func HandleMethod(handlerFunc http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		handlerFunc(w, r)
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	// Create user (POST /users)
	mux.Handle("/users", manager.With(
		HandleMethod(h.CreateUser, http.MethodPost),
	))

	// Verify OTP (POST /users/verify)
	mux.Handle("/users/verify", manager.With(
		HandleMethod(h.VerifyOTP, http.MethodPost),
	))

	// Resend OTP (GET /users/resend-otp)
	mux.Handle("/users/resend-otp", manager.With(
		HandleMethod(h.ResendOTP, http.MethodGet),
	))

	// Login (POST /users/login)
	mux.Handle("/users/login", manager.With(
		HandleMethod(h.Login, http.MethodPost),
	))

	// Get profile (GET /users/profile)
	mux.Handle("/users/profile", manager.With(
		HandleMethod(h.GetUserByEmail, http.MethodGet),
	))

	// Change password (POST /users/update/password)
	mux.Handle("/users/update/password", manager.With(
		HandleMethod(h.ChangePassword, http.MethodPost),
	))

	// Update profile (POST /users/update/profile)
	mux.Handle("/users/update/profile", manager.With(
		HandleMethod(h.UpdateProfile, http.MethodPost),
	))
}
