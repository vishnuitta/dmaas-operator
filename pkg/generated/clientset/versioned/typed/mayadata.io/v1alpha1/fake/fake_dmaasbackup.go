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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/mayadata-io/dmaas-operator/pkg/apis/mayadata.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDMaaSBackups implements DMaaSBackupInterface
type FakeDMaaSBackups struct {
	Fake *FakeMayadataV1alpha1
	ns   string
}

var dmaasbackupsResource = schema.GroupVersionResource{Group: "mayadata.io", Version: "v1alpha1", Resource: "dmaasbackups"}

var dmaasbackupsKind = schema.GroupVersionKind{Group: "mayadata.io", Version: "v1alpha1", Kind: "DMaaSBackup"}

// Get takes name of the dMaaSBackup, and returns the corresponding dMaaSBackup object, and an error if there is any.
func (c *FakeDMaaSBackups) Get(name string, options v1.GetOptions) (result *v1alpha1.DMaaSBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dmaasbackupsResource, c.ns, name), &v1alpha1.DMaaSBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DMaaSBackup), err
}

// List takes label and field selectors, and returns the list of DMaaSBackups that match those selectors.
func (c *FakeDMaaSBackups) List(opts v1.ListOptions) (result *v1alpha1.DMaaSBackupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dmaasbackupsResource, dmaasbackupsKind, c.ns, opts), &v1alpha1.DMaaSBackupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DMaaSBackupList{ListMeta: obj.(*v1alpha1.DMaaSBackupList).ListMeta}
	for _, item := range obj.(*v1alpha1.DMaaSBackupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dMaaSBackups.
func (c *FakeDMaaSBackups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dmaasbackupsResource, c.ns, opts))

}

// Create takes the representation of a dMaaSBackup and creates it.  Returns the server's representation of the dMaaSBackup, and an error, if there is any.
func (c *FakeDMaaSBackups) Create(dMaaSBackup *v1alpha1.DMaaSBackup) (result *v1alpha1.DMaaSBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dmaasbackupsResource, c.ns, dMaaSBackup), &v1alpha1.DMaaSBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DMaaSBackup), err
}

// Update takes the representation of a dMaaSBackup and updates it. Returns the server's representation of the dMaaSBackup, and an error, if there is any.
func (c *FakeDMaaSBackups) Update(dMaaSBackup *v1alpha1.DMaaSBackup) (result *v1alpha1.DMaaSBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dmaasbackupsResource, c.ns, dMaaSBackup), &v1alpha1.DMaaSBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DMaaSBackup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDMaaSBackups) UpdateStatus(dMaaSBackup *v1alpha1.DMaaSBackup) (*v1alpha1.DMaaSBackup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dmaasbackupsResource, "status", c.ns, dMaaSBackup), &v1alpha1.DMaaSBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DMaaSBackup), err
}

// Delete takes name of the dMaaSBackup and deletes it. Returns an error if one occurs.
func (c *FakeDMaaSBackups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dmaasbackupsResource, c.ns, name), &v1alpha1.DMaaSBackup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDMaaSBackups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dmaasbackupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DMaaSBackupList{})
	return err
}

// Patch applies the patch and returns the patched dMaaSBackup.
func (c *FakeDMaaSBackups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DMaaSBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dmaasbackupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DMaaSBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DMaaSBackup), err
}
