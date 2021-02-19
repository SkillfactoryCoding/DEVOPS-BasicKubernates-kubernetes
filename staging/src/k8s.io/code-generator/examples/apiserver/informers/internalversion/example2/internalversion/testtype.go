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

package internalversion

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	example2 "k8s.io/code-generator/examples/apiserver/apis/example2"
	clientsetinternalversion "k8s.io/code-generator/examples/apiserver/clientset/internalversion"
	internalinterfaces "k8s.io/code-generator/examples/apiserver/informers/internalversion/internalinterfaces"
	internalversion "k8s.io/code-generator/examples/apiserver/listers/example2/internalversion"
)

// TestTypeInformer provides access to a shared informer and lister for
// TestTypes.
type TestTypeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.TestTypeLister
}

type testTypeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewTestTypeInformer constructs a new informer for TestType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTestTypeInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTestTypeInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredTestTypeInformer constructs a new informer for TestType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTestTypeInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecondExample().TestTypes().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecondExample().TestTypes().Watch(context.TODO(), options)
			},
		},
		&example2.TestType{},
		resyncPeriod,
		indexers,
	)
}

func (f *testTypeInformer) defaultInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTestTypeInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *testTypeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&example2.TestType{}, f.defaultInformer)
}

func (f *testTypeInformer) Lister() internalversion.TestTypeLister {
	return internalversion.NewTestTypeLister(f.Informer().GetIndexer())
}