package gateway

import (
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

	// api exec
	resp, err := http.Get(u.String())
	if err != nil {
		return "", fmt.Errorf("failed AuthorizeGateway.GetAuthorizePage: %w", err)
	}
	if resp.StatusCode >= http.StatusBadRequest {
		return "", fmt.Errorf("failed AuthorizeGateway.GetAuthorizePage: %s", resp.Status)
	}
	return "", nil
}
