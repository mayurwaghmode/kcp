/*
Copyright The KCP Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"

	workloadv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/workload/v1alpha1"
	versioned "github.com/kcp-dev/kcp/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kcp-dev/kcp/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kcp-dev/kcp/pkg/client/listers/workload/v1alpha1"
)

// WorkloadClusterInformer provides access to a shared informer and lister for
// WorkloadClusters.
type WorkloadClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.WorkloadClusterLister
}

type workloadClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewWorkloadClusterInformer constructs a new informer for WorkloadCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewWorkloadClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredWorkloadClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredWorkloadClusterInformer constructs a new informer for WorkloadCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredWorkloadClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return NewFilteredWorkloadClusterInformerWithOptions(client, tweakListOptions, cache.WithResyncPeriod(resyncPeriod), cache.WithIndexers(indexers))
}

func NewFilteredWorkloadClusterInformerWithOptions(client versioned.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc, opts ...cache.SharedInformerOption) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformerWithOptions(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WorkloadV1alpha1().WorkloadClusters().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WorkloadV1alpha1().WorkloadClusters().Watch(context.TODO(), options)
			},
		},
		&workloadv1alpha1.WorkloadCluster{},
		opts...,
	)
}

func (f *workloadClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	indexers := cache.Indexers{}
	for k, v := range f.factory.ExtraClusterScopedIndexers() {
		indexers[k] = v
	}

	return NewFilteredWorkloadClusterInformerWithOptions(client,
		f.tweakListOptions,
		cache.WithResyncPeriod(resyncPeriod),
		cache.WithIndexers(indexers),
		cache.WithKeyFunction(f.factory.KeyFunction()),
	)
}

func (f *workloadClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&workloadv1alpha1.WorkloadCluster{}, f.defaultInformer)
}

func (f *workloadClusterInformer) Lister() v1alpha1.WorkloadClusterLister {
	return v1alpha1.NewWorkloadClusterLister(f.Informer().GetIndexer())
}
