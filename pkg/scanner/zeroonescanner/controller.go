package zeroonescanner

import (
	"context"
	"github.com/google/uuid"
	batchv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k "k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

type Controller struct {
	cs                      k.Interface
	namespace               string
	scannerImage            string
	resultSenderImage       string
	dependencyTrackEndpoint string
	apiKey                  string
}

func NewController(namespace, scannerImage, resultSenderImage, dependencyTrackEndpoint, apiKey string) (*Controller, error) {
	conf, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	cs, err := k.NewForConfig(conf)
	if err != nil {
		return nil, err
	}

	return &Controller{
		cs:                      cs,
		namespace:               namespace,
		scannerImage:            scannerImage,
		resultSenderImage:       resultSenderImage,
		dependencyTrackEndpoint: dependencyTrackEndpoint,
		apiKey:                  apiKey,
	}, nil
}

func (c *Controller) StartScan(registryUrl, repositoryId, username, password, imageName, imageTag string) error {
	ctx := context.Background()
	jobController := c.cs.BatchV1().Jobs(c.namespace)
	job := c.genJob(registryUrl, repositoryId, username, password, imageName, imageTag)
	options := metav1.CreateOptions{}

	_, err := jobController.Create(ctx, job, options)
	return err
}

func (c *Controller) genJob(registryUrl, repositoryId, username, password, imageName, imageTag string) *batchv1.Job {
	imageFullName := imageName + ":" + imageTag
	return &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: "scanner-" + username + "-" + uuid.New().String(),
		},
		Spec: batchv1.JobSpec{
			BackoffLimit:            int32ptr(1),
			TTLSecondsAfterFinished: int32ptr(300),
			Template: apiv1.PodTemplateSpec{
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "dependency-track-sender",
							Image: c.resultSenderImage,
							Env: []apiv1.EnvVar{
								{Name: "API_KEY", Value: c.apiKey},
								{Name: "HOSTNAME", Value: c.dependencyTrackEndpoint},
								{Name: "PROJECT_NAME", Value: repositoryId},
								{Name: "PROJECT_VERSION", Value: imageTag},
							},
						},
						{
							Name:  "scanner",
							Image: c.scannerImage,
							Env: []apiv1.EnvVar{
								{Name: "IMAGE", Value: imageFullName},
								{Name: "AUTH", Value: "true"},
								{Name: "REGISTRY_URL", Value: registryUrl},
								{Name: "ID", Value: username},
								{Name: "PASSWORD", Value: password},
							},
						},
					},
					RestartPolicy: apiv1.RestartPolicyNever,
				},
			},
		},
	}
}

func int32ptr(num int) *int32 {
	n := int32(num)
	return &n
}
