package usecase

import "errors"

type IssueTokenUseCase interface {
	Handle(input IssueTokenInput) (IssueTokenOutput, error)
}

type IssueTokenInput struct {
	ClientID     string
	ClientSecret string

	Code        string
	GrantType   string
	RedirectURI string
}

type IssueTokenOutput struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type IssueTokenInteractor struct{}

func (i *IssueTokenInteractor) Handle(in IssueTokenInput) (IssueTokenOutput, error) {
	// TODO: hard-coded. change way to auth with fetching id and secret from db
	if in.ClientID != "client_id" || in.ClientSecret != "client_secret" {
		return IssueTokenOutput{}, errors.New("invalid basic auth")
	}
	// TODO: verify auth parameters...

	return IssueTokenOutput{
		AccessToken: "access-token", // TODO: fix hard-coded
		TokenType:   "Bearer",
	}, nil
}
