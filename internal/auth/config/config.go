package config

import (
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	EnvLocal = "local"
	EnvDev   = "development"
	EnvStg   = "staging"
	EnvProd  = "production"
)

type Config struct {
	Env  string
	HTTP struct {
		Port uint16
	}
}

// NewConfig returns config, incluedes environment vars.
func NewConfig() Config {
	return initConfig()
}

func initConfig() Config {
	// config file setting
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")

	// env setting
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
	})

	if err := viper.ReadInConfig(); err != nil {
		panic("config file is not found")
	}
	fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		panic("config file unmarshal failed")
	}

	return c
}
