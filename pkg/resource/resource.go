package resource

import (
	"fmt"

	"github.com/giorgosdi/kubectl-kopy/pkg/secret"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Kind struct {
	// NOTE: Kind is an interface so that resource.go does not need to know what is returning (Secret, Deployment etc)
	Kind   interface{}
	Client kubernetes.Clientset
}

func getClients(conf string) *kubernetes.Clientset {
	config, _ := clientcmd.BuildConfigFromFlags("", conf)
	resourceClient, _ := kubernetes.NewForConfig(config)

	return resourceClient
}

func (k *Kind) GetClientset(kind, kubeconfig string, kcs kubernetes.Clientset) {
	fmt.Println("INSIDE RESOURCE:GETCLIENTSET", kcs)
	clientset := getClients(kubeconfig)
	if kind == "secret" {
		s := &secret.Secret{}
		s.Client = *clientset
		kcs = *clientset
		fmt.Println(kcs)
	}
}
