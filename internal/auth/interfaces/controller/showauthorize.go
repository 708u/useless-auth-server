package controller

import (
	"net/http"

	"github.com/708u/useless-auth-server/internal/auth/interfaces/presenter/html"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter"
	commonHTML "github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter/html"
)

type ShowAuthorize struct {
	Renderer presenter.Renderer
}

func (s *ShowAuthorize) Action(w http.ResponseWriter, r *http.Request) {
	// TODO: fix
	_ = s.Renderer.Set(commonHTML.NewRenderHandler(w, html.T, html.Authorize)).Render()
}
