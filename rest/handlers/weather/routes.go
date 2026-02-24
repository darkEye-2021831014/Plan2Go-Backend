package weather

import (
	"net/http"
	"plan2go-backend/rest/middleware"
)

func (h *Handler) WeatherRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle(
		"GET /users/weather",
		manager.With(
			http.HandlerFunc(h.GetWeatherHandler),
		),
	)
}
