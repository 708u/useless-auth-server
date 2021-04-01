package auth

import (
	"net/http"

	infraHTTP "github.com/708u/useless-auth-server/internal/auth/infrastructure/http"
)

type Server struct {
	Router http.Handler
}

func NewServer() (*Server, error) {
	r := infraHTTP.NewRouter()
	return &Server{
		Router: r,
	}, nil
}
