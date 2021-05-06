package usecase

type UseCase struct {
	// OAuth2/OIDC
	GetAuthorize GetAuthorizeUsecase
	GetToken     GetTokenUseCase

	// resource
	FetchResource FetchResourceUseCase
}
