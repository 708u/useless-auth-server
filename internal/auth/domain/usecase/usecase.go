package usecase

type UseCase struct {
	GetAuthorize GetAuthorizeUseCase
	IssueToken   IssueTokenUseCase
}
