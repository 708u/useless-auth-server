package html

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter"
)

const (
	templatePath = "template/"
	templateHTML = ".html"
)

const (
	PathIndex = "index"

	pathHeader = "_header"
	pathFooter = "_footer"
)

//go:embed template/*
var f embed.FS

// RenderOutputOptions represents New Renderer options func
type RenderOutputOptions func(*HTMLRenderHandler)

// Renderer represents html rendering
type HTMLRenderHandler struct {
	writer   http.ResponseWriter
	template string
	Output   interface{}
}

// NewRenderHandler returns Renderer. optionally you can pass output data.
func NewRenderHandler(w http.ResponseWriter, t string, opts ...RenderOutputOptions) presenter.RenderHandler {
	r := &HTMLRenderHandler{
		writer:   w,
		template: t,
	}
	for _, f := range opts {
		f(r)
	}
	return r
}

// Render generate http template.
func (r *HTMLRenderHandler) Render() error {
	if err := r.render(); err != nil {
		return fmt.Errorf("html render failed: %w", err)
	}
	return nil
}

func (r *HTMLRenderHandler) render() error {
	t, err := template.ParseFS(
		f,
		templatePath+r.template+templateHTML,
		templatePath+pathHeader+templateHTML,
		templatePath+pathFooter+templateHTML,
	)
	if err != nil {
		return fmt.Errorf("failed to load template: %w", err)
	}
	if err := t.Execute(r.writer, r.Output); err != nil {
		return fmt.Errorf("failed to exec template: %w", err)
	}
	return nil
}

// WithOutput set optional output data.
func WithOutput(o interface{}) RenderOutputOptions {
	return func(r *HTMLRenderHandler) {
		r.Output = o
	}
}
