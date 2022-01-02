package engine

import "github.com/spf13/viper"

type Config struct {
	Title       string `mapstructure:"title"`
	Description string `mapstructure:"description"`
}

func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	var config Config
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	err := viper.Unmarshal(&config)
	return config, err
}
