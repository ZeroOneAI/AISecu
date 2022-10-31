package dependencytrack

import (
	"encoding/json"
	"github.com/ZeroOneAI/AISecu/pkg/projectmanager"
	"github.com/ZeroOneAI/AISecu/pkg/projectmanager/projectmanagercontroller"
	"net/http"
)

type Controller struct {
	dependencyTrackApiKey    string
	dependencyTrackHostname  string
	client                   *http.Client
	projectManagerController projectmanager.Interface
}

func New(hostname, apiKey string) *Controller {
	return &Controller{
		dependencyTrackApiKey:   apiKey,
		dependencyTrackHostname: hostname,
		// TODO remove hardcode
		projectManagerController: projectmanagercontroller.New("http://projectmanager:8080"),
		client:                   &http.Client{},
	}
}

type Project struct {
	Uuid          string `json:"uuid"`
	Name          string `json:"name"`
	LastBomImport int64  `json:"lastBomImport"`
	Version       string `json:"version"`
}

type Metrics struct {
	Critical                             int     `json:"critical"`
	High                                 int     `json:"high"`
	Medium                               int     `json:"medium"`
	Low                                  int     `json:"low"`
	Unassigned                           int     `json:"unassigned"`
	Vulnerabilities                      int     `json:"vulnerabilities"`
	VulnerableComponents                 int     `json:"vulnerableComponents"`
	Components                           int     `json:"components"`
	Suppressed                           int     `json:"suppressed"`
	FindingsTotal                        int     `json:"findingsTotal"`
	FindingsAudited                      int     `json:"findingsAudited"`
	FindingsUnaudited                    int     `json:"findingsUnaudited"`
	InheritedRiskScore                   float64 `json:"inheritedRiskScore"`
	PolicyViolationsFail                 int     `json:"policyViolationsFail"`
	PolicyViolationsWarn                 int     `json:"policyViolationsWarn"`
	PolicyViolationsInfo                 int     `json:"policyViolationsInfo"`
	PolicyViolationsTotal                int     `json:"policyViolationsTotal"`
	PolicyViolationsAudited              int     `json:"policyViolationsAudited"`
	PolicyViolationsUnaudited            int     `json:"policyViolationsUnaudited"`
	PolicyViolationsSecurityTotal        int     `json:"policyViolationsSecurityTotal"`
	PolicyViolationsSecurityAudited      int     `json:"policyViolationsSecurityAudited"`
	PolicyViolationsSecurityUnaudited    int     `json:"policyViolationsSecurityUnaudited"`
	PolicyViolationsLicenseTotal         int     `json:"policyViolationsLicenseTotal"`
	PolicyViolationsLicenseAudited       int     `json:"policyViolationsLicenseAudited"`
	PolicyViolationsLicenseUnaudited     int     `json:"policyViolationsLicenseUnaudited"`
	PolicyViolationsOperationalTotal     int     `json:"policyViolationsOperationalTotal"`
	PolicyViolationsOperationalAudited   int     `json:"policyViolationsOperationalAudited"`
	PolicyViolationsOperationalUnaudited int     `json:"policyViolationsOperationalUnaudited"`
	FirstOccurrence                      int64   `json:"firstOccurrence"`
	LastOccurrence                       int64   `json:"lastOccurrence"`
}

type Repository struct {
	registryType    string
	registryUrl     string
	accountUsername string
	repositoryName  string
}

func (c *Controller) GetImageMetrics(imageId string) (*Metrics, error) {
	image, err := c.projectManagerController.GetImageByImageId(imageId)
	if err != nil {
		return nil, err
	}
	return c.getImageMetrics(image.RepositoryId, image.Tag)
}

func (c *Controller) getImageMetrics(repositoryId, tag string) (*Metrics, error) {
	projectId, err := c.getProjectId(repositoryId, tag)
	if err != nil {
		return nil, err
	}
	return c.getMetricsById(projectId)
}

func (c *Controller) getMetricsById(projectId string) (*Metrics, error) {
	endpoint := c.dependencyTrackHostname + "/api/v1/metrics/project/" + projectId + "/current"

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
	resp := &Metrics{}
	err = json.NewDecoder(respRaw.Body).Decode(resp)
	return resp, err
}

func (c *Controller) getProjectId(projectName, projectVersion string) (string, error) {
	endpoint := c.dependencyTrackHostname + "/api/v1/project/lookup?name=" + projectName + "&version=" + projectVersion

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return "", err
	}
	req.Close = true
	req.Header.Add("X-API-Key", c.dependencyTrackApiKey)

	respRaw, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer respRaw.Body.Close()
	resp := Project{}
	err = json.NewDecoder(respRaw.Body).Decode(&resp)
	if err != nil {
		return "", err
	}
	return resp.Uuid, nil
}

func (c *Controller) getProjectUuidListByProjectName(projectName string) ([]string, error) {
	endpoint := c.dependencyTrackHostname + "/api/v1/project?name=" + projectName

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
	resp := make([]*Project, 0)

	err = json.NewDecoder(respRaw.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}

	projectUuidList := make([]string, len(resp))
	for index := range projectUuidList {
		projectUuidList[index] = resp[index].Uuid
	}
	return projectUuidList, nil
}
