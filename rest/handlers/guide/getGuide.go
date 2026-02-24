package guide

import (
	"encoding/json"
	"net/http"
)

type GuidesRequest struct {
	City string `json:"city"` // optional
}

func (h *GuideHandler) GetGuidesinfo(w http.ResponseWriter, r *http.Request) {
	var req GuidesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	guides, err := h.guideRepo.GetGuides(req.City)
	if err != nil {
		http.Error(w, "Failed to fetch guides: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(guides)
}
