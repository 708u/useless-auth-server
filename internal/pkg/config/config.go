package config

import (
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// InitConfig sets configuration info and marshal config struct via marshalFn.
// e.g.
//
// common.InitConfig(cName, cPath, cType, func() {
// if err := viper.Unmarshal(&c); err != nil {
//		panic("config file unmarshal failed")
//		}
// })
func InitConfig(cName, cPath, cType string, marshalFn func()) {
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

	marshalFn()
}
