// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1beta1 "k8s.io/api/extensions/v1beta1"
	informers "k8s.io/client-go/informers/extensions/v1beta1"
	listers "k8s.io/client-go/listers/extensions/v1beta1"
	"k8s.io/client-go/tools/cache"
)

type networkPolicyInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.NetworkPolicyInformer = &networkPolicyInformer{}

func NewNetworkPolicyInformer(f xnsinformers.SharedInformerFactory) informers.NetworkPolicyInformer {
	resource := v1beta1.SchemeGroupVersion.WithResource("networkpolicies")
	converter := xnsinformers.NewListWatchConverter(
		f.GetScheme(),
		&v1beta1.NetworkPolicy{},
		&v1beta1.NetworkPolicyList{},
	)

	informer := f.ForResource(resource, xnsinformers.ResourceOptions{
		ClusterScoped:      false,
		ListWatchConverter: converter,
	})

	return &networkPolicyInformer{informer: informer.Informer()}
}

func (i *networkPolicyInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *networkPolicyInformer) Lister() listers.NetworkPolicyLister {
	return listers.NewNetworkPolicyLister(i.informer.GetIndexer())
}
