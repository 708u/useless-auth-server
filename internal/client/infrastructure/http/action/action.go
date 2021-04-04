package action

import "net/http"

// Action has all actions.
type Action struct {
	HealthCheck http.HandlerFunc
}
