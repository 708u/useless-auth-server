package valueobject

import (
	"errors"
	"fmt"
	"net/url"
)

type AuthQuery struct {
	ClientID     string
	ResponseType string
	RedirectURI  string
}

var (
	ErrClientIDEmpty     = errors.New("client id is empty")
	ErrResponseTypeEmpty = errors.New("response type is empty")
)

func NewAuthQuery(cID, responseType, redirectURI string) (AuthQuery, error) {
	if cID == "" {
		return AuthQuery{}, ErrClientIDEmpty
	}
	if responseType == "" {
		return AuthQuery{}, ErrResponseTypeEmpty
	}

	rURI, err := url.QueryUnescape(redirectURI)
	if err != nil {
		return AuthQuery{}, fmt.Errorf("failed to create auth query: %w", err)
	}
	if rURI == "" {
		return AuthQuery{}, fmt.Errorf("decode failed. invalid redirect uri: %s", rURI)
	}

	return AuthQuery{
		ClientID:     cID,
		ResponseType: responseType,
		RedirectURI:  rURI,
	}, nil
}
