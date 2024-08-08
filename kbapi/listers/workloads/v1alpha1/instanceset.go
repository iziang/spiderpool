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
	v1alpha1 "github.com/spidernet-io/spiderpool/kbapi/workloads/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InstanceSetLister helps list InstanceSets.
// All objects returned here must be treated as read-only.
type InstanceSetLister interface {
	// List lists all InstanceSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InstanceSet, err error)
	// InstanceSets returns an object that can list and get InstanceSets.
	InstanceSets(namespace string) InstanceSetNamespaceLister
	InstanceSetListerExpansion
}

// instanceSetLister implements the InstanceSetLister interface.
type instanceSetLister struct {
	indexer cache.Indexer
}

// NewInstanceSetLister returns a new InstanceSetLister.
func NewInstanceSetLister(indexer cache.Indexer) InstanceSetLister {
	return &instanceSetLister{indexer: indexer}
}

// List lists all InstanceSets in the indexer.
func (s *instanceSetLister) List(selector labels.Selector) (ret []*v1alpha1.InstanceSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InstanceSet))
	})
	return ret, err
}

// InstanceSets returns an object that can list and get InstanceSets.
func (s *instanceSetLister) InstanceSets(namespace string) InstanceSetNamespaceLister {
	return instanceSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InstanceSetNamespaceLister helps list and get InstanceSets.
// All objects returned here must be treated as read-only.
type InstanceSetNamespaceLister interface {
	// List lists all InstanceSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InstanceSet, err error)
	// Get retrieves the InstanceSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.InstanceSet, error)
	InstanceSetNamespaceListerExpansion
}

// instanceSetNamespaceLister implements the InstanceSetNamespaceLister
// interface.
type instanceSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all InstanceSets in the indexer for a given namespace.
func (s instanceSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.InstanceSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InstanceSet))
	})
	return ret, err
}

// Get retrieves the InstanceSet from the indexer for a given namespace and name.
func (s instanceSetNamespaceLister) Get(name string) (*v1alpha1.InstanceSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("instanceset"), name)
	}
	return obj.(*v1alpha1.InstanceSet), nil
}
