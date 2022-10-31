package image

import (
	"errors"
	auth3 "github.com/ZeroOneAI/AISecu/pkg/image/auth"
	"github.com/ZeroOneAI/AISecu/pkg/utils"
	"os"
)

const (
	ImageEnvKey   = "IMAGE"
	UseAuthEnvKey = "AUTH"
)

type Image struct {
	image string
	auth  auth3.InfoInterface
}

func (s *Image) Image() string                          { return s.image }
func (s *Image) AuthInfoInterface() auth3.InfoInterface { return s.auth }

func NewInfo() (*Image, error) {
	var errMsg = ""
	var useAuth = false
	var authInfo auth3.InfoInterface = nil
	var err error = nil

	image, imageEnvExist := os.LookupEnv(ImageEnvKey)
	if !imageEnvExist {
		errMsg = "Env \"" + ImageEnvKey + "\" not valid\n"
	}

	useAuth, err = utils.StringToBool(utils.GetEnvOrDefault(UseAuthEnvKey, "false"))
	if err != nil {
		errMsg += "Env \"" + UseAuthEnvKey + "\" not valid\n"
		return nil, errors.New(errMsg)
	}

	if useAuth {
		authInfo, err = auth3.NewAuthInfo()
		if err != nil {
			errMsg += err.Error()
		}
	}
	if errMsg != "" {
		return nil, errors.New(errMsg)
	}

	return &Image{
		image: image,
		auth:  authInfo,
	}, nil
}
