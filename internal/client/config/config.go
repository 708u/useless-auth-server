package config

import (
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	ConfigName = "config"
	ConfigPath = "config/client"
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
	// config file setting
	viper.SetConfigName(cName)
	viper.AddConfigPath(cPath)
	viper.SetConfigType(cType)

	// env setting
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
	})

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("failed to read config: %s", err.Error()))
	}
	fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		panic("config file unmarshal failed")
	}

	return c
}
