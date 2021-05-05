package controller

import (
	"net/http"

	clientHTML "github.com/708u/useless-auth-server/internal/client/interfaces/presenter/html"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter/html"
)

type FetchResource struct {
	Renderer presenter.Renderer
}

func NewFetchResource(r presenter.Renderer) *FetchResource {
	return &FetchResource{Renderer: r}
}

func (f *FetchResource) Action(w http.ResponseWriter, r *http.Request) {
	// TODO: error handling
	_ = f.Renderer.Set(html.NewRenderHandler(w, clientHTML.T, clientHTML.FetchResourceIndex)).Render()
}
