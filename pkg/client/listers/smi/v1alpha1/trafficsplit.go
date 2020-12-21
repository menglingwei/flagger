/*
Copyright The Flagger Authors.

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
	v1alpha1 "github.com/fluxcd/flagger/pkg/apis/smi/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TrafficSplitLister helps list TrafficSplits.
type TrafficSplitLister interface {
	// List lists all TrafficSplits in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.TrafficSplit, err error)
	// TrafficSplits returns an object that can list and get TrafficSplits.
	TrafficSplits(namespace string) TrafficSplitNamespaceLister
	TrafficSplitListerExpansion
}

// trafficSplitLister implements the TrafficSplitLister interface.
type trafficSplitLister struct {
	indexer cache.Indexer
}

// NewTrafficSplitLister returns a new TrafficSplitLister.
func NewTrafficSplitLister(indexer cache.Indexer) TrafficSplitLister {
	return &trafficSplitLister{indexer: indexer}
}

// List lists all TrafficSplits in the indexer.
func (s *trafficSplitLister) List(selector labels.Selector) (ret []*v1alpha1.TrafficSplit, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TrafficSplit))
	})
	return ret, err
}

// TrafficSplits returns an object that can list and get TrafficSplits.
func (s *trafficSplitLister) TrafficSplits(namespace string) TrafficSplitNamespaceLister {
	return trafficSplitNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TrafficSplitNamespaceLister helps list and get TrafficSplits.
type TrafficSplitNamespaceLister interface {
	// List lists all TrafficSplits in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.TrafficSplit, err error)
	// Get retrieves the TrafficSplit from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.TrafficSplit, error)
	TrafficSplitNamespaceListerExpansion
}

// trafficSplitNamespaceLister implements the TrafficSplitNamespaceLister
// interface.
type trafficSplitNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TrafficSplits in the indexer for a given namespace.
func (s trafficSplitNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TrafficSplit, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TrafficSplit))
	})
	return ret, err
}

// Get retrieves the TrafficSplit from the indexer for a given namespace and name.
func (s trafficSplitNamespaceLister) Get(name string) (*v1alpha1.TrafficSplit, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("trafficsplit"), name)
	}
	return obj.(*v1alpha1.TrafficSplit), nil
}
