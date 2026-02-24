package plan

import "plan2go-backend/rest/services"

type PlanHandler struct {
	planService *services.PlanService
}

func NewPlanHandler(planService *services.PlanService) *PlanHandler {
	return &PlanHandler{planService: planService}
}