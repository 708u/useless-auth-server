package repository

type ResourceRepository interface {
	FetchUserResource(accessToken string) (string, error)
}
