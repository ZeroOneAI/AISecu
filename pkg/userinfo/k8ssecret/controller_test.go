package k8ssecret

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestController_AddAccount(t *testing.T) {
	art := assert.New(t)

	c, err := NewController("default")
	if err != nil {
		art.NoError(err)
		return
	}
	err = c.AddAccount("dockerhub", "myoon", "david18424")
	art.NoError(err)
}

func TestController_AddRepositoryFromAccount(t *testing.T) {
	art := assert.New(t)

	c, err := NewController("default")
	if err != nil {
		art.NoError(err)
		return
	}
	err = c.AddRepositoryFromAccount("dockerhub", "scan-test", "myoon")
	art.NoError(err)
}

func TestController_IsUserPermittedToRepository(t *testing.T) {
	art := assert.New(t)

	c, err := NewController("default")
	if err != nil {
		art.NoError(err)
		return
	}
	isPermitted := c.IsUserPermittedToRepository("dockerhub", "scan-test", "myoon")
	if !isPermitted {
		art.NoError(errors.New("access deny"))
	}
}
