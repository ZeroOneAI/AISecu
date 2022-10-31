package main

import (
	"errors"
	"flag"
	"github.com/ZeroOneAI/AISecu/cmd/webhook/config/viper"
	"github.com/ZeroOneAI/AISecu/cmd/webhook/webhook"
	"github.com/ZeroOneAI/AISecu/cmd/webhook/webhook/docker_hub"
	"github.com/ZeroOneAI/AISecu/pkg/userinfo"
	"github.com/ZeroOneAI/AISecu/pkg/userinfo/projectmanager"
	"github.com/gin-gonic/gin"
)

func main() {
	port := ":8080"
	r, err := newRoute()
	if err != nil {
		panic(err)
	}

	if err = r.Run(port); err != nil {
		panic(err)
	}
}

func newRoute() (*gin.Engine, error) {
	var webhookController webhook.Interface
	var userInfoController userinfo.Interface
	var err error

	configFilePath := flag.String("config", "", "Define config file path")
	flag.Parse()
	if configFilePath == nil {
		return nil, errors.New("config option must be defined")
	}
	config, err := viper.NewFromFile(*configFilePath)
	if err != nil {
		return nil, err
	}
	userInfoController = projectmanager.New(config.UserInfoEndpoint)
	webhookController = webhook.NewController(userInfoController)

	dockerHubController, err := docker_hub.NewController(userInfoController, config.Scanner)
	if err != nil {
		return nil, err
	}
	if err = webhookController.AddRegistry(docker_hub.RegistryName, dockerHubController); err != nil {
		return nil, err
	}

	return webhookController.Router(), nil
}
