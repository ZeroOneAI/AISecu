package main

import (
	"github.com/ZeroOneAI/AISecu/cmd/resultsender/sender/dependency_track"
	"github.com/ZeroOneAI/AISecu/pkg/utils"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	apiKey := utils.MustGetEnv("API_KEY")
	hostname := utils.GetEnvOrDefault("HOSTNAME", "http://localhost")
	projectName := utils.MustGetEnv("PROJECT_NAME")
	projectVersion := utils.MustGetEnv("PROJECT_VERSION")

	controller := dependency_track.NewController(apiKey, hostname)
	r := gin.Default()
	fin := make(chan error, 1)
	r.POST("/scan/result", func(c *gin.Context) {
		bom, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
			fin <- err
			return
		}
		if err = controller.Send(bom, projectName, projectVersion); err != nil {
			fin <- err
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		c.JSON(http.StatusOK, gin.H{})
		fin <- nil
	})
	go func() {
		err := <-fin
		time.Sleep(time.Second)
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	}()
	if err := r.Run(":3000"); err != nil {
		panic(err)
	}
}
