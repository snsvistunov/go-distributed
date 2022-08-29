package config

import "github.com/spf13/viper"

type Config struct {
	SourceFile string `mapstructure:"SOURCE_FILE"`
	OutputFile string `mapstructure:"OUTPUT_FILE"`
}

var Cfg *Config

func LoadConfig(path string) (err error) {
	var config Config
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	Cfg = &config
	return
}
