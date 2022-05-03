package kopy

import (
	"github.com/giorgosdi/kubectl-kopy/pkg/options"
	"github.com/giorgosdi/kubectl-kopy/pkg/resource"
	"github.com/giorgosdi/kubectl-kopy/pkg/secret"
	"k8s.io/client-go/kubernetes"
)

type ResourceKind interface {
	GetClientset(kind, kubeconfig string) *kubernetes.Clientset
}

type ResourceService struct {
	kind ResourceKind
}

//NOTE: This is not really used
type kopyResource interface {
	Kopy(target string)
	GetResource(ns, name string)
}

type KopyResource struct {
	resource kopyResource
}

type kopyService struct {
	RS ResourceService
	KS KopyResource
}

func (rS *ResourceService) GetClient(kind, kubeconfig string) *kubernetes.Clientset {
	client := rS.kind.GetClientset(kind, kubeconfig)
	return client
}

func (kS *kopyService) RetrieveResource(name, kind, namespace, kubeconfig string) {
	client := kS.RS.GetClient(kind, kubeconfig)
	kS.KS.resource = &secret.Secret{
		Client: *client,
	}
	kS.KS.resource.GetResource(namespace, name)
}

func (kService kopyService) Kopy(target string) {
	kService.KS.resource.Kopy(target)
}

func (kService *kopyService) GetResource(ns, name string) {
	kService.KS.resource.GetResource(ns, name)
}

func KopyObject(o *options.KopyOptions) {
	rs := ResourceService{}
	rs.kind = &resource.Kind{}
	ks := KopyResource{}
	kopyService := kopyService{
		RS: rs,
		KS: ks,
	}
	kopyService.RetrieveResource(o.Name, o.Kind, o.Namespace, o.Kubeconfig)
	kopyService.Kopy(o.Target)

}
