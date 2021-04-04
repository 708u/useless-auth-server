package action

import "net/http"

// Action have all actions.
type Action struct {
	HealthCheck http.HandlerFunc
}
