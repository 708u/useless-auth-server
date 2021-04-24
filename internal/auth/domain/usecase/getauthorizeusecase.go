package usecase

type GetAuthorizeUseCase interface {
	Handle(input GetAuthorizeInput) (GetAuthorizeOutput, error)
}

type GetAuthorizeInteractor struct {
}

type GetAuthorizeInput struct {
}

type GetAuthorizeOutput struct {
}

func (g *GetAuthorizeInteractor) Handle(in GetAuthorizeInput) (GetAuthorizeOutput, error) {
	return GetAuthorizeOutput{}, nil
}
