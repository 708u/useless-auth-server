package auth

import (
	"net/http"

	infraHttp "github.com/708u/useless-auth-server/internal/auth/infrastructure/http"
)

type Server struct {
	Router http.Handler
}

func NewServer() (*Server, error) {
	r := infraHttp.NewRouter()
	return &Server{
		Router: r,
	}, nil
}
