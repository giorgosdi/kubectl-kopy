package main

import (
	"github.com/giorgosdi/kubectl-kopy/pkg/secret"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type resource interface {
	Kopy()
}

type res struct {
	r resource
}

func (resource res) Kopy() {
	resource.r.Kopy()
}

func GetClients(conf string) *kubernetes.Clientset {
	config, _ := clientcmd.BuildConfigFromFlags("", conf)
	resourceClient, _ := kubernetes.NewForConfig(config)

	return resourceClient
}

func main() {
	clientset := GetClients("/Users/giorgos.dimitriou/.kube/legacy")
	s := secret.Secret{}
	s.Client = *clientset

	result := s.GetSecret("monitoring", "logstash-newrelic")
	s.O = *result
	rs := res{r: &s}
	rs.Kopy()
}
