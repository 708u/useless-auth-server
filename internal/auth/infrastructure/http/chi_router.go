package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	return setRoute(r)
}

func setRoute(r *chi.Mux) chi.Router {
	return r.Route("/v1", v1Route)
}

func v1Route(r chi.Router) {
	r.Get("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "/foo. Hello World from Go.")
	})
	r.Get("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "/bar. Hello World from Golang.")
	})
}
