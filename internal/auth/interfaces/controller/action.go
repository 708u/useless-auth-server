package controller

import common "github.com/708u/useless-auth-server/internal/pkg/interfaces/controller"

// Actions have all actions.
type Actions struct {
	HealthCheck common.Action

	// OAuth/OIDC
	GetAuthorize  common.Action
	ShowAuthorize common.Action

	IssueToken common.Action

	// Resource
	ShowUserResource common.Action
}
