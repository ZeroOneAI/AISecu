package trivy

import (
	"github.com/ZeroOneAI/AISecu/pkg/image"
	"os"
	"os/exec"
)

type Trivy struct{}

func NewTrivy() *Trivy { return &Trivy{} }

func (t *Trivy) Scan(image image.InfoInterface, resultFilePath string) error {
	cmd := exec.Command("trivy", "i", "-i", image.Image(), "--format", "cyclonedx", "-o", resultFilePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
