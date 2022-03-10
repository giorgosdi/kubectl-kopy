package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	coreV1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type resource interface {
	removeGen()
}

type secret struct {
	o coreV1.Secret
	n coreV1.Secret
}

func (s *secret) removeGen() {
	secretKopy := s.o.DeepCopy()
	objMeta(&secretKopy.ObjectMeta)
	s.n = *secretKopy
}

type res struct {
	r resource
}

func (resource res) removeGen() {
	resource.r.removeGen()
}

func GetClients(conf string) *kubernetes.Clientset {
	config, _ := clientcmd.BuildConfigFromFlags("", conf)
	resourceClient, _ := kubernetes.NewForConfig(config)

	return resourceClient
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

func main() {
	clientset := GetClients("/Users/giorgos.dimitriou/.kube/legacy")
	secretsClient := clientset.CoreV1().Secrets("monitoring")

	s, err := secretsClient.Get(context.TODO(), "logstash-newrelic", metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}
	sec := secret{o: *s}

	r := res{r: &sec}
	r.removeGen()
}
