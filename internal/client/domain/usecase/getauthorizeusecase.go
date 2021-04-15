package usecase

import (
	"fmt"
	"net/http"

	"github.com/708u/useless-auth-server/internal/client/domain/repository"
)

type GetAuthorizeUsecase interface {
	Handle(input GetAuthorizeInput) (GetAuthorizeOutput, error)
}

type GetAuthorizeInteractor struct {
	AuthorizeRepo repository.AuthorizeRepository
}

type GetAuthorizeInput struct {
	AuthServiceURI string
	ClientId       string
	ResponseType   string
	RedirectURI    string
}

type GetAuthorizeOutput struct {
	RedirectTo string
	StatusCode int
}

func (g *GetAuthorizeInteractor) Handle(in GetAuthorizeInput) (GetAuthorizeOutput, error) {
	redirect, err := g.AuthorizeRepo.GetAuthorizePage(in.AuthServiceURI, in.ClientId, in.ResponseType, in.RedirectURI)
	if err != nil {
		return GetAuthorizeOutput{}, fmt.Errorf("fail AuthorizeRepo.GetAuthorizePage: %w", err)
	}
	return GetAuthorizeOutput{RedirectTo: redirect, StatusCode: http.StatusFound}, nil
}
