/*
Copyright (C) 2022-2024 ApeCloud Co., Ltd

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
	v1alpha1 "github.com/apecloud/kubeblocks/apis/apps/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServiceDescriptorLister helps list ServiceDescriptors.
// All objects returned here must be treated as read-only.
type ServiceDescriptorLister interface {
	// List lists all ServiceDescriptors in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceDescriptor, err error)
	// ServiceDescriptors returns an object that can list and get ServiceDescriptors.
	ServiceDescriptors(namespace string) ServiceDescriptorNamespaceLister
	ServiceDescriptorListerExpansion
}

// serviceDescriptorLister implements the ServiceDescriptorLister interface.
type serviceDescriptorLister struct {
	indexer cache.Indexer
}

// NewServiceDescriptorLister returns a new ServiceDescriptorLister.
func NewServiceDescriptorLister(indexer cache.Indexer) ServiceDescriptorLister {
	return &serviceDescriptorLister{indexer: indexer}
}

// List lists all ServiceDescriptors in the indexer.
func (s *serviceDescriptorLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceDescriptor, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceDescriptor))
	})
	return ret, err
}

// ServiceDescriptors returns an object that can list and get ServiceDescriptors.
func (s *serviceDescriptorLister) ServiceDescriptors(namespace string) ServiceDescriptorNamespaceLister {
	return serviceDescriptorNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServiceDescriptorNamespaceLister helps list and get ServiceDescriptors.
// All objects returned here must be treated as read-only.
type ServiceDescriptorNamespaceLister interface {
	// List lists all ServiceDescriptors in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceDescriptor, err error)
	// Get retrieves the ServiceDescriptor from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ServiceDescriptor, error)
	ServiceDescriptorNamespaceListerExpansion
}

// serviceDescriptorNamespaceLister implements the ServiceDescriptorNamespaceLister
// interface.
type serviceDescriptorNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServiceDescriptors in the indexer for a given namespace.
func (s serviceDescriptorNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceDescriptor, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceDescriptor))
	})
	return ret, err
}

// Get retrieves the ServiceDescriptor from the indexer for a given namespace and name.
func (s serviceDescriptorNamespaceLister) Get(name string) (*v1alpha1.ServiceDescriptor, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("servicedescriptor"), name)
	}
	return obj.(*v1alpha1.ServiceDescriptor), nil
}
