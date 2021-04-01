package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NewRouter returns chi http handler.
func NewRouter() http.Handler {
	r := chi.NewRouter()

	return routing(r)
}

type Handler func(w http.ResponseWriter, r *http.Request) error
type Response struct {
	Result string `json:"result,omitempty"`
	Status int    `json:"status,omitempty"`
}

func HealthCheckAction(w http.ResponseWriter, r *http.Request) {
	result := Response{Result: "OK", Status: http.StatusOK}
	res, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func routing(r *chi.Mux) chi.Router {
	// health check
	r.Get("/health", HealthCheckAction)

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
