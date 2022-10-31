package mongodb

import "context"

func (c *Controller) ListRegistryType(ctx context.Context, sp *int, ep *int) ([]string, error) {
	// TODO implementation
	return []string{"dockerhub"}, nil
}
