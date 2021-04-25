package service

import (
	"net/url"

	"github.com/708u/useless-auth-server/internal/auth/domain/model/valueobject"
)

type URLService struct {
	AppURL string
}

func NewURLService(url string) URLServicer {
	return &URLService{
		AppURL: url,
	}
}

func (u *URLService) GenAuthorizeURI(authQuery valueobject.AuthQuery) string {
	aURL, _ := url.Parse(u.AppURL)
	aURL.Path = "authorize"

	// set querystrings
	q := url.Values{}
	q.Add("client_id", authQuery.ClientID)
	q.Add("response_type", authQuery.ResponseType)
	q.Add("redirect_uri", authQuery.RedirectURI)
	aURL.RawQuery = q.Encode()

	return aURL.String()
}
