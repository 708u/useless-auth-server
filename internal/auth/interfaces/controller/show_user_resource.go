package controller

import (
	"net/http"

	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter/json"
)

type ShowUserResource struct {
	Renderer presenter.Renderer
}

func (s *ShowUserResource) Action(w http.ResponseWriter, r *http.Request) {
	// TODO: temoral render
	s.Renderer.Set(json.NewRenderHandler(w, r, nil, 200)).Render()
}
