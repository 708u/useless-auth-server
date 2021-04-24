package gateway

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const (
	authServerAuthorizePath = "/v1/authorize"
)

type AuthorizationGateway struct {
}

func (a *AuthorizationGateway) GetAuthorizePage(oURI, cID, resType, rURI string) (string, error) {
	u, err := url.Parse(oURI)
	if err != nil {
		return "", fmt.Errorf("failed AuthorizeGateway.GetAuthorizePage: %w", err)
	}

	q := u.Query()
	q.Add("client_id", cID)
	q.Add("response_type", resType)
	q.Add("redirect_uri", rURI)
	u.RawQuery = q.Encode()
	u.Path = authServerAuthorizePath

	client := http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	// api exec
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed AuthorizeGateway.GetAuthorizePage: %w", err)
	}
	if resp.StatusCode >= http.StatusBadRequest {
		return "", fmt.Errorf("failed AuthorizeGateway.GetAuthorizePage: %s", resp.Status)
	}
	s, ok := resp.Header["Location"]
	if !ok {
		return "", errors.New("failed to parse location from response")
	}

	return s[0], nil
}
