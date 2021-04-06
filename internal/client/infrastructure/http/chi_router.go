package http

import (
	"fmt"
	"net/http"

	"github.com/708u/useless-auth-server/internal/client/infrastructure/http/action"
	"github.com/708u/useless-auth-server/internal/client/interfaces/presenter"
	"github.com/708u/useless-auth-server/internal/client/interfaces/presenter/html"
	"github.com/go-chi/chi/v5"
)

// NewRouter returns chi http handler.
func NewRouter(a *action.Action) http.Handler {
	r := chi.NewRouter()

	return routing(r, a)
}

func routing(r *chi.Mux, action *action.Action) chi.Router {
	// health check
	r.Get("/health", action.HealthCheck)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		renderer := presenter.NewRenderer(html.NewRenderHandler(w, html.PathIndex))
		if err := renderer.Render(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	r.Route("/v1", v1Route(action))
	return r
}

func v1Route(action *action.Action) func(chi.Router) {
	return func(r chi.Router) {
		r.Get("/foo", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "/foo. Hello World from Go.")
		})
		r.Get("/bar", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "/bar. Hello World from Golang.")
		})
	}
}
