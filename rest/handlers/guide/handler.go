package guide

import "plan2go-backend/repo"

type GuideHandler struct {
	guideRepo repo.GuideRepo
}

func NewGuideHandler(guideRepo repo.GuideRepo) *GuideHandler {
	return &GuideHandler{guideRepo: guideRepo}
}
