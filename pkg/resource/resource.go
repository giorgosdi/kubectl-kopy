package resource

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Kind struct {
	// NOTE: Kind is an interface so that resource.go does not need to know what is returning (Secret, Deployment etc)
	Kind interface{}
}

func getClients(conf string) *kubernetes.Clientset {
	config, _ := clientcmd.BuildConfigFromFlags("", conf)
	resourceClient, _ := kubernetes.NewForConfig(config)

	return resourceClient
}

func (k *Kind) GetClientset(kind, kubeconfig string) *kubernetes.Clientset {
	clientset := getClients(kubeconfig)
	return clientset
}
