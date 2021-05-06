package usecase

import "errors"

type GetUserResourceUseCase interface {
	Handle(in GetUserResourceInput) (GetUserResourceOutput, error)
}

type GetUserResourceInteractor struct {
}

type GetUserResourceInput struct {
	AccessToken string
}

type GetUserResourceOutput struct {
	Name        string
	Description string
}

func (g *GetUserResourceInteractor) Handle(in GetUserResourceInput) (GetUserResourceOutput, error) {
	// TODO: hard code
	if in.AccessToken != "access-token" {
		return GetUserResourceOutput{}, errors.New("failed GetUserResourceUseCase.handle: access token is invalid")
	}
	// TODO: hard code
	return GetUserResourceOutput{
		Name:        "foo",
		Description: "this is it",
	}, nil
}
