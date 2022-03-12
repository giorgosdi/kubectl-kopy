package secret

import (
	"context"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	coreV1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

type Secret struct {
	O      coreV1.Secret
	N      coreV1.Secret
	Client kubernetes.Clientset
}

func objMeta(meta *metav1.ObjectMeta) {
	newAnnotations := removeAnnotations(meta.Annotations)
	emptyManagedFields := []metav1.ManagedFieldsEntry{}
	meta.SetSelfLink("")
	meta.SetNamespace("")
	meta.SetResourceVersion("")
	meta.SetUID("")
	meta.SetCreationTimestamp(metav1.Time{})
	meta.SetManagedFields(emptyManagedFields)
	meta.SetAnnotations(newAnnotations)
	meta.SetOwnerReferences([]metav1.OwnerReference{})
}

func removeAnnotations(annotations map[string]string) map[string]string {
	delete(annotations, "kubectl.kubernetes.io/last-applied-configuration")
	delete(annotations, "hash")
	delete(annotations, "fluxcd.io/sync-checksum")
	delete(annotations, "helm.fluxcd.io/antecedent")
	return annotations
}

func removeLabels(labels map[string]string) map[string]string {
	delete(labels, "fluxcd.io/sync-gc-mark")
	return labels
}

func (s *Secret) RemoveGen() {
	secretKopy := s.O.DeepCopy()
	objMeta(&secretKopy.ObjectMeta)
	s.N = *secretKopy
}

func (s *Secret) Kopy() {
	s.RemoveGen()
	result, err := s.Client.CoreV1().Secrets("default").Create(context.TODO(), &s.N, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(result)
}

func (s *Secret) GetSecret(ns, secret string) *coreV1.Secret {
	result, _ := s.Client.CoreV1().Secrets(ns).Get(context.TODO(), secret, metav1.GetOptions{})
	return result
}
