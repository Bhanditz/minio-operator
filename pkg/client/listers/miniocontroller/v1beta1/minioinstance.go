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

package v1beta1

import (
	v1beta1 "github.com/nitisht/minio-operator/pkg/apis/miniocontroller/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MinioInstanceLister helps list MinioInstances.
type MinioInstanceLister interface {
	// List lists all MinioInstances in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.MinioInstance, err error)
	// MinioInstances returns an object that can list and get MinioInstances.
	MinioInstances(namespace string) MinioInstanceNamespaceLister
	MinioInstanceListerExpansion
}

// minioInstanceLister implements the MinioInstanceLister interface.
type minioInstanceLister struct {
	indexer cache.Indexer
}

// NewMinioInstanceLister returns a new MinioInstanceLister.
func NewMinioInstanceLister(indexer cache.Indexer) MinioInstanceLister {
	return &minioInstanceLister{indexer: indexer}
}

// List lists all MinioInstances in the indexer.
func (s *minioInstanceLister) List(selector labels.Selector) (ret []*v1beta1.MinioInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.MinioInstance))
	})
	return ret, err
}

// MinioInstances returns an object that can list and get MinioInstances.
func (s *minioInstanceLister) MinioInstances(namespace string) MinioInstanceNamespaceLister {
	return minioInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MinioInstanceNamespaceLister helps list and get MinioInstances.
type MinioInstanceNamespaceLister interface {
	// List lists all MinioInstances in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.MinioInstance, err error)
	// Get retrieves the MinioInstance from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.MinioInstance, error)
	MinioInstanceNamespaceListerExpansion
}

// minioInstanceNamespaceLister implements the MinioInstanceNamespaceLister
// interface.
type minioInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MinioInstances in the indexer for a given namespace.
func (s minioInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.MinioInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.MinioInstance))
	})
	return ret, err
}

// Get retrieves the MinioInstance from the indexer for a given namespace and name.
func (s minioInstanceNamespaceLister) Get(name string) (*v1beta1.MinioInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("minioinstance"), name)
	}
	return obj.(*v1beta1.MinioInstance), nil
}
