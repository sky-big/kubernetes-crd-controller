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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/sky-big/kubernetes-crd-controller/pkg/apis/example/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CRD1Lister helps list CRD1s.
type CRD1Lister interface {
	// List lists all CRD1s in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CRD1, err error)
	// CRD1s returns an object that can list and get CRD1s.
	CRD1s(namespace string) CRD1NamespaceLister
	CRD1ListerExpansion
}

// cRD1Lister implements the CRD1Lister interface.
type cRD1Lister struct {
	indexer cache.Indexer
}

// NewCRD1Lister returns a new CRD1Lister.
func NewCRD1Lister(indexer cache.Indexer) CRD1Lister {
	return &cRD1Lister{indexer: indexer}
}

// List lists all CRD1s in the indexer.
func (s *cRD1Lister) List(selector labels.Selector) (ret []*v1alpha1.CRD1, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CRD1))
	})
	return ret, err
}

// CRD1s returns an object that can list and get CRD1s.
func (s *cRD1Lister) CRD1s(namespace string) CRD1NamespaceLister {
	return cRD1NamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CRD1NamespaceLister helps list and get CRD1s.
type CRD1NamespaceLister interface {
	// List lists all CRD1s in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CRD1, err error)
	// Get retrieves the CRD1 from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CRD1, error)
	CRD1NamespaceListerExpansion
}

// cRD1NamespaceLister implements the CRD1NamespaceLister
// interface.
type cRD1NamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CRD1s in the indexer for a given namespace.
func (s cRD1NamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CRD1, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CRD1))
	})
	return ret, err
}

// Get retrieves the CRD1 from the indexer for a given namespace and name.
func (s cRD1NamespaceLister) Get(name string) (*v1alpha1.CRD1, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("crd1"), name)
	}
	return obj.(*v1alpha1.CRD1), nil
}
