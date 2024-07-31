package handler

import (
	"encoding/json"
	"net/http"
)

type HealthHandler struct{}

func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "UP"})
}
