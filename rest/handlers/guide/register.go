package guide

import (
	"encoding/json"
	"net/http"
	"plan2go-backend/repo"
	"plan2go-backend/util"
)

func (h *GuideHandler) CreateGuide(w http.ResponseWriter, r *http.Request) {
	var guide repo.Guide

	if err := json.NewDecoder(r.Body).Decode(&guide); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if guide.City == "" || guide.HourlyFee <= 0 {
		http.Error(w, "Missing required fields: city, hourly_fee", http.StatusBadRequest)
		return
	}

	createdGuide, err := h.guideRepo.CreateGuide(guide)
	if err != nil {
		http.Error(w, "Guide already exists or DB error: "+err.Error(), http.StatusConflict)
		return
	}

	util.SendData(w, map[string]interface{}{
		"message": "Guide created successfully",
		"guide":   createdGuide,
	}, http.StatusCreated)
}
