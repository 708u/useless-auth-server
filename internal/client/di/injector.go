package di

import (
	"net/http"

	"github.com/708u/useless-auth-server/internal/client"
	"github.com/708u/useless-auth-server/internal/client/config"
	infraHTTP "github.com/708u/useless-auth-server/internal/client/infrastructure/http"
	"github.com/708u/useless-auth-server/internal/client/interfaces/controller"
	common "github.com/708u/useless-auth-server/internal/pkg/interfaces/controller"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter"
)

var conf = InjectConfig()

func NewServer() *client.Server {
	return &client.Server{
		Router: InjectRouter(),
		Config: conf,
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
	r := InjectRenderer()

	return &controller.Actions{
		HealthCheck: &common.HealthCheck{},
		ShowIndex:   controller.NewShowIndex(r, conf.Auth.URL),
	}
}

func InjectRenderer() presenter.Renderer {
	return presenter.NewRenderer()
}
