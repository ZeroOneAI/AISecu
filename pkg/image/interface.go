package image

import (
	"github.com/ZeroOneAI/AISecu/pkg/image/auth"
)

type InfoInterface interface {
	Image() string
	AuthInfoInterface() auth.InfoInterface
}
