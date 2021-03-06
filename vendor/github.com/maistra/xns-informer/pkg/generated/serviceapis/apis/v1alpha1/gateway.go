// Code generated by xns-informer-gen. DO NOT EDIT.

package v1alpha1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "sigs.k8s.io/service-apis/apis/v1alpha1"
	informers "sigs.k8s.io/service-apis/pkg/client/informers/externalversions/apis/v1alpha1"
	listers "sigs.k8s.io/service-apis/pkg/client/listers/apis/v1alpha1"
)

type gatewayInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.GatewayInformer = &gatewayInformer{}

func NewGatewayInformer(f xnsinformers.SharedInformerFactory) informers.GatewayInformer {
	resource := v1alpha1.SchemeGroupVersion.WithResource("gateways")
	converter := xnsinformers.NewListWatchConverter(
		f.GetScheme(),
		&v1alpha1.Gateway{},
		&v1alpha1.GatewayList{},
	)

	informer := f.ForResource(resource, xnsinformers.ResourceOptions{
		ClusterScoped:      false,
		ListWatchConverter: converter,
	})

	return &gatewayInformer{informer: informer.Informer()}
}

func (i *gatewayInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *gatewayInformer) Lister() listers.GatewayLister {
	return listers.NewGatewayLister(i.informer.GetIndexer())
}
