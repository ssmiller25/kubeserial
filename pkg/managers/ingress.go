package managers

import (
	appv1alpha1 "github.com/janekbaraniewski/kubeserial/pkg/apis/app/v1alpha1"
	v1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (m *Manager)CreateIngress(cr *appv1alpha1.KubeSerial, device *appv1alpha1.Device, domain string) *v1beta1.Ingress {
	name := m.GetName(cr.Name, device.Name)
	labels := map[string]string{
		"app": name,
	}
	return &v1beta1.Ingress{
		ObjectMeta:	metav1.ObjectMeta {
			Name:			name,
			Namespace:		cr.Namespace,
			Labels:			labels,
			Annotations: 	cr.Spec.Ingress.Annotations,
		},
		Spec:		v1beta1.IngressSpec{
			Rules:		[]v1beta1.IngressRule{
				{
					Host:				device.Name + domain,
					IngressRuleValue:	v1beta1.IngressRuleValue{
						HTTP: &v1beta1.HTTPIngressRuleValue{
							Paths:	[]v1beta1.HTTPIngressPath{
								{
									Path: 		"/",
									Backend: 	v1beta1.IngressBackend {
										ServiceName: m.GetName(cr.Name, device.Name),
										ServicePort: intstr.FromInt(80),
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
