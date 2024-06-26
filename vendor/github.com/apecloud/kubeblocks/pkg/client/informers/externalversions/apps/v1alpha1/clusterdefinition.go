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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	appsv1alpha1 "github.com/apecloud/kubeblocks/apis/apps/v1alpha1"
	versioned "github.com/apecloud/kubeblocks/pkg/client/clientset/versioned"
	internalinterfaces "github.com/apecloud/kubeblocks/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/apecloud/kubeblocks/pkg/client/listers/apps/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterDefinitionInformer provides access to a shared informer and lister for
// ClusterDefinitions.
type ClusterDefinitionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterDefinitionLister
}

type clusterDefinitionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterDefinitionInformer constructs a new informer for ClusterDefinition type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterDefinitionInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterDefinitionInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterDefinitionInformer constructs a new informer for ClusterDefinition type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterDefinitionInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1alpha1().ClusterDefinitions().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1alpha1().ClusterDefinitions().Watch(context.TODO(), options)
			},
		},
		&appsv1alpha1.ClusterDefinition{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterDefinitionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterDefinitionInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterDefinitionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&appsv1alpha1.ClusterDefinition{}, f.defaultInformer)
}

func (f *clusterDefinitionInformer) Lister() v1alpha1.ClusterDefinitionLister {
	return v1alpha1.NewClusterDefinitionLister(f.Informer().GetIndexer())
}
