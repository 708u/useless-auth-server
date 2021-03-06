package controller

import common "github.com/708u/useless-auth-server/internal/pkg/interfaces/controller"

// Action has all actions.
type Actions struct {
	HealthCheck common.Action

	// OAuth2.0/OIDC
	CallbackGetToken common.Action
	GetAuthorize     common.Action
	ShowIndex        common.Action

	// client app
	FetchResource common.Action
}
