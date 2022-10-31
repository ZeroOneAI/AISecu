package datastruct

import (
	"errors"
	"github.com/ZeroOneAI/AISecu/cmd/projectmanager/config/viper/datastruct/mongo"
	"github.com/ZeroOneAI/AISecu/pkg/db/mongodb"
	"github.com/spf13/viper"
)

const (
	TypeMongoDb = "mongoDB"
)

type Config struct {
	Type    string        `yaml:"type"`
	MongoDB *mongo.Config `yaml:"mongoDB"`
}

func SetConfigDefault(basePath string, config *viper.Viper) {
	config.SetDefault(basePath+".type", TypeMongoDb)
	mongo.SetConfigDefault(basePath+".mongoDB", config)
}

func (c *Config) GetController() (*mongodb.Controller, error) {
	switch c.Type {
	case TypeMongoDb:
		return c.MongoDB.CreateController()
	}

	return nil, errors.New("unknown db type")
}
