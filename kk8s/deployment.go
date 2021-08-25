package kk8s

import (
	v1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeploymentParams struct {
	Name       string
	Labels     map[string]string
	Containers []ContainerParams
}

type ContainerParams struct {
}

func GenDeployment(dp DeploymentParams) *v1.Deployment {
	dp.Labels["app"] = dp.Name
	deployment := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: dp.Name,
		},
		Spec: v1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": dp.Name,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: dp.Labels,
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:            "vc-web",
							Image:           "nginx:1.13.5-alpine",
							ImagePullPolicy: "IfNotPresent",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
	return deployment
}

func int32Ptr(i int32) *int32 {
	return &i
}
