package config_test

import (
	"os"
	"testing"

	"github.com/708u/useless-auth-server/internal/client/config"
	"github.com/google/go-cmp/cmp"
)

func TestConfig_NewConfig(t *testing.T) {
	t.Parallel()

	type args struct {
		name string
		path string
		typ  string
	}
	wantConf := config.Config{
		Env:  "testing",
		HTTP: config.HTTP{Port: 9000},
		Auth: config.Auth{
			URL:          "http://localhost:9001",
			ResponseType: "code",
		},
		Client: config.Client{
			ID:          "useless-auth-server-client",
			Secret:      "secret",
			RedirectURI: "http://localhost:9000/callback",
		},
	}
	tests := []struct {
		name          string
		args          args
		getWantConfig func() config.Config
		wantRecover   bool
		setEnv        func()
	}{
		{
			name: "success new config",
			args: args{
				name: "test_config",
				path: "testdata",
				typ:  "yml",
			},
			getWantConfig: func() config.Config { return wantConf },
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
			getWantConfig: func() config.Config {
				c := wantConf
				c.Env = "override_testing"
				c.HTTP.Port = 1234
				return c
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

			if diff := cmp.Diff(got, tt.getWantConfig()); diff != "" {
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
