package usecase

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
	return IssueTokenOutput{}, nil
}
