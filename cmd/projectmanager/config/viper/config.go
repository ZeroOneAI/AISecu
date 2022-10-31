package viper

import (
	"github.com/ZeroOneAI/AISecu/cmd/projectmanager/config/viper/datastruct"
	"github.com/spf13/viper"
)

type Config struct {
	Db          *datastruct.Config `yaml:"db"`
	ApiRootPath string             `yaml:"apiRootPath"`
}

func SetConfigDefault(config *viper.Viper) {
	datastruct.SetConfigDefault("db", config)
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
