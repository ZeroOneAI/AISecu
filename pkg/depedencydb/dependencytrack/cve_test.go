package dependencytrack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestController_getCVEById(t *testing.T) {
	art := assert.New(t)

	c := &Controller{
		dependencyTrackApiKey:   "1CrdXtGgGXML5I7b1YsOcp7hbIQMWbL6",
		dependencyTrackHostname: "https://dependency-track.cloud.ainode.ai",
		client:                  &http.Client{},
	}

	cve, err := c.getCVEById("430bcaec-71c2-4474-9728-2f600734543b")
	if err != nil {
		art.Error(err)
		return
	}
	fmt.Println(len(cve))
}
