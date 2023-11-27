package config

import "github.com/spf13/viper"

func init() {
	// environment - could be "local", "prod", "dev"
	viper.SetDefault("env", "prod")
}
