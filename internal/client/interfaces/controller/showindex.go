package controller

import (
	"net/http"

	"github.com/708u/useless-auth-server/internal/client/interfaces/presenter/html"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter"
)

type ShowIndex struct {
	Renderer presenter.Renderer
}

// NewShowIndex returns ShowIndex pointerJ
func NewShowIndex(r presenter.Renderer) *ShowIndex {
	return &ShowIndex{Renderer: r}
}

// Action shows index
func (s *ShowIndex) Action(w http.ResponseWriter, r *http.Request) {
	s.Renderer.Set(html.NewRenderHandler(w, html.Index))
	s.Renderer.Render()
}
