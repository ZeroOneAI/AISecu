package dependencytrack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestController_ListImageTags(t *testing.T) {
	art := assert.New(t)

	c := &Controller{
		dependencyTrackApiKey:   "1CrdXtGgGXML5I7b1YsOcp7hbIQMWbL6",
		dependencyTrackHostname: "https://dependency-track.cloud.ainode.ai",
		client:                  &http.Client{},
	}

	uuids, err := c.getProjectUuidListByProjectName("")
	art.NoError(err)
	fmt.Println(uuids)
}
