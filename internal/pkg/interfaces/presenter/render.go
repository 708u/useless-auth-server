package presenter

import (
	"errors"
	"fmt"
)

// Renderer represents output boundary
type RenderHandler interface {
	Handle() error
}

type Renderer interface {
	Set(h RenderHandler) Renderer
	Render() error
}

type Render struct {
	RenderHandler RenderHandler
}

// NewRenderer returns renderer
func NewRenderer() Renderer {
	return &Render{}
}

// Set sets render handler.
func (r *Render) Set(h RenderHandler) Renderer {
	r.RenderHandler = h
	return r
}

// Render renders output.
// it uses its own render handler to render the output
func (r *Render) Render() error {
	if r.RenderHandler == nil {
		return errors.New("render handler must be set")
	}

	err := r.RenderHandler.Handle()
	if err != nil {
		return fmt.Errorf("render failed: %w", err)
	}
	return nil
}
