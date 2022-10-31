package webhook

import (
	"errors"
	"github.com/ZeroOneAI/AISecu/pkg/userinfo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	router             *gin.Engine
	registryMap        map[string]RegistryInterface
	userInfoController userinfo.Interface
}

func (c *Controller) Router() *gin.Engine { return c.router }

func NewController(info userinfo.Interface) *Controller {
	c := &Controller{
		registryMap:        make(map[string]RegistryInterface),
		router:             gin.Default(),
		userInfoController: info,
	}

	c.router.POST("/webhook/scan/:repositoryId", c.scanRoute)

	return c
}

func (c *Controller) AddRegistry(registryName string, r RegistryInterface) error {
	if _, exist := c.registryMap[registryName]; exist {
		return errors.New(registryName + " is already exist")
	}
	c.registryMap[registryName] = r
	return nil
}

func (c *Controller) scanRoute(con *gin.Context) {
	registry, err := c.getRegistryInterface(con)
	if err != nil {
		if err.Error() == "registry not found" {
			con.JSON(http.StatusNotFound, gin.H{})
		} else {
			con.JSON(http.StatusInternalServerError, gin.H{})
		}
		return
	}
	registry.Scan(con)
}

func (c *Controller) getRegistryInterface(con *gin.Context) (RegistryInterface, error) {
	repositoryId := con.Param("repositoryId")

	con.AddParam("repositoryId", repositoryId)
	err := c.userInfoController.SetRepositoryInContextByRepositoryId(con, repositoryId)
	if err != nil {
		return nil, err
	}
	err = c.userInfoController.SetAccountInContextByAccountId(con, con.Param("accountId"))
	if err != nil {
		return nil, err
	}
	registry, exist := c.registryMap[con.Param("registryType")]
	if !exist {
		return nil, errors.New("registry not found")
	}
	return registry, nil
}
