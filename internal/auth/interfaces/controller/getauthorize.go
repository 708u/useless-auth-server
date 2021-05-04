package controller

import (
	"fmt"
	"net/http"

	"github.com/708u/useless-auth-server/internal/auth/domain/usecase"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter/redirect"
	"github.com/gorilla/schema"
)

type GetAuthorize struct {
	UseCase  usecase.GetAuthorizeUseCase
	Renderer presenter.Renderer

	AppURL string
}

type GetAuthorizeRequest struct {
	ClientID     string `schema:"client_id"`
	RedirectURI  string `schema:"redirect_uri"`
	ResponseType string `schema:"response_type"`
}

func (g *GetAuthorize) Action(w http.ResponseWriter, r *http.Request) {
	var req GetAuthorizeRequest
	decoder := schema.NewDecoder()
	if err := decoder.Decode(&req, r.URL.Query()); err != nil {
		// TODO: error handling
		return
	}

	in := usecase.GetAuthorizeInput{
		ClientID:     req.ClientID,
		RedirectURI:  req.RedirectURI,
		ResponseType: req.ResponseType,
	}
	out, err := g.UseCase.Handle(in)
	if err != nil {
		// TODO: error handling
		fmt.Println(out.URI)
	}

	// TODO: fix
	_ = g.Renderer.Set(redirect.NewRenderHandler(w, r, out.URI)).Render()
}
