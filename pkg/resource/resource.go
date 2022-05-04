package resource

import (
	"github.com/giorgosdi/kubectl-kopy/pkg/deployment"
	"github.com/giorgosdi/kubectl-kopy/pkg/secret"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Kind interface {
	// NOTE: Kind is an interface so that resource.go does not need to know what is returning (Secret, Deployment etc)
	Kopy(target string)
	GetResource(ns, kind string)
}

func getClients(conf string) *kubernetes.Clientset {
	config, _ := clientcmd.BuildConfigFromFlags("", conf)
	resourceClient, _ := kubernetes.NewForConfig(config)

	return resourceClient
}

func GetClientset(kind, kubeconfig string) Kind {
	clientset := getClients(kubeconfig)
	if kind == "secret" {
		return &secret.Secret{
			Client: *clientset,
		}
	}
	if kind == "deployment" {
		return &deployment.Deployment{
			Client: *clientset,
		}
	}
	return nil
}
