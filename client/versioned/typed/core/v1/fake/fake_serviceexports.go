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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	corev1 "maistra.io/api/core/v1"
)

// FakeServiceExports implements ServiceExportsInterface
type FakeServiceExports struct {
	Fake *FakeCoreV1
	ns   string
}

var serviceexportsResource = schema.GroupVersionResource{Group: "maistra.io", Version: "v1", Resource: "serviceexports"}

var serviceexportsKind = schema.GroupVersionKind{Group: "maistra.io", Version: "v1", Kind: "ServiceExports"}

// Get takes name of the serviceExports, and returns the corresponding serviceExports object, and an error if there is any.
func (c *FakeServiceExports) Get(ctx context.Context, name string, options v1.GetOptions) (result *corev1.ServiceExports, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(serviceexportsResource, c.ns, name), &corev1.ServiceExports{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceExports), err
}

// List takes label and field selectors, and returns the list of ServiceExports that match those selectors.
func (c *FakeServiceExports) List(ctx context.Context, opts v1.ListOptions) (result *corev1.ServiceExportsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(serviceexportsResource, serviceexportsKind, c.ns, opts), &corev1.ServiceExportsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.ServiceExportsList{ListMeta: obj.(*corev1.ServiceExportsList).ListMeta}
	for _, item := range obj.(*corev1.ServiceExportsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceExports.
func (c *FakeServiceExports) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(serviceexportsResource, c.ns, opts))

}

// Create takes the representation of a serviceExports and creates it.  Returns the server's representation of the serviceExports, and an error, if there is any.
func (c *FakeServiceExports) Create(ctx context.Context, serviceExports *corev1.ServiceExports, opts v1.CreateOptions) (result *corev1.ServiceExports, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(serviceexportsResource, c.ns, serviceExports), &corev1.ServiceExports{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceExports), err
}

// Update takes the representation of a serviceExports and updates it. Returns the server's representation of the serviceExports, and an error, if there is any.
func (c *FakeServiceExports) Update(ctx context.Context, serviceExports *corev1.ServiceExports, opts v1.UpdateOptions) (result *corev1.ServiceExports, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(serviceexportsResource, c.ns, serviceExports), &corev1.ServiceExports{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceExports), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceExports) UpdateStatus(ctx context.Context, serviceExports *corev1.ServiceExports, opts v1.UpdateOptions) (*corev1.ServiceExports, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(serviceexportsResource, "status", c.ns, serviceExports), &corev1.ServiceExports{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceExports), err
}

// Delete takes name of the serviceExports and deletes it. Returns an error if one occurs.
func (c *FakeServiceExports) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(serviceexportsResource, c.ns, name), &corev1.ServiceExports{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceExports) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(serviceexportsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &corev1.ServiceExportsList{})
	return err
}

// Patch applies the patch and returns the patched serviceExports.
func (c *FakeServiceExports) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *corev1.ServiceExports, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(serviceexportsResource, c.ns, name, pt, data, subresources...), &corev1.ServiceExports{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceExports), err
}
