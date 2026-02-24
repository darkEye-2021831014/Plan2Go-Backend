package plan

import (
	"encoding/json"
	"net/http"
)

type PlanRequest struct {
	Division     string `json:"division"`
	District     string `json:"district"`
	Budget       string `json:"budget"`
	LocationType string `json:"locationType"`
	Days         int    `json:"days"`
}

func (h *PlanHandler) GeneratePlan(w http.ResponseWriter, r *http.Request) {
	var req PlanRequest

	// Decode JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate all required fields based on struct
	if req.Division == "" || req.District == "" || req.Budget == "" ||
		req.LocationType == "" || req.Days <= 0 {
		http.Error(w, "All fields are required: division, district, budget, locationType, days", http.StatusBadRequest)
		return
	}

	// Call service
	plan, err := h.planService.GenerateTourPlan(
		r.Context(),
		req.Division,
		req.District,
		req.Budget,
		req.LocationType,
		req.Days,
	)
	if err != nil {
		http.Error(w, "Failed to generate plan: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plan)
}
