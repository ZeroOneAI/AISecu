package webhook

import "github.com/gin-gonic/gin"

type Interface interface {
	Router() *gin.Engine
	AddRegistry(registryName string, r RegistryInterface) error
}

type RegistryInterface interface {
	Scan(c *gin.Context)
}
