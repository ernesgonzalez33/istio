// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	informers "k8s.io/client-go/informers/policy/v1beta1"
)

type Interface interface {
	PodDisruptionBudgets() informers.PodDisruptionBudgetInformer
	PodSecurityPolicies() informers.PodSecurityPolicyInformer
}

type version struct {
	factory xnsinformers.SharedInformerFactory
}

func New(factory xnsinformers.SharedInformerFactory) Interface {
	return &version{factory: factory}
}
func (v *version) PodDisruptionBudgets() informers.PodDisruptionBudgetInformer {
	return NewPodDisruptionBudgetInformer(v.factory)
}
func (v *version) PodSecurityPolicies() informers.PodSecurityPolicyInformer {
	return NewPodSecurityPolicyInformer(v.factory)
}
