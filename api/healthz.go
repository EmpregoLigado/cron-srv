package api

import (
	"encoding/json"
	"net/http"
)

func (h *APIHandler) HealthzIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{"status": "up"}
	json.NewEncoder(w).Encode(response)
}
