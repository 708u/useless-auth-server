package controller

import (
	"fmt"
	"net/http"

	"github.com/708u/useless-auth-server/internal/client/domain/usecase"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter/json"
	"github.com/gorilla/schema"
)

type CallbackGetToken struct {
	UseCase  usecase.GetTokenUseCase
	Renderer presenter.Renderer
}

type CallbackGetTokenRequest struct {
	Code string `schema:"code"`
}

func NewCallbackGetToken(u usecase.GetTokenUseCase, r presenter.Renderer) *CallbackGetToken {
	return &CallbackGetToken{
		UseCase:  u,
		Renderer: r,
	}
}

func (c *CallbackGetToken) Action(w http.ResponseWriter, r *http.Request) {
	var req CallbackGetTokenRequest
	schema.NewDecoder().Decode(&req, r.URL.Query())
	in := usecase.GetTokenInput{AuthorizationCode: req.Code}

	out, err := c.UseCase.Handle(in)
	if err != nil {
		// TODO: add error handling
		fmt.Println(err)
	}
	// TODO: temporal set
	c.Renderer.Set(json.NewRenderHandler(w, r, out, 200))

}
