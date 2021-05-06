package controller

import (
	"net/http"

	"github.com/708u/useless-auth-server/internal/client/domain/usecase"
	clientHTML "github.com/708u/useless-auth-server/internal/client/interfaces/presenter/html"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter/html"
)

type FetchResource struct {
	UseCase  usecase.FetchResourceUseCase
	Renderer presenter.Renderer
}

func NewFetchResource(u usecase.FetchResourceUseCase, r presenter.Renderer) *FetchResource {
	return &FetchResource{UseCase: u, Renderer: r}
}

func (f *FetchResource) Action(w http.ResponseWriter, r *http.Request) {
	// TODO: hard cord
	out, err := f.UseCase.Handle(usecase.FetchResourceInput{AccessToken: "access-token"})
	if err != nil {
		_ = f.Renderer.Set(html.NewRenderHandler(w, clientHTML.T, clientHTML.FetchResourceIndex)).Render()
		return
	}
	// TODO: error handling
	_ = f.Renderer.Set(html.NewRenderHandler(w, clientHTML.T, clientHTML.FetchResourceIndex, html.WithOutput(out))).Render()
}
