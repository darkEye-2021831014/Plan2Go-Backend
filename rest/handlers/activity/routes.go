package activity

import (
	"net/http"
	"plan2go-backend/rest/middleware"
)

func (h *ActivityHandler) RegisterActivityRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"POST /users/activity/track",
		manager.With(
			http.HandlerFunc(h.TrackActivity),
		),
	)
	mux.Handle(
		"POST /users/activity",
		manager.With(
			http.HandlerFunc(h.ListActivity),
		),
	)

	http.HandleFunc("/v1/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
}
