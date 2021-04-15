package json

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter"
)

type JSONRenderHandler struct {
	writer     http.ResponseWriter
	request    *http.Request
	Output     interface{}
	StatusCode int
}

// RenderOutputOptions represents New Renderer options func
type RenderOutputOptions func(*JSONRenderHandler)

// NewRenderHandler returns Renderer. optionally you can pass output data.
func NewRenderHandler(w http.ResponseWriter, r *http.Request, o interface{}, s int) presenter.RenderHandler {
	return &JSONRenderHandler{
		writer:     w,
		request:    r,
		Output:     o,
		StatusCode: s,
	}
}

// Handle generate http template.
func (r *JSONRenderHandler) Handle() error {
	if err := r.render(); err != nil {
		return fmt.Errorf("html render failed: %w", err)
	}
	return nil
}

func (r *JSONRenderHandler) render() error {
	resp, err := json.Marshal(r.Output)
	if err != nil {
		http.Error(r.writer, err.Error(), http.StatusInternalServerError)
	}

	r.writer.Header().Set("Content-Type", "application/json")
	r.writer.WriteHeader(r.StatusCode)
	r.writer.Write(resp)

	return nil
}
