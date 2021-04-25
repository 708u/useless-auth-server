package usecase

type GetTokenUseCase interface {
	Handle(in GetTokenInput) (GetTokenOutput, error)
}

type GetTokenInteractor struct{}

type GetTokenInput struct {
	AuthorizationCode string
}
type GetTokenOutput struct{}

func (g *GetTokenInteractor) Handle(in GetTokenInput) (GetTokenOutput, error) {
	// TODO: implement post token request
	return GetTokenOutput{}, nil
}
