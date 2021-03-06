package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/708u/useless-auth-server/internal/pkg/interfaces/controller"
)

func TestNewRouter(t *testing.T) {
	tests := []struct {
		name     string
		wantCode int
		wantBody string
	}{
		{
			name:     "success health check",
			wantCode: http.StatusOK,
			wantBody: `{"result":"OK"}`,
		},
	}
	t.Parallel()
	healthCheck := controller.NewHealthCheck()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/health", nil)
			got := httptest.NewRecorder()
			healthCheck.Action(got, req)

			if tt.wantCode != got.Result().StatusCode {
				t.Errorf("wantCode: %d, got: %d", tt.wantCode, got.Code)
			}

			if tt.wantBody != got.Body.String() {
				t.Errorf("wantBody: %s, got: %s", tt.wantBody, got.Body.String())
			}
		})
	}
}
