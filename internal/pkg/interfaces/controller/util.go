package controller

import (
	"errors"
	"net/http"
	"strings"
)

// ParseAccessTokenFromHeader returns access token.
// it suppose to be given http header like "Authorization: Bearer xxxxxx"
func ParseAccessTokenFromHeader(r *http.Request) (string, error) {
	h := strings.Split(r.Header.Get("Authorization"), " ")
	if len(h) != 2 {
		return "", errors.New("invalid access token")
	}

	tokenTyp := h[0]
	if tokenTyp != "Bearer" {
		return "", errors.New(`authorization header format should be "Authorization Bearer: token"`)
	}

	return h[1], nil
}
