package auth

import (
	"errors"
	"github.com/ZeroOneAI/AISecu/pkg/utils"
	"os"
)

const (
	RegistryUrlEnvKey = "REGISTRY_URL"
	IdEnvKey          = "ID"
	PwdEnvKey         = "PASSWORD"
)

type Info struct {
	registryUrl string
	id          string
	pwd         string
}

func (a *Info) RegistryUrl() string { return a.registryUrl }
func (a *Info) Id() string          { return a.id }
func (a *Info) Pwd() string         { return a.pwd }

func NewAuthInfo() (*Info, error) {
	errMsg := ""

	registryUrl := utils.GetEnvOrDefault(RegistryUrlEnvKey, "docker.io")
	id, idExist := os.LookupEnv(IdEnvKey)
	if !idExist {
		errMsg += "Env \"" + IdEnvKey + "\" not valid\n"
	}
	pwd, pwdExist := os.LookupEnv(PwdEnvKey)
	if !pwdExist {
		errMsg += "Env \"" + PwdEnvKey + "\" not valid\n"
	}

	if errMsg != "" {
		return nil, errors.New(errMsg)
	}

	return &Info{
		registryUrl: registryUrl,
		id:          id,
		pwd:         pwd,
	}, nil
}
