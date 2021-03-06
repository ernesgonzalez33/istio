// Code generated by xns-informer-gen. DO NOT EDIT.

package v1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1 "k8s.io/api/core/v1"
	informers "k8s.io/client-go/informers/core/v1"
	listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
)

type podTemplateInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.PodTemplateInformer = &podTemplateInformer{}

func NewPodTemplateInformer(f xnsinformers.SharedInformerFactory) informers.PodTemplateInformer {
	resource := v1.SchemeGroupVersion.WithResource("podtemplates")
	converter := xnsinformers.NewListWatchConverter(
		f.GetScheme(),
		&v1.PodTemplate{},
		&v1.PodTemplateList{},
	)

	informer := f.ForResource(resource, xnsinformers.ResourceOptions{
		ClusterScoped:      false,
		ListWatchConverter: converter,
	})

	return &podTemplateInformer{informer: informer.Informer()}
}

func (i *podTemplateInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *podTemplateInformer) Lister() listers.PodTemplateLister {
	return listers.NewPodTemplateLister(i.informer.GetIndexer())
}
