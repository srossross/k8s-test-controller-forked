/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/munnerz/k8s-api-pager-demo/pkg/apis/pager/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TestRunnerLister helps list TestRunners.
type TestRunnerLister interface {
	// List lists all TestRunners in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.TestRunner, err error)
	// TestRunners returns an object that can list and get TestRunners.
	TestRunners(namespace string) TestRunnerNamespaceLister
	TestRunnerListerExpansion
}

// testRunnerLister implements the TestRunnerLister interface.
type testRunnerLister struct {
	indexer cache.Indexer
}

// NewTestRunnerLister returns a new TestRunnerLister.
func NewTestRunnerLister(indexer cache.Indexer) TestRunnerLister {
	return &testRunnerLister{indexer: indexer}
}

// List lists all TestRunners in the indexer.
func (s *testRunnerLister) List(selector labels.Selector) (ret []*v1alpha1.TestRunner, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TestRunner))
	})
	return ret, err
}

// TestRunners returns an object that can list and get TestRunners.
func (s *testRunnerLister) TestRunners(namespace string) TestRunnerNamespaceLister {
	return testRunnerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TestRunnerNamespaceLister helps list and get TestRunners.
type TestRunnerNamespaceLister interface {
	// List lists all TestRunners in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.TestRunner, err error)
	// Get retrieves the TestRunner from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.TestRunner, error)
	TestRunnerNamespaceListerExpansion
}

// testRunnerNamespaceLister implements the TestRunnerNamespaceLister
// interface.
type testRunnerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TestRunners in the indexer for a given namespace.
func (s testRunnerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TestRunner, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TestRunner))
	})
	return ret, err
}

// Get retrieves the TestRunner from the indexer for a given namespace and name.
func (s testRunnerNamespaceLister) Get(name string) (*v1alpha1.TestRunner, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("testrunner"), name)
	}
	return obj.(*v1alpha1.TestRunner), nil
}
