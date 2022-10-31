package scanner

import (
	"errors"
	"github.com/ZeroOneAI/AISecu/cmd/webhook/config/viper/scanner/zeroonescanner"
	"github.com/ZeroOneAI/AISecu/pkg/scanner"
	"github.com/spf13/viper"
)

const (
	TypeZerooneScanner = "zerooneScanner"
)

type Config struct {
	Type           string                 `yaml:"type"`
	ZerooneScanner *zeroonescanner.Config `yaml:"zerooneScanner"`
}

func SetConfigDefault(basePath string, config *viper.Viper) {
	config.SetDefault(basePath+".type", TypeZerooneScanner)
	zeroonescanner.SetConfigDefault(basePath+".zerooneScanner", config)
}

func (c *Config) GetController() (scanner.Interface, error) {
	switch c.Type {
	case TypeZerooneScanner:
		return c.ZerooneScanner.CreateController()
	}

	return nil, errors.New("unknown scanner type")
}
