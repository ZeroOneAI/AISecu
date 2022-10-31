package zeroonescanner

import (
	"github.com/ZeroOneAI/AISecu/pkg/scanner/zeroonescanner"
	"github.com/spf13/viper"
)

type Config struct {
	Namespace               string `yaml:"namespace"`
	ScannerImage            string `yaml:"scannerImage"`
	ResultSenderImage       string `yaml:"resultSenderImage"`
	DependencyTrackEndpoint string `yaml:"dependencyTrackEndpoint"`
	DependencyTrackApiKey   string `yaml:"dependencyTrackApiKey"`
}

func SetConfigDefault(basePath string, config *viper.Viper) {
	config.SetDefault(basePath+".namespace", "default")
	config.SetDefault(basePath+".scannerImage", "zerooneai/scanner:t0.0.4")
	config.SetDefault(basePath+".resultSenderImage", "zerooneai/resultsender:dependencytrack-0.0.3")
	config.SetDefault(basePath+".dependencyTrackEndpoint", "")
	config.SetDefault(basePath+".dependencyTrackApiKey", "")
}

func (c *Config) CreateController() (*zeroonescanner.Controller, error) {
	return zeroonescanner.NewController(c.Namespace, c.ScannerImage, c.ResultSenderImage, c.DependencyTrackEndpoint, c.DependencyTrackApiKey)
}
