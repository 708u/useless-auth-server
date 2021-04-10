package controller_test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/708u/useless-auth-server/internal/client/interfaces/controller"
	"github.com/708u/useless-auth-server/internal/client/interfaces/presenter/html"
	"github.com/708u/useless-auth-server/internal/pkg/interfaces/presenter"
)

func TestNewShowIndex(t *testing.T) {
	t.Parallel()

	type args struct {
		r   presenter.Renderer
		url string
	}
	tests := []struct {
		name string
		args args
		want *controller.ShowIndex
	}{
		{
			name: "success new show index",
			args: args{
				r:   &presenter.Render{},
				url: "http://localhost",
			},
			want: &controller.ShowIndex{Renderer: &presenter.Render{}, AuthURL: "http://localhost"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := controller.NewShowIndex(tt.args.r, "http://localhost"); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewShowIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShowIndex_Action(t *testing.T) {
	t.Parallel()

	type fields struct {
		Renderer presenter.Renderer
	}
	tests := []struct {
		name     string
		fields   fields
		wantCode int
	}{
		{
			name:     "success show index action",
			wantCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			got := httptest.NewRecorder()

			r := presenter.NewRenderer()
			r.Set(html.NewRenderHandler(got, html.Index))

			s := &controller.ShowIndex{
				Renderer: r,
			}
			s.Action(got, req)

			if tt.wantCode != got.Code {
				t.Errorf("Action() failed. wantCode: %d, got: %d", tt.wantCode, got.Code)
			}
		})
	}
}
