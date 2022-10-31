package depedencydb

import (
	"github.com/ZeroOneAI/AISecu/pkg/depedencydb/dependencytrack"
)

type Interface interface {
	GetImageMetrics(imageId string) (*dependencytrack.Metrics, error)
	GetImageCVEListByImageId(imageId string) ([]*dependencytrack.CVE, error)
	GetImageCVEListByRepositoryIdAndTag(repositoryId, tag string) ([]*dependencytrack.CVE, error)
	DeleteProjectByProjectName(projectName string) error
}
