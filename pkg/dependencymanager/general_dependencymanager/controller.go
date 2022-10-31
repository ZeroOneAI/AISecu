package general_dependencymanager

import (
	"fmt"
	"net/http"
)

type Controller struct {
	hostname string
	client   *http.Client
}

func New(hostname string) *Controller {
	return &Controller{
		hostname: hostname,
		client:   &http.Client{},
	}
}

func (c *Controller) DeleteMetricsRelatedToRepository(repositoryId string) error {
	endpoint := c.hostname + "/internal/metrics/repository/" + repositoryId

	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}
	req.Close = true

	respRaw, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer respRaw.Body.Close()

	if respRaw.StatusCode != 200 {
		return fmt.Errorf("invalid status code [%d]", respRaw.StatusCode)
	}
	return nil
}
