// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	informers "istio.io/client-go/pkg/informers/externalversions/security/v1beta1"
	listers "istio.io/client-go/pkg/listers/security/v1beta1"
	"k8s.io/client-go/tools/cache"
)

type peerAuthenticationInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.PeerAuthenticationInformer = &peerAuthenticationInformer{}

func NewPeerAuthenticationInformer(f xnsinformers.SharedInformerFactory) informers.PeerAuthenticationInformer {
	resource := v1beta1.SchemeGroupVersion.WithResource("peerauthentications")
	converter := xnsinformers.NewListWatchConverter(
		f.GetScheme(),
		&v1beta1.PeerAuthentication{},
		&v1beta1.PeerAuthenticationList{},
	)

	informer := f.ForResource(resource, xnsinformers.ResourceOptions{
		ClusterScoped:      false,
		ListWatchConverter: converter,
	})

	return &peerAuthenticationInformer{informer: informer.Informer()}
}

func (i *peerAuthenticationInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *peerAuthenticationInformer) Lister() listers.PeerAuthenticationLister {
	return listers.NewPeerAuthenticationLister(i.informer.GetIndexer())
}
