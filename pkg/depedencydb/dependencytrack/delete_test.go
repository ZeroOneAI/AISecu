package dependencytrack

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestController_DeleteProjectByProjectName(t *testing.T) {
	art := assert.New(t)

	c := &Controller{
		dependencyTrackApiKey:   "1CrdXtGgGXML5I7b1YsOcp7hbIQMWbL6",
		dependencyTrackHostname: "https://dependency-track.cloud.ainode.ai",
		client:                  &http.Client{},
	}
	err := c.DeleteProjectByProjectName("dockerhub/docker.io/zerooneai/scan-test")
	art.NoError(err)
}
