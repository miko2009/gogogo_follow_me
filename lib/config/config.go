package config

import (
	"github.com/spf13/viper"
	"os"
)


func getViper(configName string, configType string) *viper.Viper {
	v := viper.New()
	v.SetConfigName(configName)
	v.SetConfigType(configType)

	env := os.Getenv("MY_ENV")
	switch env {
	case "dev":
		v.AddConfigPath("conf/dev")
	case "pre":
		v.AddConfigPath("conf/pre")
	case "release":
		v.AddConfigPath("conf/release")
	case "master":
		v.AddConfigPath("conf/master")
	default:
		v.AddConfigPath("conf/dev")
	}
	return v
}


func InitConfig() error {
	err := LoadJsonConfig()
	if err != nil {
		return err
	}
	errToml := LoadTomlConfig()
	return errToml
}