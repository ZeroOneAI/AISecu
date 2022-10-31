package crane

import (
	"fmt"
	"github.com/ZeroOneAI/AISecu/pkg/image"
	"github.com/ZeroOneAI/AISecu/pkg/image/auth"
	"os"
	"os/exec"
	"strings"
)

type Crane struct{}

func NewCrane() *Crane { return &Crane{} }

func (c *Crane) Pull(image image.InfoInterface) error {
	err := c.Auth(image.AuthInfoInterface())
	if err != nil {
		return err
	}

	err = c.MkdirForImage(image.Image())
	if err != nil {
		return err
	}

	return c.PullImage(image.Image())
}

func (c *Crane) PullImage(image string) error {
	cmd := exec.Command("crane", "pull", image, image)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("hi")
	return cmd.Run()
}

func (c *Crane) MkdirForImage(image string) error {
	strSlice := strings.Split(image, "/")
	if len(strSlice) > 0 {
		strSlice = strSlice[:len(strSlice)-1]
	}
	dir := strings.Join(strSlice, "/")
	return os.MkdirAll(dir, os.ModePerm)
}

func (c *Crane) Auth(auth auth.InfoInterface) error {
	if auth == nil {
		return nil
	}
	cmd := exec.Command("crane", "auth", "login", auth.RegistryUrl(), "-u", auth.Id(), "-p", auth.Pwd())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
