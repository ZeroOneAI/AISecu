package dependencytrack

import (
	"github.com/ZeroOneAI/AISecu/pkg/depedencydb/dependencytrack"
	"github.com/spf13/viper"
)

type Config struct {
	Endpoint string `json:"endpoint"`
	ApiKey   string `json:"apiKey"`
}

func SetConfigDefault(basePath string, config *viper.Viper) {
	config.SetDefault(basePath+".endpoint", "http://localhost:8080")
	config.SetDefault(basePath+".api_key", "")
}

func (c *Config) CreateController() (*dependencytrack.Controller, error) {
	return dependencytrack.New(c.Endpoint, c.ApiKey), nil
}
