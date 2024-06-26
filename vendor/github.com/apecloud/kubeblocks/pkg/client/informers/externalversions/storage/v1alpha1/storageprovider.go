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

	storagev1alpha1 "github.com/apecloud/kubeblocks/apis/storage/v1alpha1"
	versioned "github.com/apecloud/kubeblocks/pkg/client/clientset/versioned"
	internalinterfaces "github.com/apecloud/kubeblocks/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/apecloud/kubeblocks/pkg/client/listers/storage/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// StorageProviderInformer provides access to a shared informer and lister for
// StorageProviders.
type StorageProviderInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.StorageProviderLister
}

type storageProviderInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewStorageProviderInformer constructs a new informer for StorageProvider type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewStorageProviderInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredStorageProviderInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredStorageProviderInformer constructs a new informer for StorageProvider type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredStorageProviderInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1alpha1().StorageProviders().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1alpha1().StorageProviders().Watch(context.TODO(), options)
			},
		},
		&storagev1alpha1.StorageProvider{},
		resyncPeriod,
		indexers,
	)
}

func (f *storageProviderInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredStorageProviderInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *storageProviderInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&storagev1alpha1.StorageProvider{}, f.defaultInformer)
}

func (f *storageProviderInformer) Lister() v1alpha1.StorageProviderLister {
	return v1alpha1.NewStorageProviderLister(f.Informer().GetIndexer())
}
