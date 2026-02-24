package user

import (
	"net/http"
	"plan2go-backend/rest/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(h.CreateUser),
		),
	)
	mux.Handle(
		"POST /users/verify",
		manager.With(
			http.HandlerFunc(h.VerifyOTP),
		),
	)
	mux.Handle(
		"GET /users/resend-otp",
		manager.With(
			http.HandlerFunc(h.ResendOTP),
		),
	)
	mux.Handle(
		"POST /users/login",
		manager.With(
			http.HandlerFunc(h.Login),
		),
	)
	mux.Handle(
		"GET /users/profile",
		manager.With(
			http.HandlerFunc(h.GetUserByEmail),
		),
	)
	mux.Handle(
		"POST /users/update/password",
		manager.With(
			http.HandlerFunc(h.ChangePassword),
		),
	)
	mux.Handle(
		"POST /users/update/profile",
		manager.With(
			http.HandlerFunc(h.UpdateProfile),
		),
	)

}
