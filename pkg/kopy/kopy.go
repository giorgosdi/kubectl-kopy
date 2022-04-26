package kopy

import (
	"fmt"

	"github.com/giorgosdi/kubectl-kopy/pkg/options"
	"github.com/giorgosdi/kubectl-kopy/pkg/resource"
	"github.com/giorgosdi/kubectl-kopy/pkg/secret"
	"k8s.io/client-go/kubernetes"
)

type ResourceKind interface {
	GetClientset(kind, kubeconfig string, kcs kubernetes.Clientset)
}

type ResourceService struct {
	kind   ResourceKind
	Client kubernetes.Clientset
}

//NOTE: This is not really used
type kopyResource interface {
	Kopy()
	GetResource(ns, name string)
}

type KopyResource struct {
	resource kopyResource
}

type kopyService struct {
	RS ResourceService
	KS KopyResource
}

func (rS *ResourceService) GetClient(kind, kubeconfig string) {
	rS.kind.GetClientset(kind, kubeconfig, rS.Client)
	fmt.Println("INSIDE KOPY:GETCLIENT", rS.Client)
}

func (kS *kopyService) RetrieveResource(name, kind, namespace, kubeconfig string) {
	fmt.Println(name)
	fmt.Println(kind)
	fmt.Println(namespace)
	fmt.Println(kubeconfig)
	kS.RS.GetClient(kind, kubeconfig)
	//kS.RS.kind.GetClientset(kind, kubeconfig)
	fmt.Println("After kind")
	kS.KS.resource = &secret.Secret{
		Client: kS.RS.Client,
	}
	kS.KS.resource.GetResource(namespace, name)
}

func (kService kopyService) Kopy() {
	kService.KS.resource.Kopy()
}

func (kService *kopyService) GetResource(ns, name string) {
	fmt.Println("IM HERE")
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
	kopyService.Kopy()

}
