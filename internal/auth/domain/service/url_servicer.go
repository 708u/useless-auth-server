package service

import "github.com/708u/useless-auth-server/internal/auth/domain/model/valueobject"

type URLServicer interface {
	GenAuthorizeURI(authQuery valueobject.AuthQuery) string
}
