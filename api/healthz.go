package api

import (
	"net/http"
)

func (h *APIHandler) HealthzIndex(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "up"}
	JSON(w, http.StatusOK, response)
}
