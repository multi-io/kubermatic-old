package kubermatic

import (
	operatorv1alpha1 "github.com/kubermatic/kubermatic/api/pkg/crd/operator/v1alpha1"
	"github.com/kubermatic/kubermatic/api/pkg/resources/reconciling"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	policyv1beta1 "k8s.io/api/policy/v1beta1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func uiPodLabels() map[string]string {
	return map[string]string{
		nameLabel:    "kubermatic-ui",
		versionLabel: "v1",
	}
}

func UIDeploymentCreator(ns string, cfg *operatorv1alpha1.KubermaticConfiguration) reconciling.NamedDeploymentCreatorGetter {
	return func() (string, reconciling.DeploymentCreator) {
		return uiDeploymentName, func(d *appsv1.Deployment) (*appsv1.Deployment, error) {
			specLabels := uiPodLabels()

			d.Spec.Replicas = i32ptr(2)
			d.Spec.Selector = &metav1.LabelSelector{
				MatchLabels: specLabels,
			}

			d.Spec.Template.Labels = specLabels
			d.Spec.Template.Spec.ImagePullSecrets = []corev1.LocalObjectReference{
				{
					Name: dockercfgSecretName,
				},
			}

			d.Spec.Template.Spec.Containers = []corev1.Container{
				{
					Name:            "webserver",
					Image:           "quay.io/kubermatic/ui-v2:v1.3.0",
					ImagePullPolicy: corev1.PullIfNotPresent,
					Ports: []corev1.ContainerPort{
						{
							Name:          "http",
							ContainerPort: 8080,
							Protocol:      corev1.ProtocolTCP,
						},
					},
					VolumeMounts: []corev1.VolumeMount{
						{
							MountPath: "/dist/config/",
							Name:      "config",
							ReadOnly:  true,
						},
					},
					Resources: corev1.ResourceRequirements{
						Requests: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse("100m"),
							corev1.ResourceMemory: resource.MustParse("64Mi"),
						},
						Limits: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse("250m"),
							corev1.ResourceMemory: resource.MustParse("128Mi"),
						},
					},
					TerminationMessagePolicy: corev1.TerminationMessageReadFile,
					TerminationMessagePath:   "/dev/termination-log",
				},
			}

			d.Spec.Template.Spec.Volumes = []corev1.Volume{
				{
					Name: "config",
					VolumeSource: corev1.VolumeSource{
						ConfigMap: &corev1.ConfigMapVolumeSource{
							DefaultMode:          i32ptr(420),
							LocalObjectReference: corev1.LocalObjectReference{Name: uiConfigConfigMapName},
						},
					},
				},
			}

			return d, nil
		}
	}
}

func UIPDBCreator(ns string, cfg *operatorv1alpha1.KubermaticConfiguration) reconciling.NamedPodDisruptionBudgetCreatorGetter {
	name := "kubermatic-ui-v2"

	return func() (string, reconciling.PodDisruptionBudgetCreator) {
		return name, func(pdb *policyv1beta1.PodDisruptionBudget) (*policyv1beta1.PodDisruptionBudget, error) {
			min := intstr.FromInt(1)

			pdb.Spec.MinAvailable = &min
			pdb.Spec.Selector = &metav1.LabelSelector{
				MatchLabels: uiPodLabels(),
			}

			return pdb, nil
		}
	}
}

func UIServiceCreator(ns string, cfg *operatorv1alpha1.KubermaticConfiguration) reconciling.NamedServiceCreatorGetter {
	return func() (string, reconciling.ServiceCreator) {
		return uiServiceName, func(s *corev1.Service) (*corev1.Service, error) {
			s.Spec.Type = corev1.ServiceTypeNodePort
			s.Spec.Selector = uiPodLabels()

			s.Spec.Ports = mergeServicePort(s.Spec.Ports, corev1.ServicePort{
				Port:       80,
				TargetPort: intstr.FromInt(8080),
				Protocol:   corev1.ProtocolTCP,
			})

			return s, nil
		}
	}
}
