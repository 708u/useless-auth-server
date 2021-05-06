package di

import (
	"net/http"

	"github.com/708u/useless-auth-server/internal/client"
	"github.com/708u/useless-auth-server/internal/client/config"
	"github.com/708u/useless-auth-server/internal/client/domain/repository"
	"github.com/708u/useless-auth-server/internal/client/domain/usecase"
	infraHTTP "github.com/708u/useless-auth-server/internal/client/infrastructure/http"
	"github.com/708u/useless-auth-server/internal/client/interfaces/controller"
	"github.com/708u/useless-auth-server/internal/client/interfaces/gateway"
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
	usecase := InjectUseCase()

	return &controller.Actions{
		HealthCheck: &common.HealthCheck{},

		CallbackGetToken: controller.NewCallbackGetToken(
			conf.Auth.URL,
			conf.Client.RedirectURI,
			usecase.GetToken,
			r,
		),
		GetAuthorize: controller.NewGetAuthorize(
			usecase.GetAuthorize,
			r,
			conf.Auth.URL,
			conf.Client.ID,
			conf.Client.RedirectURI,
			conf.Auth.ResponseType,
		),
		ShowIndex: controller.NewShowIndex(r, conf.Auth.URL),

		FetchResource: controller.NewFetchResource(usecase.FetchResource, r),
	}
}

func InjectRenderer() presenter.Renderer {
	return presenter.NewRenderer()
}

func InjectUseCase() *usecase.UseCase {
	r := InjectRepository()
	return &usecase.UseCase{
		GetAuthorize: &usecase.GetAuthorizeInteractor{AuthorizeRepo: r.AuthorizeRepo},
		GetToken:     &usecase.GetTokenInteractor{AuthorizeRepo: r.AuthorizeRepo},
		FetchResource: &usecase.FetchResourceInteractor{
			ResourceRepo: r.ResourceRepo,
		},
	}
}

func InjectRepository() *repository.Repo {
	return &repository.Repo{
		AuthorizeRepo: &gateway.AuthorizationGateway{},
		ResourceRepo: &gateway.ResourceGateway{
			ResourceSrvURL: conf.Auth.URL,
		},
	}
}
