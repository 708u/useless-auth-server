package repository

type AuthorizeRepository interface {
	GetAuthorizePage(oURI, cID, rt, rURI string) (string, error)
	GetAccessToken(oURI, code, rURI string) (string, error)
}
