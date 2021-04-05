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
	Env  string
	HTTP HTTP
}

type HTTP struct {
	Port uint16
}

// NewConfig returns config, incluedes environment vars.
func NewConfig(cName, cPath, cType string) Config {
	var c Config
	common.InitConfig(cName, cPath, cType, func() {
		// var c Config
		if err := viper.Unmarshal(&c); err != nil {
			panic("config file unmarshal failed")
		}
	})
	return c
}
