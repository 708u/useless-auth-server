package controller

import (
	"net/http"

	"github.com/708u/useless-auth-server/internal/auth/domain/usecase"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter"
)

type GetAuthorize struct {
	UseCase  usecase.GetAuthorizeUseCase
	Renderer presenter.Renderer
}

func (g *GetAuthorize) Action(w http.ResponseWriter, r *http.Request) {
	in := usecase.GetAuthorizeInput{}
	g.UseCase.Handle(in)

	http.Redirect(w, r, "http://localhost:9001", http.StatusFound)
}
