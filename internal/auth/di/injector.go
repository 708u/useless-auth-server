package di

import (
	"net/http"

	"github.com/708u/useless-auth-server/internal/auth"
	"github.com/708u/useless-auth-server/internal/auth/config"
	"github.com/708u/useless-auth-server/internal/auth/domain/service"
	"github.com/708u/useless-auth-server/internal/auth/domain/usecase"
	infraHTTP "github.com/708u/useless-auth-server/internal/auth/infrastructure/http"
	"github.com/708u/useless-auth-server/internal/auth/interfaces/controller"
	common "github.com/708u/useless-auth-server/internal/pkg/interfaces/controller"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter"
)

var conf = injectConfig()

func injectConfig() config.Config {
	return config.NewConfig(
		config.ConfigName, config.ConfigPath, config.ConfigType,
	)
}

func NewServer() *auth.Server {
	return &auth.Server{
		Router: injectRouter(),
		Config: conf,
	}
}

func injectRouter() http.Handler {
	return infraHTTP.NewRouter(injectAction())
}

func injectAction() *controller.Actions {
	u := injectUseCase()

	return &controller.Actions{
		HealthCheck: common.NewHealthCheck(),
		GetAuthorize: &controller.GetAuthorize{
			UseCase:  u.GetAuthorize,
			Renderer: presenter.NewRenderer(),
			AppURL:   conf.HTTP.URL,
		},
	}
}

func injectUseCase() *usecase.UseCase {
	srv := injectService()

	return &usecase.UseCase{
		GetAuthorize: &usecase.GetAuthorizeInteractor{
			URLService: srv.URL,
		},
	}
}

func injectService() *service.Service {
	return &service.Service{
		URL: service.NewURLService(conf.HTTP.URL),
	}
}
