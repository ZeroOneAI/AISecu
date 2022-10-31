package projectmanagercontroller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ZeroOneAI/AISecu/cmd/projectmanager/router/gin"
	"net/http"
)

type Controller struct {
	hostname string
}

func New(hostname string) *Controller {
	return &Controller{
		hostname: hostname,
	}
}

func (c *Controller) GetImageByImageId(imageId string) (*gin.Image, error) {
	rawResp, err := http.Get(c.hostname + "/api/image/detail/" + imageId)
	if err != nil {
		return nil, err
	}
	if rawResp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Invalid response status code %d", rawResp.StatusCode))
	}
	resp := gin.GetImageResponse{}
	err = json.NewDecoder(rawResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp.Image, nil
}

func (c *Controller) GetLatestImageByRepositoryId(repositoryId string) (*gin.Image, error) {
	rawResp, err := http.Get(c.hostname + "/api/repository/latest/" + repositoryId)
	if err != nil {
		return nil, err
	}
	if rawResp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Invalid response status code %d", rawResp.StatusCode))
	}
	resp := gin.GetLatestImageByRepository{}
	err = json.NewDecoder(rawResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp.Image, nil
}
