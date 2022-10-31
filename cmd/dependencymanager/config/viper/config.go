package viper

import (
	"errors"
	"github.com/ZeroOneAI/AISecu/cmd/dependencymanager/config/viper/dependencytrack"
	"github.com/ZeroOneAI/AISecu/pkg/depedencydb"
	"github.com/spf13/viper"
)

const (
	TypeDependencyTrack = "dependencyTrack"
)

type Config struct {
	Type            string                  `json:"type"`
	DependencyTrack *dependencytrack.Config `json:"dependencyTrack"`
}

func SetConfigDefault(config *viper.Viper) {
	config.SetDefault("type", TypeDependencyTrack)
	dependencytrack.SetConfigDefault("dependency_track", config)
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

func (c *Config) GetController() (depedencydb.Interface, error) {
	switch c.Type {
	case TypeDependencyTrack:
		return c.DependencyTrack.CreateController()
	}
	return nil, errors.New("unknown dependency type")
}
