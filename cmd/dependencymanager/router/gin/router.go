package gin

import (
	"github.com/ZeroOneAI/AISecu/cmd/dependencymanager/config/viper"
	"github.com/ZeroOneAI/AISecu/pkg/depedencydb"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Router struct {
	*gin.Engine
	dependencyManager depedencydb.Interface
	urlRootPath       string
}

const (
	ExternalApiPrefix = "/api"
	InternalApiPrefix = "/internal"
)

func NewRouter(configFilePath string) (*Router, error) {
	config, err := viper.NewFromFile(configFilePath)
	if err != nil {
		return nil, err
	}
	dependencyManager, err := config.GetController()
	if err != nil {
		return nil, err
	}
	r := &Router{
		Engine:            gin.Default(),
		dependencyManager: dependencyManager,
		urlRootPath:       "",
	}
	r.setRoute()
	return r, nil
}

func (r *Router) setRoute() {
	r.GET(r.urlRootPath+ExternalApiPrefix+"/metrics/image/:imageId", r.GetImageMetrics)
	r.DELETE(r.urlRootPath+InternalApiPrefix+"/metrics/repository/:repositoryId", r.DeleteMetricsRelatedToRepository)
	r.GET(r.urlRootPath+ExternalApiPrefix+"/cve/image/:imageId", r.GetImageCVEList)
}

func (r *Router) GetImageMetrics(c *gin.Context) {
	imageId := c.Param("imageId")

	metrics, err := r.dependencyManager.GetImageMetrics(imageId)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, metrics)
}

func (r *Router) GetImageCVEList(c *gin.Context) {
	imageId := c.Param("imageId")

	cveList, err := r.dependencyManager.GetImageCVEListByImageId(imageId)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, cveList)
}

func (r *Router) DeleteMetricsRelatedToRepository(c *gin.Context) {
	repositoryId := c.Param("repositoryId")

	err := r.dependencyManager.DeleteProjectByProjectName(repositoryId)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
