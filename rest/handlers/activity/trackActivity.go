package activity

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TrackActivityRequest struct {
	UserID      int    `json:"user_id"`
	Action      string `json:"action"`
	Description string `json:"description"`
	Page        string `json:"page"`
}

func (h *ActivityHandler) TrackActivity(w http.ResponseWriter, r *http.Request) {
	var req TrackActivityRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Println("TrackActivity Request:", req)

	if err := h.Service.TrackActivity(req.UserID, req.Action, req.Description, req.Page); err != nil {
		http.Error(w, "Failed to save activity", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "activity saved"}`))
}
