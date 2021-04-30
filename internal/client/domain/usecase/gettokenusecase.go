package usecase

import (
	"fmt"

	"github.com/708u/useless-auth-server/internal/client/domain/repository"
)

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
type GetTokenOutput struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func (g *GetTokenInteractor) Handle(in GetTokenInput) (GetTokenOutput, error) {
	accessToken, err := g.AuthorizeRepo.GetAccessToken(in.AuthServerURI, in.AuthorizationCode, in.RedirectURI)
	if err != nil {
		return GetTokenOutput{}, fmt.Errorf("failed GetTokenUseCase.Handle: %w", err)
	}

	return GetTokenOutput{
		AccessToken: accessToken.Value,
		TokenType:   accessToken.TokenType,
	}, nil
}
