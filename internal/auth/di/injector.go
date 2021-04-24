package di

import (
	"net/http"

	"github.com/708u/useless-auth-server/internal/auth"
	"github.com/708u/useless-auth-server/internal/auth/config"
	"github.com/708u/useless-auth-server/internal/auth/domain/usecase"
	infraHTTP "github.com/708u/useless-auth-server/internal/auth/infrastructure/http"
	"github.com/708u/useless-auth-server/internal/auth/interfaces/controller"
	common "github.com/708u/useless-auth-server/internal/pkg/interfaces/controller"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter"
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

func InjectAction() *controller.Actions {
	r := presenter.NewRenderer()
	usecase := InjectUseCase()

	return &controller.Actions{
		HealthCheck: common.NewHealthCheck(),
		GetAuthorize: &controller.GetAuthorize{
			UseCase:  usecase.GetAuthorize,
			Renderer: r,
		},
	}
}

func InjectRouter() http.Handler {
	return infraHTTP.NewRouter(InjectAction())
}

func InjectUseCase() *usecase.UseCase {
	return &usecase.UseCase{
		GetAuthorize: &usecase.GetAuthorizeInteractor{},
	}
}
