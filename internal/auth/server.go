package auth

import (
	"net/http"

	"github.com/708u/useless-auth-server/internal/auth/config"
)

type Server struct {
	Router http.Handler
	Config config.Config
}
