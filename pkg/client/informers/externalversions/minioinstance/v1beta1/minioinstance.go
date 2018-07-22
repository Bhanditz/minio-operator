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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	time "time"

	minioinstancev1beta1 "github.com/nitisht/minio-operator/pkg/apis/minioinstance/v1beta1"
	versioned "github.com/nitisht/minio-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/nitisht/minio-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/nitisht/minio-operator/pkg/client/listers/minioinstance/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MinioInstanceInformer provides access to a shared informer and lister for
// MinioInstances.
type MinioInstanceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.MinioInstanceLister
}

type minioInstanceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMinioInstanceInformer constructs a new informer for MinioInstance type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMinioInstanceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMinioInstanceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMinioInstanceInformer constructs a new informer for MinioInstance type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMinioInstanceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MinioV1beta1().MinioInstances(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MinioV1beta1().MinioInstances(namespace).Watch(options)
			},
		},
		&minioinstancev1beta1.MinioInstance{},
		resyncPeriod,
		indexers,
	)
}

func (f *minioInstanceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMinioInstanceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *minioInstanceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&minioinstancev1beta1.MinioInstance{}, f.defaultInformer)
}

func (f *minioInstanceInformer) Lister() v1beta1.MinioInstanceLister {
	return v1beta1.NewMinioInstanceLister(f.Informer().GetIndexer())
}
