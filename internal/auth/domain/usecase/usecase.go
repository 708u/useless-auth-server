package usecase

type UseCase struct {
	// OAuth2/OIDC
	GetAuthorize GetAuthorizeUseCase
	IssueToken   IssueTokenUseCase

	// Resource
	GetUserResource GetUserResourceUseCase
}
