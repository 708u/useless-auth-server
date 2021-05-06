package repository

import "github.com/708u/useless-auth-server/internal/client/interfaces/gateway"

type ResourceRepository interface {
	FetchUserResource(accessToken string) (gateway.TmpResource, error)
}
