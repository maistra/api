// Copyright Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by xns-informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	informers "github.com/maistra/xns-informer/pkg/informers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	v1 "maistra.io/api/client/listers/federation/v1"
	versioned "maistra.io/api/client/versioned"
	internalinterfaces "maistra.io/api/client/xnsinformer/internalinterfaces"
	federationv1 "maistra.io/api/federation/v1"
)

// ServiceMeshPeerInformer provides access to a shared informer and lister for
// ServiceMeshPeers.
type ServiceMeshPeerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ServiceMeshPeerLister
}

type serviceMeshPeerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespaces       informers.NamespaceSet
}

// NewServiceMeshPeerInformer constructs a new informer for ServiceMeshPeer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceMeshPeerInformer(client versioned.Interface, namespaces informers.NamespaceSet,
	resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServiceMeshPeerInformer(client, namespaces, resyncPeriod, indexers, nil)
}

// NewFilteredServiceMeshPeerInformer constructs a new informer for ServiceMeshPeer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceMeshPeerInformer(client versioned.Interface, namespaces informers.NamespaceSet,
	resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	newInformer := func(namespace string) cache.SharedIndexInformer {
		return cache.NewSharedIndexInformer(
			&cache.ListWatch{
				ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
					if tweakListOptions != nil {
						tweakListOptions(&options)
					}
					return client.FederationV1().ServiceMeshPeers(namespace).List(context.TODO(), options)
				},
				WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
					if tweakListOptions != nil {
						tweakListOptions(&options)
					}
					return client.FederationV1().ServiceMeshPeers(namespace).Watch(context.TODO(), options)
				},
			},
			&federationv1.ServiceMeshPeer{},
			resyncPeriod,
			indexers,
		)
	}

	return informers.NewMultiNamespaceInformer(namespaces, resyncPeriod, newInformer)
}

func (f *serviceMeshPeerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServiceMeshPeerInformer(client, f.namespaces, resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serviceMeshPeerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&federationv1.ServiceMeshPeer{}, f.defaultInformer)
}

func (f *serviceMeshPeerInformer) Lister() v1.ServiceMeshPeerLister {
	return v1.NewServiceMeshPeerLister(f.Informer().GetIndexer())
}
