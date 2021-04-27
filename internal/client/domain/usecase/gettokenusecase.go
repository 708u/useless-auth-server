package usecase

import "github.com/708u/useless-auth-server/internal/client/domain/repository"

type GetTokenUseCase interface {
	Handle(in GetTokenInput) (GetTokenOutput, error)
}

type GetTokenInteractor struct {
	AuthorizeRepo repository.AuthorizeRepository
}

type GetTokenInput struct {
	AuthServerURI     string
	AuthorizationCode string
	RedirectURI       string
}
type GetTokenOutput struct{}

func (g *GetTokenInteractor) Handle(in GetTokenInput) (GetTokenOutput, error) {
	g.AuthorizeRepo.GetAccessToken(in.AuthServerURI, in.AuthorizationCode, in.RedirectURI)
	// TODO: implement post token request
	return GetTokenOutput{}, nil
}
