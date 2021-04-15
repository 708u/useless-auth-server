package controller

import (
	"fmt"
	"net/http"

	"github.com/708u/useless-auth-server/internal/client/domain/usecase"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter/redirect"
)

type GetAuthorize struct {
	UseCase  usecase.GetAuthorizeUsecase
	Renderer presenter.Renderer

	AuthServerURI string
	ClientID      string
	RedirectURI   string
	ResponseType  string
}

func NewGetAuthorize(u usecase.GetAuthorizeUsecase, r presenter.Renderer, uri, cID, rURI, rt string) *GetAuthorize {
	return &GetAuthorize{
		UseCase:  u,
		Renderer: r,

		AuthServerURI: uri,
		ClientID:      cID,
		RedirectURI:   rURI,
		ResponseType:  rt,
	}
}

func (a *GetAuthorize) Action(w http.ResponseWriter, r *http.Request) {
	in := usecase.GetAuthorizeInput{
		AuthServiceURI: a.AuthServerURI,
		ClientId:       a.ClientID,
		ResponseType:   a.ResponseType,
		RedirectURI:    a.RedirectURI,
	}

	out, err := a.UseCase.Handle(in)
	if err != nil {
		// TODO: implement error renderer and replace it
		http.Error(w, fmt.Sprintf("internal error: %s", err.Error()), http.StatusInternalServerError)
	}

	a.Renderer.Set(redirect.NewRenderHandler(w, r, out.RedirectTo))
	a.Renderer.Render()
}
