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

// Render renders output.
// it uses its own render handler to render the output
func (r *Renderer) Render() error {
	err := r.RenderHandler.Render()
	if err != nil {
		return fmt.Errorf("render failed: %w", err)
	}
	return nil
}
