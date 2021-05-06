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
	r := presenter.NewRenderer()
	usecase := injectUseCase()

	return &controller.Actions{
		HealthCheck: common.NewHealthCheck(),
		// OAuth2/OIDC
		GetAuthorize: &controller.GetAuthorize{
			UseCase:  usecase.GetAuthorize,
			Renderer: r,
			AppURL:   conf.HTTP.URL,
		},
		ShowAuthorize: &controller.ShowAuthorize{
			Renderer: r,
		},
		IssueToken: &controller.IssueToken{
			UseCase:  usecase.IssueToken,
			Renderer: r,
		},
		// Resource
		ShowUserResource: &controller.ShowUserResource{
			UseCase:  usecase.GetUserResource,
			Renderer: r,
		},
	}
}

func injectUseCase() *usecase.UseCase {
	srv := injectService()

	return &usecase.UseCase{
		// OAuth2/OIDC
		GetAuthorize: &usecase.GetAuthorizeInteractor{
			URLService: srv.URL,
		},
		IssueToken: &usecase.IssueTokenInteractor{},

		// Resource
		GetUserResource: &usecase.GetUserResourceInteractor{},
	}
}

func injectService() *service.Service {
	return &service.Service{
		URL: service.NewURLService(conf.HTTP.URL),
	}
}
