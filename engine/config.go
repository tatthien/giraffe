package engine

import (
	"github.com/spf13/viper"
)

type SiteConfig struct {
	BaseURL     string `mapstructure:"baseURL"`
	Title       string `mapstructure:"title"`
	Description string `mapstructure:"description"`
	Port        string `mapstructure:"port"`
	ContentDir  string `mapstructure:"contentDir"`
	OutputDir   string `mapstructure:"outputDir"`
}

func LoadConfig(path string) (SiteConfig, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("OutputDir", "dist")
	viper.SetDefault("Port", "3333")

	viper.AutomaticEnv()

	var config SiteConfig
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	err := viper.Unmarshal(&config)
	return config, err
}
