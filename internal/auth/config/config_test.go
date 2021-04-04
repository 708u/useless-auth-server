package config_test

import (
	"os"
	"testing"

	"github.com/708u/useless-auth-server/internal/auth/config"
	"github.com/google/go-cmp/cmp"
)

func TestConfig_NewConfig(t *testing.T) {
	t.Parallel()
	type args struct {
		name string
		path string
		typ  string
	}
	tests := []struct {
		name        string
		args        args
		want        config.Config
		wantRecover bool
		setEnv      func()
	}{
		{
			name: "success new config",
			args: args{
				name: "test_config",
				path: "testdata",
				typ:  "yml",
			},
			want: config.Config{
				Env:  "testing",
				HTTP: config.HTTP{Port: 8080},
			},
		},
		{
			name: "success new config. override env",
			args: args{
				name: "test_config",
				path: "testdata",
				typ:  "yml",
			},
			setEnv: func() {
				e := "ENV"
				p := "HTTP_PORT"
				os.Setenv(e, "override_testing")
				os.Setenv(p, "1234")
			},
			want: config.Config{
				Env:  "override_testing",
				HTTP: config.HTTP{Port: 1234},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.setEnv != nil {
				tt.setEnv()
			}
			got := config.NewConfig(tt.args.name, tt.args.path, tt.args.typ)

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("(-got +want):\n%s", diff)
			}
		})
	}
}

func TestConfig_ConfigFileNotFound(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("failed config file not found")
		}
	}()
	config.NewConfig("non_existing_config", "testdata", "yml")
}
