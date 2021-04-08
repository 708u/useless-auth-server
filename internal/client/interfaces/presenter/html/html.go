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
	Index = "index"
)

//go:embed template/*
var f embed.FS

// template
var t = template.Must(template.ParseFS(
	f,
	templatePath+"*.html",
))

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
	template string
	Output   interface{}
}

func NewLayout(h Header, o interface{}) Layout {
	return Layout{Header: h, Output: o}
}

func NewHeader(t string) Header {
	return Header{Title: t}
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

// Handle generate http template.
func (r *HTMLRenderHandler) Handle() error {
	if err := r.render(); err != nil {
		return fmt.Errorf("html render failed: %w", err)
	}
	return nil
}

func (r *HTMLRenderHandler) render() error {
	layout := NewLayout(NewHeader(r.template), r.Output)
	if err := t.ExecuteTemplate(r.writer, r.template+templateHTML, layout); err != nil {
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
