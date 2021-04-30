package repository

import "github.com/708u/useless-auth-server/internal/client/domain/model/valueobject"

type AuthorizeRepository interface {
	GetAuthorizePage(oURI, cID, rt, rURI string) (string, error)
	GetAccessToken(oURI, code, rURI string) (valueobject.AccessToken, error)
}
