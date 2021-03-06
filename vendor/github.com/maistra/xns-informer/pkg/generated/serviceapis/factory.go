// Code generated by xns-informer-gen. DO NOT EDIT.

package serviceapis

import (
	apis "github.com/maistra/xns-informer/pkg/generated/serviceapis/apis"
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type SharedInformerFactory interface {
	Apis() apis.Interface
	ForResource(resource schema.GroupVersionResource) (GenericInformer, error)
}

type sharedInformerFactory struct {
	factory xnsinformers.SharedInformerFactory
}

// NewSharedInformerFactory returns a new SharedInformerFactory.
func NewSharedInformerFactory(f xnsinformers.SharedInformerFactory) SharedInformerFactory {
	return &sharedInformerFactory{factory: f}
}

// Apis returns a new apis.Interface.
func (f *sharedInformerFactory) Apis() apis.Interface {
	return apis.New(f.factory)
}
