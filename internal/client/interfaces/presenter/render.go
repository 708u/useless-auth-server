package presenter

import "fmt"

// Renderer represents output boundary
type RenderHandler interface {
	Render() error
}

type Renderer struct {
	RenderHandler RenderHandler
}

func NewRenderer(r RenderHandler) *Renderer {
	return &Renderer{RenderHandler: r}
}

func (r *Renderer) Render() error {
	err := r.RenderHandler.Render()
	if err != nil {
		return fmt.Errorf("render failed: %w", err)
	}
	return nil
}
