/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "k8s.io/api/settings/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PodPresetLister helps list PodPresets.
type PodPresetLister interface {
	// List lists all PodPresets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PodPreset, err error)
	// PodPresets returns an object that can list and get PodPresets.
	PodPresets(namespace string) PodPresetNamespaceLister
	PodPresetsWithMultiTenancy(namespace string, tenant string) PodPresetNamespaceLister
	PodPresetListerExpansion
}

// podPresetLister implements the PodPresetLister interface.
type podPresetLister struct {
	indexer cache.Indexer
}

// NewPodPresetLister returns a new PodPresetLister.
func NewPodPresetLister(indexer cache.Indexer) PodPresetLister {
	return &podPresetLister{indexer: indexer}
}

// List lists all PodPresets in the indexer.
func (s *podPresetLister) List(selector labels.Selector) (ret []*v1alpha1.PodPreset, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PodPreset))
	})
	return ret, err
}

// PodPresets returns an object that can list and get PodPresets.
func (s *podPresetLister) PodPresets(namespace string) PodPresetNamespaceLister {
	return podPresetNamespaceLister{indexer: s.indexer, namespace: namespace, tenant: "default"}
}

func (s *podPresetLister) PodPresetsWithMultiTenancy(namespace string, tenant string) PodPresetNamespaceLister {
	return podPresetNamespaceLister{indexer: s.indexer, namespace: namespace, tenant: tenant}
}

// PodPresetNamespaceLister helps list and get PodPresets.
type PodPresetNamespaceLister interface {
	// List lists all PodPresets in the indexer for a given tenant/namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PodPreset, err error)
	// Get retrieves the PodPreset from the indexer for a given tenant/namespace and name.
	Get(name string) (*v1alpha1.PodPreset, error)
	PodPresetNamespaceListerExpansion
}

// podPresetNamespaceLister implements the PodPresetNamespaceLister
// interface.
type podPresetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
	tenant    string
}

// List lists all PodPresets in the indexer for a given namespace.
func (s podPresetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PodPreset, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.tenant, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PodPreset))
	})
	return ret, err
}

// Get retrieves the PodPreset from the indexer for a given namespace and name.
func (s podPresetNamespaceLister) Get(name string) (*v1alpha1.PodPreset, error) {
	key := s.tenant + "/" + s.namespace + "/" + name
	if s.tenant == "default" {
		key = s.namespace + "/" + name
	}
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("podpreset"), name)
	}
	return obj.(*v1alpha1.PodPreset), nil
}
