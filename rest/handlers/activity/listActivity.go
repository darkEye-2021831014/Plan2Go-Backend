package activity

import (
	"encoding/json"
	"net/http"
)

type ListActivityRequest struct {
	UserID int `json:"user_id"`
}

func (h *ActivityHandler) ListActivity(w http.ResponseWriter, r *http.Request) {
	var req ListActivityRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	activities, err := h.Service.QueryActivity(req.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activities)
}
