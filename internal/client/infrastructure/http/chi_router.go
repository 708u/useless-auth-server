package http

import (
	"net/http"

	"github.com/708u/useless-auth-server/internal/client/interfaces/controller"
	"github.com/go-chi/chi/v5"
)

// NewRouter returns chi http handler.
func NewRouter(a *controller.Actions) http.Handler {
	r := chi.NewRouter()

	return routing(r, a)
}

func routing(r *chi.Mux, a *controller.Actions) chi.Router {
	// health check
	r.Get("/health", a.HealthCheck.Action)

	r.Get("/", a.ShowIndex.Action)

	r.Route("/v1", v1Route(a))
	return r
}

func v1Route(a *controller.Actions) func(chi.Router) {
	return func(r chi.Router) {
		r.Get("/authorize", a.GetAuthorize.Action)

		r.Get("/callback", a.CallbackGetToken.Action)
	}
}
