package repository

type AuthorizeRepository interface {
	GetAuthorizePage(oURI, cID, rt, rURI string) (string, error)
}
