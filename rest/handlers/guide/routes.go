package guide

import (
	"net/http"
	"plan2go-backend/rest/middleware"
)

func (h *GuideHandler) GuideRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle(
		"POST /users/guide",
		manager.With(
			http.HandlerFunc(h.GetGuidesinfo),
		),
	)
	mux.Handle(
		"POST /users/guide/register",
		manager.With(
			http.HandlerFunc(h.CreateGuide),
		),
	)
}
