package dependency_track

import "net/http"

type Controller struct {
	apiKey   string
	hostname string
	client   *http.Client
}

func NewController(apiKey, hostname string) *Controller {
	return &Controller{
		apiKey:   apiKey,
		hostname: hostname,
		client:   &http.Client{},
	}
}
