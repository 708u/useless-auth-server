package action

import (
	"encoding/json"
	"net/http"
)

type HealthCheckResponse struct {
	Result string `json:"result,omitempty"`
}

// HealthCheckActions Returns http 200 status for health check
func HealthCheckAction(w http.ResponseWriter, r *http.Request) {
	result := HealthCheckResponse{Result: "OK"}
	res, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
