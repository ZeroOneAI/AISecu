package projectmanager

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	schema "github.com/ZeroOneAI/AISecu/cmd/projectmanager/router/gin"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	hostname string
	client   *http.Client
}

func New(hostname string) *Controller {
	return &Controller{hostname: hostname, client: &http.Client{}}
}

func (c *Controller) SetRepositoryInContextByRepositoryId(con *gin.Context, repositoryId string) error {
	endpoint := c.hostname + "/api/repository/detail/" + repositoryId

	rawResp, err := c.client.Get(endpoint)
	if err != nil {
		return err
	}
	defer rawResp.Body.Close()

	if rawResp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("http request error with status code [%d]", rawResp.StatusCode))
	}
	resp := schema.GetRepositoryResponse{}
	err = json.NewDecoder(rawResp.Body).Decode(&resp)
	if err != nil {
		return err
	}

	con.AddParam("repositoryName", resp.Repository.Name)
	con.AddParam("accountId", resp.Repository.AccountId)
	return nil
}

func (c *Controller) SetAccountInContextByAccountId(con *gin.Context, accountId string) error {
	endpoint := c.hostname + "/api/account/detail/" + accountId

	rawResp, err := c.client.Get(endpoint)
	if err != nil {
		return err
	}
	defer rawResp.Body.Close()

	if rawResp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("http request error with status code [%d]", rawResp.StatusCode))
	}
	resp := schema.GetAccountResponse{}
	err = json.NewDecoder(rawResp.Body).Decode(&resp)
	if err != nil {
		return err
	}

	con.AddParam("registryType", resp.Account.RegistryType)
	con.AddParam("registryUrl", resp.Account.RegistryUrl)
	con.AddParam("username", resp.Account.Username)
	return nil
}

func (c *Controller) GetPassword(accountId string) (string, error) {
	endpoint := c.hostname + "/internal/account/private/" + accountId
	rawResp, err := c.client.Get(endpoint)
	if err != nil {
		return "", err
	}
	defer rawResp.Body.Close()

	if rawResp.StatusCode != 200 {
		return "", errors.New(rawResp.Status)
	}
	resp := schema.GetAccountPrivateResponse{}
	err = json.NewDecoder(rawResp.Body).Decode(&resp)
	if err != nil {
		return "", err
	}
	return resp.AccountPrivate.Password, nil
}

func (c *Controller) RepositoryPushed(repositoryId, tag string) error {
	endpoint := c.hostname + "/api/repository/image/" + repositoryId

	reqBodyStruct := schema.CreateOrUpdateImageRequest{
		Tag: tag,
	}
	reqBodyBytes, err := json.Marshal(reqBodyStruct)
	if err != nil {
		return err
	}
	reqBody := bytes.NewBuffer(reqBodyBytes)
	req, err := http.NewRequest("PUT", endpoint, reqBody)
	if err != nil {
		return err
	}
	rawResp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer rawResp.Body.Close()

	resp := schema.CreateOrUpdateImageResponse{}
	json.NewDecoder(rawResp.Body).Decode(&resp)
	return nil
}
