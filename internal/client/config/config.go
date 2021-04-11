package config

import (
	common "github.com/708u/useless-auth-server/internal/pkg/config"
	"github.com/spf13/viper"
)

const (
	ConfigName = "config"
	ConfigPath = "configs/client"
	ConfigType = "yml"
)

type Config struct {
	Env string

	HTTP   HTTP
	Auth   Auth
	Client Client
}

type HTTP struct {
	Port uint16
}

type Auth struct {
	URL          string
	ResponseType string
}

type Client struct {
	ID          string
	Secret      string
	RedirectURI string
}

// NewConfig returns config, incluedes environment vars.
func NewConfig(cName, cPath, cType string) Config {
	var c Config
	common.InitConfig(cName, cPath, cType, func() {
		viper.SetDefault("Auth.ResponseType", "code")
		// var c Config
		if err := viper.Unmarshal(&c); err != nil {
			panic("config file unmarshal failed")
		}
	})
	return c
}
