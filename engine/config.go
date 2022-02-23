package engine

import (
	"github.com/spf13/viper"
)

type SiteConfig struct {
	BaseURL     string `mapstructure:"baseURL"`
	Title       string `mapstructure:"title"`
	Description string `mapstructure:"description"`
}

func LoadConfig(path string) (SiteConfig, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	var config SiteConfig
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	err := viper.Unmarshal(&config)
	return config, err
}
