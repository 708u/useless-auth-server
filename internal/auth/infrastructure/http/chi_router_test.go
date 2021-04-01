package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	infraHTTP "github.com/708u/useless-auth-server/internal/auth/infrastructure/http"
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
			wantBody: "OK\n",
		},
	}
	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/health", nil)
			got := httptest.NewRecorder()
			infraHTTP.HealthCheckAction(got, req)

			if tt.wantCode != got.Result().StatusCode {
				t.Errorf("wantCode: %d, got: %d", tt.wantCode, got.Code)
			}

			a := tt.wantBody
			b := got.Body.String()
			if a != b {
				t.Errorf("wantBody: %s, got: %s", tt.wantBody, got.Body.String())
			}
		})
	}
}
