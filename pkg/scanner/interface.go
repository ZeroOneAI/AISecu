package scanner

type Interface interface {
	StartScan(registryUrl, repositoryId, username, password, imageName, imageTag string) error
}
