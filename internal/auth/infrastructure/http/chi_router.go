package http

import (
	"fmt"
	"net/http"

	"github.com/708u/useless-auth-server/internal/auth/action"
	"github.com/go-chi/chi/v5"
)

// NewRouter returns chi http handler.
func NewRouter() http.Handler {
	r := chi.NewRouter()

	return routing(r)
}

func routing(r *chi.Mux) chi.Router {
	// health check
	r.Get("/health", action.HealthCheckAction)

	r.Route("/v1", v1Route)
	return r
}

func v1Route(r chi.Router) {
	r.Get("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "/foo. Hello World from Go.")
	})
	r.Get("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "/bar. Hello World from Golang.")
	})
}
