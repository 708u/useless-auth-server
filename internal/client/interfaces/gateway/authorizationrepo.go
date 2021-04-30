package gateway

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/708u/useless-auth-server/internal/client/domain/model/valueobject"
)

const (
	authServerAuthorizePath = "/v1/authorize"
	authServerTokenEndpoint = "/v1/token"
)

const (
	authGrantTypeCode = "authorization_code"
)

var emptyAC valueobject.AccessToken

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

func (a *AuthorizationGateway) GetAccessToken(oURI, code, rURI string) (valueobject.AccessToken, error) {
	u, err := url.Parse(oURI)
	if err != nil {
		return emptyAC, fmt.Errorf("failed AuthorizeGateway.GetAccessToken: %w", err)
	}

	form := url.Values{}
	form.Add("grant_type", authGrantTypeCode)
	form.Add("code", code)
	form.Add("redirect_uri", rURI)

	reqBody := strings.NewReader(form.Encode())
	req, _ := http.NewRequest(http.MethodPost, u.String(), reqBody)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.URL.Path = authServerTokenEndpoint
	// TODO: set client id and secret
	req.SetBasicAuth("client_id", "client_secret")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return emptyAC, fmt.Errorf("failed AuthorizeGateway.GetAccessToken: %w", err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return emptyAC, err
	}

	var accessToken valueobject.AccessToken
	if err := json.Unmarshal(respBody, &accessToken); err != nil {
		return emptyAC, fmt.Errorf("failed AuthorizeGateway.GetAccessToken: %w", err)
	}

	return accessToken, nil
}
