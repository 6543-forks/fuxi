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

// FakeNodeLivenesses implements NodeLivenessInterface
type FakeNodeLivenesses struct {
	Fake *FakeFuxiV1
	ns   string
}

var nodelivenessesResource = schema.GroupVersionResource{Group: "fuxi.nio.io", Version: "v1", Resource: "nodelivenesses"}

var nodelivenessesKind = schema.GroupVersionKind{Group: "fuxi.nio.io", Version: "v1", Kind: "NodeLiveness"}

// Get takes name of the nodeLiveness, and returns the corresponding nodeLiveness object, and an error if there is any.
func (c *FakeNodeLivenesses) Get(name string, options v1.GetOptions) (result *fuxiv1.NodeLiveness, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nodelivenessesResource, c.ns, name), &fuxiv1.NodeLiveness{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fuxiv1.NodeLiveness), err
}

// List takes label and field selectors, and returns the list of NodeLivenesses that match those selectors.
func (c *FakeNodeLivenesses) List(opts v1.ListOptions) (result *fuxiv1.NodeLivenessList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nodelivenessesResource, nodelivenessesKind, c.ns, opts), &fuxiv1.NodeLivenessList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &fuxiv1.NodeLivenessList{ListMeta: obj.(*fuxiv1.NodeLivenessList).ListMeta}
	for _, item := range obj.(*fuxiv1.NodeLivenessList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodeLivenesses.
func (c *FakeNodeLivenesses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nodelivenessesResource, c.ns, opts))

}

// Create takes the representation of a nodeLiveness and creates it.  Returns the server's representation of the nodeLiveness, and an error, if there is any.
func (c *FakeNodeLivenesses) Create(nodeLiveness *fuxiv1.NodeLiveness) (result *fuxiv1.NodeLiveness, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nodelivenessesResource, c.ns, nodeLiveness), &fuxiv1.NodeLiveness{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fuxiv1.NodeLiveness), err
}

// Update takes the representation of a nodeLiveness and updates it. Returns the server's representation of the nodeLiveness, and an error, if there is any.
func (c *FakeNodeLivenesses) Update(nodeLiveness *fuxiv1.NodeLiveness) (result *fuxiv1.NodeLiveness, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nodelivenessesResource, c.ns, nodeLiveness), &fuxiv1.NodeLiveness{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fuxiv1.NodeLiveness), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNodeLivenesses) UpdateStatus(nodeLiveness *fuxiv1.NodeLiveness) (*fuxiv1.NodeLiveness, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(nodelivenessesResource, "status", c.ns, nodeLiveness), &fuxiv1.NodeLiveness{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fuxiv1.NodeLiveness), err
}

// Delete takes name of the nodeLiveness and deletes it. Returns an error if one occurs.
func (c *FakeNodeLivenesses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(nodelivenessesResource, c.ns, name), &fuxiv1.NodeLiveness{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodeLivenesses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nodelivenessesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &fuxiv1.NodeLivenessList{})
	return err
}

// Patch applies the patch and returns the patched nodeLiveness.
func (c *FakeNodeLivenesses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *fuxiv1.NodeLiveness, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nodelivenessesResource, c.ns, name, pt, data, subresources...), &fuxiv1.NodeLiveness{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fuxiv1.NodeLiveness), err
}
