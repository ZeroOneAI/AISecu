package dependencytrack

import (
	"encoding/json"
	"net/http"
)

type Component struct {
	Uuid    string `json:"uuid"`
	Name    string `json:"name"`
	Version string `json:"version"`
	Purl    string `json:"purl"`
	Project string `json:"project"`
}

type Vulnerability struct {
	Uuid           string      `json:"uuid"`
	Source         string      `json:"source"`
	VulnId         string      `json:"vulnId"`
	Severity       string      `json:"severity"`
	SeverityRank   int         `json:"severityRank"`
	CweId          int         `json:"cweId,omitempty"`
	CweName        string      `json:"cweName,omitempty"`
	Description    string      `json:"description"`
	Recommendation interface{} `json:"recommendation"`
}

type Analysis struct {
	IsSuppressed bool `json:"isSuppressed"`
}

type Attribution struct {
	AnalyzerIdentity string `json:"analyzerIdentity"`
	AttributedOn     int64  `json:"attributedOn"`
}

type CVE struct {
	Component     Component     `json:"component"`
	Vulnerability Vulnerability `json:"vulnerability"`
	Analysis      Analysis      `json:"analysis"`
	Attribution   Attribution   `json:"attribution"`
	Matrix        string        `json:"matrix"`
}

func (c *Controller) GetImageCVEListByImageId(imageId string) ([]*CVE, error) {
	image, err := c.projectManagerController.GetImageByImageId(imageId)
	if err != nil {
		return nil, err
	}
	return c.GetImageCVEListByRepositoryIdAndTag(image.RepositoryId, image.Tag)
}

func (c *Controller) GetImageCVEListByRepositoryIdAndTag(repositoryId, tag string) ([]*CVE, error) {
	projectId, err := c.getProjectId(repositoryId, tag)
	if err != nil {
		return nil, err
	}
	return c.getCVEById(projectId)
}

func (c *Controller) getCVEById(projectId string) ([]*CVE, error) {
	endpoint := c.dependencyTrackHostname + "/api/v1/finding/project/" + projectId + "?suppressed=false"

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Close = true
	req.Header.Add("X-API-Key", c.dependencyTrackApiKey)

	respRaw, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer respRaw.Body.Close()
	resp := make([]*CVE, 0)
	err = json.NewDecoder(respRaw.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
