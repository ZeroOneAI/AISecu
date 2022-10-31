package dependencymanager

type Interface interface {
	DeleteMetricsRelatedToRepository(repositoryId string) error
}
