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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	fuxiv1 "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeWorkloadses implements WorkloadsInterface
type FakeWorkloadses struct {
	Fake *FakeFuxiV1
	ns   string
}

var workloadsesResource = schema.GroupVersionResource{Group: "fuxi.nio.io", Version: "v1", Resource: "workloadses"}

var workloadsesKind = schema.GroupVersionKind{Group: "fuxi.nio.io", Version: "v1", Kind: "Workloads"}

// Get takes name of the workloads, and returns the corresponding workloads object, and an error if there is any.
func (c *FakeWorkloadses) Get(name string, options v1.GetOptions) (result *fuxiv1.Workloads, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(workloadsesResource, c.ns, name), &fuxiv1.Workloads{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fuxiv1.Workloads), err
}

// List takes label and field selectors, and returns the list of Workloadses that match those selectors.
func (c *FakeWorkloadses) List(opts v1.ListOptions) (result *fuxiv1.WorkloadsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(workloadsesResource, workloadsesKind, c.ns, opts), &fuxiv1.WorkloadsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &fuxiv1.WorkloadsList{ListMeta: obj.(*fuxiv1.WorkloadsList).ListMeta}
	for _, item := range obj.(*fuxiv1.WorkloadsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested workloadses.
func (c *FakeWorkloadses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(workloadsesResource, c.ns, opts))

}

// Create takes the representation of a workloads and creates it.  Returns the server's representation of the workloads, and an error, if there is any.
func (c *FakeWorkloadses) Create(workloads *fuxiv1.Workloads) (result *fuxiv1.Workloads, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(workloadsesResource, c.ns, workloads), &fuxiv1.Workloads{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fuxiv1.Workloads), err
}

// Update takes the representation of a workloads and updates it. Returns the server's representation of the workloads, and an error, if there is any.
func (c *FakeWorkloadses) Update(workloads *fuxiv1.Workloads) (result *fuxiv1.Workloads, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(workloadsesResource, c.ns, workloads), &fuxiv1.Workloads{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fuxiv1.Workloads), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWorkloadses) UpdateStatus(workloads *fuxiv1.Workloads) (*fuxiv1.Workloads, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(workloadsesResource, "status", c.ns, workloads), &fuxiv1.Workloads{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fuxiv1.Workloads), err
}

// Delete takes name of the workloads and deletes it. Returns an error if one occurs.
func (c *FakeWorkloadses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(workloadsesResource, c.ns, name), &fuxiv1.Workloads{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWorkloadses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(workloadsesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &fuxiv1.WorkloadsList{})
	return err
}

// Patch applies the patch and returns the patched workloads.
func (c *FakeWorkloadses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *fuxiv1.Workloads, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(workloadsesResource, c.ns, name, pt, data, subresources...), &fuxiv1.Workloads{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fuxiv1.Workloads), err
}
