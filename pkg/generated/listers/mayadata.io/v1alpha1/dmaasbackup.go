/*
Copyright 2020 The MayaData Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    https://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/mayadata-io/dmaas-operator/pkg/apis/mayadata.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DMaaSBackupLister helps list DMaaSBackups.
type DMaaSBackupLister interface {
	// List lists all DMaaSBackups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DMaaSBackup, err error)
	// DMaaSBackups returns an object that can list and get DMaaSBackups.
	DMaaSBackups(namespace string) DMaaSBackupNamespaceLister
	DMaaSBackupListerExpansion
}

// dMaaSBackupLister implements the DMaaSBackupLister interface.
type dMaaSBackupLister struct {
	indexer cache.Indexer
}

// NewDMaaSBackupLister returns a new DMaaSBackupLister.
func NewDMaaSBackupLister(indexer cache.Indexer) DMaaSBackupLister {
	return &dMaaSBackupLister{indexer: indexer}
}

// List lists all DMaaSBackups in the indexer.
func (s *dMaaSBackupLister) List(selector labels.Selector) (ret []*v1alpha1.DMaaSBackup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DMaaSBackup))
	})
	return ret, err
}

// DMaaSBackups returns an object that can list and get DMaaSBackups.
func (s *dMaaSBackupLister) DMaaSBackups(namespace string) DMaaSBackupNamespaceLister {
	return dMaaSBackupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DMaaSBackupNamespaceLister helps list and get DMaaSBackups.
type DMaaSBackupNamespaceLister interface {
	// List lists all DMaaSBackups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DMaaSBackup, err error)
	// Get retrieves the DMaaSBackup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DMaaSBackup, error)
	DMaaSBackupNamespaceListerExpansion
}

// dMaaSBackupNamespaceLister implements the DMaaSBackupNamespaceLister
// interface.
type dMaaSBackupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DMaaSBackups in the indexer for a given namespace.
func (s dMaaSBackupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DMaaSBackup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DMaaSBackup))
	})
	return ret, err
}

// Get retrieves the DMaaSBackup from the indexer for a given namespace and name.
func (s dMaaSBackupNamespaceLister) Get(name string) (*v1alpha1.DMaaSBackup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dmaasbackup"), name)
	}
	return obj.(*v1alpha1.DMaaSBackup), nil
}