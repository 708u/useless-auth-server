package di

import (
	"net/http"

	"github.com/708u/useless-auth-server/internal/auth"
	"github.com/708u/useless-auth-server/internal/auth/config"
	infraHTTP "github.com/708u/useless-auth-server/internal/auth/infrastructure/http"
	"github.com/708u/useless-auth-server/internal/auth/interfaces/controller"
	common "github.com/708u/useless-auth-server/internal/pkg/interfaces/controller"
)

func NewServer() *auth.Server {
	return &auth.Server{
		Router: InjectRouter(),
		Config: InjectConfig(),
	}
}

func InjectConfig() config.Config {
	return config.NewConfig(
		config.ConfigName, config.ConfigPath, config.ConfigType,
	)
}

func InjectRouter() http.Handler {
	return infraHTTP.NewRouter(InjectAction())
}

func InjectAction() *controller.Actions {
	return &controller.Actions{
		HealthCheck: common.NewHealthCheck(),
	}
}
