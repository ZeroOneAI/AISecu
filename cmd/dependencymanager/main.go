package main

import (
	"flag"
	"github.com/ZeroOneAI/AISecu/cmd/dependencymanager/router/gin"
	"log"
)

func main() {
	configFilePath := flag.String("config", "", "Define config file path")
	flag.Parse()
	if configFilePath == nil {
		log.Panicln("config option must be defined")
		return
	}
	r, err := gin.NewRouter(*configFilePath)
	if err != nil {
		log.Panicln(err)
		return
	}

	r.Run(":8080")
}
