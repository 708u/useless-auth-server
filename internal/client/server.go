package client

import (
	"net/http"

	"github.com/708u/useless-auth-server/internal/client/config"
)

type Server struct {
	Router http.Handler
	Config config.Config
}
