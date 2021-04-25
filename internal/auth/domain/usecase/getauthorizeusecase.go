package usecase

import (
	"fmt"

	"github.com/708u/useless-auth-server/internal/auth/domain/model/valueobject"
	"github.com/708u/useless-auth-server/internal/auth/domain/service"
)

type GetAuthorizeUseCase interface {
	Handle(input GetAuthorizeInput) (GetAuthorizeOutput, error)
}

type GetAuthorizeInput struct {
	ClientID     string
	RedirectURI  string
	ResponseType string
}

type GetAuthorizeOutput struct {
	URI string
}

type GetAuthorizeInteractor struct {
	URLService service.URLServicer
}

func (g *GetAuthorizeInteractor) Handle(in GetAuthorizeInput) (GetAuthorizeOutput, error) {
	authQuery, err := valueobject.NewAuthQuery(in.ClientID, in.ResponseType, in.RedirectURI)
	if err != nil {
		return GetAuthorizeOutput{}, fmt.Errorf("failed GetAuthorizeUseCase.Handle: %w", err)
	}

	return GetAuthorizeOutput{
		URI: g.URLService.GenAuthorizeURI(authQuery),
	}, nil
}
