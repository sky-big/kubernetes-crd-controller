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

package v1alpha1

import (
	time "time"

	examplev1alpha1 "github.com/sky-big/kubernetes-crd-controller/pkg/apis/example/v1alpha1"
	versioned "github.com/sky-big/kubernetes-crd-controller/pkg/client/clientset/versioned"
	internalinterfaces "github.com/sky-big/kubernetes-crd-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/sky-big/kubernetes-crd-controller/pkg/client/listers/example/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CRD1Informer provides access to a shared informer and lister for
// CRD1s.
type CRD1Informer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CRD1Lister
}

type cRD1Informer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCRD1Informer constructs a new informer for CRD1 type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCRD1Informer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCRD1Informer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCRD1Informer constructs a new informer for CRD1 type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCRD1Informer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleV1alpha1().CRD1s(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleV1alpha1().CRD1s(namespace).Watch(options)
			},
		},
		&examplev1alpha1.CRD1{},
		resyncPeriod,
		indexers,
	)
}

func (f *cRD1Informer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCRD1Informer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cRD1Informer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&examplev1alpha1.CRD1{}, f.defaultInformer)
}

func (f *cRD1Informer) Lister() v1alpha1.CRD1Lister {
	return v1alpha1.NewCRD1Lister(f.Informer().GetIndexer())
}
