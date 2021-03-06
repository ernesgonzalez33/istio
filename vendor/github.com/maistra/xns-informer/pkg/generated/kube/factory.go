// Code generated by xns-informer-gen. DO NOT EDIT.

package kube

import (
	admissionregistration "github.com/maistra/xns-informer/pkg/generated/kube/admissionregistration"
	apps "github.com/maistra/xns-informer/pkg/generated/kube/apps"
	autoscaling "github.com/maistra/xns-informer/pkg/generated/kube/autoscaling"
	batch "github.com/maistra/xns-informer/pkg/generated/kube/batch"
	certificates "github.com/maistra/xns-informer/pkg/generated/kube/certificates"
	coordination "github.com/maistra/xns-informer/pkg/generated/kube/coordination"
	core "github.com/maistra/xns-informer/pkg/generated/kube/core"
	discovery "github.com/maistra/xns-informer/pkg/generated/kube/discovery"
	events "github.com/maistra/xns-informer/pkg/generated/kube/events"
	extensions "github.com/maistra/xns-informer/pkg/generated/kube/extensions"
	flowcontrol "github.com/maistra/xns-informer/pkg/generated/kube/flowcontrol"
	networking "github.com/maistra/xns-informer/pkg/generated/kube/networking"
	node "github.com/maistra/xns-informer/pkg/generated/kube/node"
	policy "github.com/maistra/xns-informer/pkg/generated/kube/policy"
	rbac "github.com/maistra/xns-informer/pkg/generated/kube/rbac"
	scheduling "github.com/maistra/xns-informer/pkg/generated/kube/scheduling"
	settings "github.com/maistra/xns-informer/pkg/generated/kube/settings"
	storage "github.com/maistra/xns-informer/pkg/generated/kube/storage"
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type SharedInformerFactory interface {
	Admissionregistration() admissionregistration.Interface
	Apps() apps.Interface
	Autoscaling() autoscaling.Interface
	Batch() batch.Interface
	Certificates() certificates.Interface
	Coordination() coordination.Interface
	Core() core.Interface
	Discovery() discovery.Interface
	Events() events.Interface
	Extensions() extensions.Interface
	Flowcontrol() flowcontrol.Interface
	Networking() networking.Interface
	Node() node.Interface
	Policy() policy.Interface
	Rbac() rbac.Interface
	Scheduling() scheduling.Interface
	Settings() settings.Interface
	Storage() storage.Interface
	ForResource(resource schema.GroupVersionResource) (GenericInformer, error)
}

type sharedInformerFactory struct {
	factory xnsinformers.SharedInformerFactory
}

// NewSharedInformerFactory returns a new SharedInformerFactory.
func NewSharedInformerFactory(f xnsinformers.SharedInformerFactory) SharedInformerFactory {
	return &sharedInformerFactory{factory: f}
}

// Admissionregistration returns a new admissionregistration.Interface.
func (f *sharedInformerFactory) Admissionregistration() admissionregistration.Interface {
	return admissionregistration.New(f.factory)
}

// Apps returns a new apps.Interface.
func (f *sharedInformerFactory) Apps() apps.Interface {
	return apps.New(f.factory)
}

// Autoscaling returns a new autoscaling.Interface.
func (f *sharedInformerFactory) Autoscaling() autoscaling.Interface {
	return autoscaling.New(f.factory)
}

// Batch returns a new batch.Interface.
func (f *sharedInformerFactory) Batch() batch.Interface {
	return batch.New(f.factory)
}

// Certificates returns a new certificates.Interface.
func (f *sharedInformerFactory) Certificates() certificates.Interface {
	return certificates.New(f.factory)
}

// Coordination returns a new coordination.Interface.
func (f *sharedInformerFactory) Coordination() coordination.Interface {
	return coordination.New(f.factory)
}

// Core returns a new core.Interface.
func (f *sharedInformerFactory) Core() core.Interface {
	return core.New(f.factory)
}

// Discovery returns a new discovery.Interface.
func (f *sharedInformerFactory) Discovery() discovery.Interface {
	return discovery.New(f.factory)
}

// Events returns a new events.Interface.
func (f *sharedInformerFactory) Events() events.Interface {
	return events.New(f.factory)
}

// Extensions returns a new extensions.Interface.
func (f *sharedInformerFactory) Extensions() extensions.Interface {
	return extensions.New(f.factory)
}

// Flowcontrol returns a new flowcontrol.Interface.
func (f *sharedInformerFactory) Flowcontrol() flowcontrol.Interface {
	return flowcontrol.New(f.factory)
}

// Networking returns a new networking.Interface.
func (f *sharedInformerFactory) Networking() networking.Interface {
	return networking.New(f.factory)
}

// Node returns a new node.Interface.
func (f *sharedInformerFactory) Node() node.Interface {
	return node.New(f.factory)
}

// Policy returns a new policy.Interface.
func (f *sharedInformerFactory) Policy() policy.Interface {
	return policy.New(f.factory)
}

// Rbac returns a new rbac.Interface.
func (f *sharedInformerFactory) Rbac() rbac.Interface {
	return rbac.New(f.factory)
}

// Scheduling returns a new scheduling.Interface.
func (f *sharedInformerFactory) Scheduling() scheduling.Interface {
	return scheduling.New(f.factory)
}

// Settings returns a new settings.Interface.
func (f *sharedInformerFactory) Settings() settings.Interface {
	return settings.New(f.factory)
}

// Storage returns a new storage.Interface.
func (f *sharedInformerFactory) Storage() storage.Interface {
	return storage.New(f.factory)
}
