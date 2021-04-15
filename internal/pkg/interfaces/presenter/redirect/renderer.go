package redirect

import (
	"net/http"

	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter"
)

type HTTPRedirectRenderHandler struct {
	writer     http.ResponseWriter
	request    *http.Request
	body       string
	statusCode int
}

// RenderOutputOptions represents New Renderer options func
type RenderOutputOptions func(*HTTPRedirectRenderHandler)

// NewRenderHandler returns redirect renderer. default status code is 302. optionally you can pass status code to change the other.
func NewRenderHandler(w http.ResponseWriter, r *http.Request, b string, opts ...RenderOutputOptions) presenter.RenderHandler {
	h := &HTTPRedirectRenderHandler{
		writer:     w,
		request:    r,
		body:       b,
		statusCode: http.StatusFound,
	}
	for _, f := range opts {
		f(h)
	}
	return h
}

// WithStatusCode returns functional options. we can change the other status code.
func WithStatusCode(c int) RenderOutputOptions {
	return func(h *HTTPRedirectRenderHandler) {
		h.statusCode = c
	}
}

// Handle writes redirect response.
func (h *HTTPRedirectRenderHandler) Handle() error {
	http.Redirect(h.writer, h.request, h.body, h.statusCode)
	return nil
}
