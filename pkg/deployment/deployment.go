package deployment

import (
	"context"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	appsV1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
)

type Deployment struct {
	O      appsV1.Deployment
	N      appsV1.Deployment
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

func (d *Deployment) RemoveGen() {
	deploymentKopy := d.O.DeepCopy()
	objMeta(&deploymentKopy.ObjectMeta)
	d.N = *deploymentKopy
}

func (d *Deployment) Kopy(target string) {
	d.RemoveGen()
	_, err := d.Client.AppsV1().Deployments(target).Create(context.TODO(), &d.N, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func (d *Deployment) GetResource(ns, secret string) {
	result, _ := d.Client.AppsV1().Deployments(ns).Get(context.TODO(), secret, metav1.GetOptions{})
	d.O = *result
}
