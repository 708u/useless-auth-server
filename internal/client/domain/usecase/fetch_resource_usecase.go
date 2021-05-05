package usecase

import "github.com/708u/useless-auth-server/internal/client/domain/repository"

type FetchResourceUseCase interface {
	Handle(input FetchResourceInput) (FetchResourceOutput, error)
}

type FetchResourceInput struct {
	AccessToken string
}

type FetchResourceOutput struct {
	URI string
}

type FetchResourceInteractor struct {
	ResourceRepo repository.ResourceRepository
}

func (f *FetchResourceInteractor) Handle(in FetchResourceInput) (FetchResourceOutput, error) {
	_, err := f.ResourceRepo.FetchUserResource(in.AccessToken)
	if err != nil {
		return FetchResourceOutput{}, err
	}

	return FetchResourceOutput{}, nil
}
