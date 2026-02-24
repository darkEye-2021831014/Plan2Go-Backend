package plan

import (
	"net/http"
	"plan2go-backend/rest/middleware"
)

func (h *PlanHandler) PlanRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle(
		"POST /users/plan",
		manager.With(
			http.HandlerFunc(h.GeneratePlan),
		),
	)

}
