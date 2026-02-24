package activity

import "plan2go-backend/rest/services"

type ActivityHandler struct {
	Service *services.ActivityService
}

func NewActivityHandler(s *services.ActivityService) *ActivityHandler {
	return &ActivityHandler{Service: s}
}
