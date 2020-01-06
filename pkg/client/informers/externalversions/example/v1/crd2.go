/*
Copyright 2019 JD Cloud

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

package v1

import (
	time "time"

	examplev1 "github.com/sky-big/kubernetes-crd-controller/pkg/apis/example/v1"
	versioned "github.com/sky-big/kubernetes-crd-controller/pkg/client/clientset/versioned"
	internalinterfaces "github.com/sky-big/kubernetes-crd-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/sky-big/kubernetes-crd-controller/pkg/client/listers/example/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CRD2Informer provides access to a shared informer and lister for
// CRD2s.
type CRD2Informer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.CRD2Lister
}

type cRD2Informer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCRD2Informer constructs a new informer for CRD2 type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCRD2Informer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCRD2Informer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCRD2Informer constructs a new informer for CRD2 type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCRD2Informer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleV1().CRD2s(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleV1().CRD2s(namespace).Watch(options)
			},
		},
		&examplev1.CRD2{},
		resyncPeriod,
		indexers,
	)
}

func (f *cRD2Informer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCRD2Informer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cRD2Informer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&examplev1.CRD2{}, f.defaultInformer)
}

func (f *cRD2Informer) Lister() v1.CRD2Lister {
	return v1.NewCRD2Lister(f.Informer().GetIndexer())
}
