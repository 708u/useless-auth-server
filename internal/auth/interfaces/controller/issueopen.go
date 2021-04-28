package controller

import (
	"fmt"
	"net/http"

	"github.com/708u/useless-auth-server/internal/auth/domain/usecase"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter/json"
	"github.com/gorilla/schema"
)

type IssueToken struct {
	UseCase  usecase.IssueTokenUseCase
	Renderer presenter.Renderer
}

type IssueActionRequest struct {
	Code        string `schema:"code"`
	GrantType   string `schema:"grant_type"`
	RedirectURI string `schema:"redirect_uri"`
}

func (i *IssueToken) Action(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var req IssueActionRequest
	if err := schema.NewDecoder().Decode(&req, r.PostForm); err != nil {
		// TODO: fix
		fmt.Printf("failed issue token: %s", err.Error())
		return
	}
	cID, cSecret, ok := r.BasicAuth()
	if !ok {
		// TODO: fix
		fmt.Printf("failed to parse basic auth")
		return
	}

	in := usecase.IssueTokenInput{
		ClientID:     cID,
		ClientSecret: cSecret,
		Code:         req.Code,
		GrantType:    req.GrantType,
		RedirectURI:  req.RedirectURI,
	}
	out, err := i.UseCase.Handle(in)
	if err != nil {
		fmt.Printf("failed issue token: %s", err.Error())
		return
	}
	i.Renderer.Set(json.NewRenderHandler(w, r, out, 200)).Render()
}
