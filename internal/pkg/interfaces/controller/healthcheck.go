package controller

import (
	"encoding/json"
	"net/http"
)

type HealthCheck struct {
}

type HealthCheckResponse struct {
	Result string `json:"result,omitempty"`
}

// NewHealthCheck returns HealthCheck
func NewHealthCheck() *HealthCheck {
	return &HealthCheck{}
}

// Action returns http 200 status for health check
func (h *HealthCheck) Action(w http.ResponseWriter, r *http.Request) {
	result := HealthCheckResponse{Result: "OK"}
	res, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
