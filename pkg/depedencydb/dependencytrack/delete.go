package dependencytrack

import (
	"fmt"
	"net/http"
)

func (c *Controller) DeleteProjectByProjectName(projectName string) error {
	projectIdList, err := c.getProjectUuidListByProjectName(projectName)
	if err != nil {
		return err
	}
	for _, projectId := range projectIdList {
		err = c.deleteProjectById(projectId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Controller) deleteProjectById(projectId string) error {
	endpoint := c.dependencyTrackHostname + "/api/v1/project/" + projectId

	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}
	req.Close = true
	req.Header.Add("X-API-Key", c.dependencyTrackApiKey)

	respRaw, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer respRaw.Body.Close()

	if respRaw.StatusCode != 204 {
		return fmt.Errorf("invalid status code [%d]", respRaw.StatusCode)
	}
	return nil
}
