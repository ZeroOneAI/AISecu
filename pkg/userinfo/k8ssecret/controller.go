package k8ssecret

import (
	"context"
	"errors"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	k "k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

const (
	PasswordKey = "password"
)

type Controller struct {
	cs        k.Interface
	namespace string
}

func NewController(namespace string) (*Controller, error) {
	conf, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	cs, err := k.NewForConfig(conf)
	if err != nil {
		return nil, err
	}

	return &Controller{
		cs:        cs,
		namespace: namespace,
	}, nil
}

func (c *Controller) AddAccount(registryName, accountUsername, accountPassword string) error {
	ctx := context.Background()
	secretController := c.cs.CoreV1().Secrets(c.namespace)
	secret := genSecret(registryName, accountUsername, accountPassword)
	options := metav1.CreateOptions{}

	_, err := secretController.Create(ctx, secret, options)
	return err
}

func (c *Controller) AddRepositoryFromAccount(registryName, repositoryName, accountUsername string) error {
	ctx := context.Background()
	secretController := c.cs.CoreV1().Secrets(c.namespace)
	secretName := genSecretName(registryName, accountUsername)
	labelPatch := []byte(`[{"op":"add","path":"/metadata/annotations/` + repositoryName + `","value":"true"}]`)
	options := metav1.PatchOptions{}

	_, err := secretController.Patch(ctx, secretName, types.JSONPatchType, labelPatch, options)

	return err
}

func (c *Controller) IsUserPermittedToRepository(registryName, repositoryName, accountUsername string) bool {
	ctx := context.Background()
	secretController := c.cs.CoreV1().Secrets(c.namespace)
	secretName := genSecretName(registryName, accountUsername)
	options := metav1.GetOptions{}

	secret, err := secretController.Get(ctx, secretName, options)
	if err != nil {
		return false
	}
	_, exist := secret.ObjectMeta.Annotations[repositoryName]
	return exist
}

func (c *Controller) GetPassword(registryName, accountUsername string) (string, error) {
	ctx := context.Background()
	secretController := c.cs.CoreV1().Secrets(c.namespace)
	secretName := genSecretName(registryName, accountUsername)
	options := metav1.GetOptions{}

	secret, err := secretController.Get(ctx, secretName, options)
	if err != nil {
		return "", err
	}
	pw, exist := secret.Data[PasswordKey]
	if !exist {
		return "", errors.New("no password data exist")
	}
	return string(pw), nil
}

func genSecretName(registryName, accountUsername string) string {
	return registryName + "-" + accountUsername
}

func genSecret(registryName, accountUsername, accountPassword string) *apiv1.Secret {
	// TODO ecr 의 경우 accountUsername 이 전부 root 로 동일함 (도메인이 분리됨)
	secretName := genSecretName(registryName, accountUsername)
	return &apiv1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: secretName,
			// TODO Annotations 가 아예 없다면 AddRepositoryFromAccount 호출 시 Patch 에서 에러가 발생
			Annotations: map[string]string{"a": "a"},
		},
		Data: map[string][]byte{
			PasswordKey: []byte(accountPassword),
		},
	}
}
