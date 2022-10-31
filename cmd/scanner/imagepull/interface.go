package imagepull

import (
	"github.com/ZeroOneAI/AISecu/pkg/image"
)

type Interface interface {
	Pull(image image.InfoInterface) error
}
