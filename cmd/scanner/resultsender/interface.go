package resultsender

type Interface interface {
	Send(filepath string) error
}
