package viper

import (
	"github.com/ZeroOneAI/AISecu/cmd/webhook/config/viper/scanner"
	"github.com/spf13/viper"
)

type Config struct {
	UserInfoEndpoint string          `yaml:"userInfoEndpoint"`
	Scanner          *scanner.Config `yaml:"scanner"`
}

func SetConfigDefault(config *viper.Viper) {
	config.SetDefault("userInfoEndpoint", "http://localhost:8080")
	scanner.SetConfigDefault("scanner", config)
}

func NewFromFile(filepath string) (*Config, error) {
	viperConfig := viper.New()

	SetConfigDefault(viperConfig)
	viperConfig.SetConfigFile(filepath)

	err := viperConfig.ReadInConfig()
	if err != nil {
		return nil, err
	}

	conf := &Config{}
	err = viperConfig.Unmarshal(conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
