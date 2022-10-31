package projectmanager

import (
	"github.com/ZeroOneAI/AISecu/cmd/projectmanager/router/gin"
)

type Interface interface {
	GetImageByImageId(imageId string) (*gin.Image, error)
}
