package scanner

import (
	"github.com/ZeroOneAI/AISecu/pkg/image"
)

type Interface interface {
	Scan(image image.InfoInterface, resultFilePath string) error
}
