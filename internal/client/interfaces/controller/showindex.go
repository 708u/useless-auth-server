package controller

import (
	"net/http"

	template "github.com/708u/useless-auth-server/internal/client/interfaces/presenter/html"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter/html"
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
	// TODO: fix
	_ = s.Renderer.Set(html.NewRenderHandler(w, template.T, template.Index, html.WithOutput(s.AuthURL))).Render()
}
