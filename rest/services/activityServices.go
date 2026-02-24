package services

import (
	"plan2go-backend/models"
	"plan2go-backend/repo"
)

type ActivityService struct {
	Repo *repo.ActivityRepo
}

func NewActivityService(r *repo.ActivityRepo) *ActivityService {
	return &ActivityService{Repo: r}
}

func (s *ActivityService) TrackActivity(userID int, action, description, page string) error {
	activity := models.Activity{
		UserID:      userID,
		Action:      action,
		Description: description,
		Page:        page,
	}
	return s.Repo.CreateActivity(activity)
}

func (s *ActivityService) QueryActivity(userID int) ([]models.Activity, error) {
	return s.Repo.QueryActivity(userID)
}
