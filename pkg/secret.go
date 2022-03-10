package secret

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	coreV1 "k8s.io/api/core/v1"
	corev1Type "k8s.io/client-go/kubernetes/typed/core/v1"
)

type Secret struct {
	o      coreV1.Secret
	n      coreV1.Secret
	client corev1Type.CoreV1Client
}

func objMeta(meta *metav1.ObjectMeta) {
	newAnnotations := removeAnnotations(meta.Annotations)
	emptyManagedFields := []metav1.ManagedFieldsEntry{}
	meta.SetSelfLink("")
	meta.SetResourceVersion("")
	meta.SetUID("")
	meta.SetCreationTimestamp(metav1.Time{})
	meta.SetManagedFields(emptyManagedFields)
	meta.SetAnnotations(newAnnotations)
}

func removeAnnotations(annotations map[string]string) map[string]string {
	delete(annotations, "kubectl.kubernetes.io/last-applied-configuration")
	delete(annotations, "hash")
	delete(annotations, "fluxcd.io/sync-checksum")
	return annotations
}

func (s *Secret) removeGen() {
	secretKopy := s.o.DeepCopy()
	objMeta(&secretKopy.ObjectMeta)
	s.n = *secretKopy
}

func (s *Secret) GetSecret(ns, secret string) *coreV1.Secret {
	result, _ := s.client.Secrets(ns).Get(context.TODO(), secret, metav1.GetOptions{})
	return result
}
