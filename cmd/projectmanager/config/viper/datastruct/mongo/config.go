package mongo

import (
	"errors"
	"github.com/ZeroOneAI/AISecu/pkg/db/mongodb"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type Config struct {
	Protocol           string `yaml:"protocol"`
	Hostname           string `yaml:"hostname"`
	Port               string `yaml:"port"`
	DatabaseName       string `yaml:"databaseName"`
	Auth               *Auth  `yaml:"auth"`
	DependencyEndpoint string `yaml:"dependencyEndpoint"`
}

type Auth struct {
	User             string `yaml:"user"`
	PasswordFilepath string `yaml:"passwordFilepath"`
}

func SetConfigDefault(basePath string, config *viper.Viper) {
	config.SetDefault(basePath+".protocol", "mongodb")
	config.SetDefault(basePath+".hostname", "localhost")
	config.SetDefault(basePath+".port", "27017")
	config.SetDefault(basePath+".databaseName", "project_manager")
	config.SetDefault(basePath+".dependencyEndpoint", "http://127.0.0.1:8080")
}

func (c *Config) CreateController() (*mongodb.Controller, error) {
	if c == nil {
		return nil, errors.New("Nil Detect")
	}

	var mongoAuthCredential *options.Credential = nil

	if c.Auth != nil {
		rawPassword, err := os.ReadFile(c.Auth.PasswordFilepath)
		if err != nil {
			return nil, err
		}
		mongoAuthCredential = &options.Credential{
			Username: c.Auth.User,
			Password: string(rawPassword),
		}
	}
	return mongodb.NewController(c.Protocol, c.Hostname, c.Port, c.DatabaseName, mongoAuthCredential, c.DependencyEndpoint)
}
