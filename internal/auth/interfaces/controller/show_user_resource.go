package controller

import (
	"fmt"
	"net/http"

	"github.com/708u/useless-auth-server/internal/auth/domain/usecase"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/controller"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter"
	jsonRender "github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter/json"
)

type ShowUserResource struct {
	UseCase  usecase.GetUserResourceUseCase
	Renderer presenter.Renderer
}

type ShowUserResourceResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (s *ShowUserResource) Action(w http.ResponseWriter, r *http.Request) {
	token, err := controller.ParseAccessTokenFromHeader(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("parsing token failed: %s", err.Error()), http.StatusUnauthorized)
	}

	out, err := s.UseCase.Handle(usecase.GetUserResourceInput{AccessToken: token})
	if err != nil {
		http.Error(w, "invalid access token", http.StatusUnauthorized)
	}

	resp := ShowUserResourceResponse{
		Name:        out.Name,
		Description: out.Description,
	}
	// TODO: temporal render
	s.Renderer.Set(jsonRender.NewRenderHandler(w, r, resp, 200)).Render()
}
