package serviceaccount

import (
	"context"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	coreV1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

type ServiceAccount struct {
	O      coreV1.ServiceAccount
	N      coreV1.ServiceAccount
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
	delete(annotations, "helm.sh/resource-policy")
	return annotations
}

func removeLabels(labels map[string]string) map[string]string {
	delete(labels, "fluxcd.io/sync-gc-mark")
	return labels
}

func (s *ServiceAccount) RemoveGen() {
	ServiceAccountKopy := s.O.DeepCopy()
	objMeta(&ServiceAccountKopy.ObjectMeta)
	s.N = *ServiceAccountKopy
}

func (s *ServiceAccount) Kopy(target string) {
	s.RemoveGen()
	_, err := s.Client.CoreV1().ServiceAccounts(target).Create(context.TODO(), &s.N, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func (s *ServiceAccount) GetResource(ns, serviceAccount string) {
	result, _ := s.Client.CoreV1().ServiceAccounts(ns).Get(context.TODO(), serviceAccount, metav1.GetOptions{})
	s.O = *result
}
