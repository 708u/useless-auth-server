package controller

import (
	"fmt"
	"net/http"

	"github.com/708u/useless-auth-server/internal/auth/domain/usecase"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter/json"
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

type IssueActionResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func (i *IssueToken) Action(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var req IssueActionRequest
	if err := presenter.Decoder.Decode(&req, r.PostForm); err != nil {
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

	out, err := i.UseCase.Handle(
		usecase.IssueTokenInput{
			ClientID:     cID,
			ClientSecret: cSecret,
			Code:         req.Code,
			GrantType:    req.GrantType,
			RedirectURI:  req.RedirectURI,
		})

	if err != nil {
		fmt.Printf("failed issue token: %s", err.Error())
		return
	}

	resp := IssueActionResponse{
		AccessToken: out.AccessToken,
		TokenType:   out.TokenType,
	}

	i.Renderer.Set(json.NewRenderHandler(w, r, resp, 200)).Render()
}