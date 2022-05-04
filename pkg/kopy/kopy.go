package kopy

import (
	"fmt"

	"github.com/giorgosdi/kubectl-kopy/pkg/options"
	"github.com/giorgosdi/kubectl-kopy/pkg/resource"
)

type ResourceService struct {
	kind       string
	kubeconfig string
}

type KopyResource struct {
	resource kopyResource
}

type kopyResource interface {
	Kopy(target string)
	GetResource(ns, name string)
}

type kopyService struct {
	RS ResourceService
	KS KopyResource
}

func (rS *ResourceService) getClient() resource.Kind {
	return resource.GetClientset(rS.kind, rS.kubeconfig)
}

func (kS *kopyService) RetrieveResource(name, kind, namespace, kubeconfig string) {
	kS.KS.resource = kS.RS.getClient()
	kS.KS.resource.GetResource(namespace, name)
}

func (kService kopyService) Kopy(target string) {
	kService.KS.resource.Kopy(target)
}

func (kService *kopyService) GetResource(ns, name string) {
	kService.KS.resource.GetResource(ns, name)
}

func KopyObject(o *options.KopyOptions) {
	rs := ResourceService{
		kind:       o.Kind,
		kubeconfig: o.Kubeconfig,
	}
	ks := KopyResource{}
	kopyService := kopyService{
		RS: rs,
		KS: ks,
	}
	kopyService.RetrieveResource(o.Name, o.Kind, o.Namespace, o.Kubeconfig)
	kopyService.Kopy(o.Target)
	fmt.Println()
	fmt.Printf("%s %s was copied in %s namespace successfully", o.Name, o.Kind, o.Target)
	fmt.Println()

}
