package controller

import (
	"net/http"

	"github.com/708u/useless-auth-server/internal/client/interfaces/presenter/html"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter"
)

type ShowIndex struct {
	AuthURL  string
	Renderer presenter.Renderer
}

// NewShowIndex returns ShowIndex pointerJ
func NewShowIndex(r presenter.Renderer, url string) *ShowIndex {
	return &ShowIndex{
		Renderer: r,
		AuthURL:  url,
	}
}

// Action shows index
func (s *ShowIndex) Action(w http.ResponseWriter, r *http.Request) {
	s.Renderer.Set(html.NewRenderHandler(w, html.Index, html.WithOutput(s.AuthURL)))
	s.Renderer.Render()
}
