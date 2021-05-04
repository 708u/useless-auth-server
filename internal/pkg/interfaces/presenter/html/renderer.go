package html

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter"
)

const templateHTML = ".html"

type Layout struct {
	Header Header
	Output interface{}
}

type Header struct {
	Title string
}

// RenderOutputOptions represents New Renderer options func
type RenderOutputOptions func(*HTMLRenderHandler)

// Renderer represents html rendering
type HTMLRenderHandler struct {
	writer   http.ResponseWriter
	template *template.Template
	route    string
	Output   interface{}
}

func NewLayout(h Header, o interface{}) Layout {
	return Layout{Header: h, Output: o}
}

func NewHeader(t string) Header {
	return Header{Title: t}
}

// NewRenderHandler returns Renderer. optionally you can pass output data.
func NewRenderHandler(w http.ResponseWriter, t *template.Template, r string, opts ...RenderOutputOptions) presenter.RenderHandler {
	h := &HTMLRenderHandler{
		writer:   w,
		route:    r,
		template: t,
	}
	for _, f := range opts {
		f(h)
	}
	return h
}

// Handle generate http template.
func (r *HTMLRenderHandler) Handle() error {
	if err := r.render(); err != nil {
		return fmt.Errorf("html render failed: %w", err)
	}
	return nil
}

func (r *HTMLRenderHandler) render() error {
	layout := NewLayout(NewHeader(r.route), r.Output)
	if err := r.template.ExecuteTemplate(r.writer, r.route+templateHTML, layout); err != nil {
		http.Error(r.writer, fmt.Sprintf("failed to exec template: %v", err.Error()), http.StatusInternalServerError)
	}

	return nil
}

// WithOutput set optional output data.
func WithOutput(o interface{}) RenderOutputOptions {
	return func(r *HTMLRenderHandler) {
		r.Output = o
	}
}
