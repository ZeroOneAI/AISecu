package docker_hub

import (
	"fmt"
	scanner2 "github.com/ZeroOneAI/AISecu/cmd/webhook/config/viper/scanner"
	"github.com/ZeroOneAI/AISecu/pkg/scanner"
	"github.com/ZeroOneAI/AISecu/pkg/userinfo"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	RegistryName = "dockerhub"
	RegistryUrl  = "docker.io"
)

type Controller struct {
	userInfo          userinfo.Interface
	scannerController scanner.Interface
}

func NewController(userInfo userinfo.Interface, config *scanner2.Config) (*Controller, error) {
	var scannerController scanner.Interface
	var err error

	scannerController, err = config.GetController()
	if err != nil {
		return nil, err
	}

	return &Controller{
		userInfo:          userInfo,
		scannerController: scannerController,
	}, nil
}

func (con *Controller) Scan(c *gin.Context) {
	imageData := &PayLoad{}
	username := c.Param("username")
	repositoryName := c.Param("repositoryName")
	repositoryId := c.Param("repositoryId")
	accountId := c.Param("accountId")

	err := c.Bind(imageData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	if imageData.Repository.Name != repositoryName || imageData.Repository.Owner != username {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	pw, err := con.userInfo.GetPassword(accountId)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	err = con.scannerController.StartScan(RegistryUrl, repositoryId, username, pw, imageData.Repository.RepoName, imageData.PushData.Tag)
	if err != nil {
		fmt.Println("start scan error", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	err = con.userInfo.RepositoryPushed(repositoryId, imageData.PushData.Tag)
	if err != nil {
		fmt.Println("repository push to userinfo failed")
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	fmt.Println("success")
	c.JSON(http.StatusOK, gin.H{})
}
