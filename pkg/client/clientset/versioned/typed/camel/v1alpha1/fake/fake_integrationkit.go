/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

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
	v1alpha1 "github.com/apache/camel-k/pkg/apis/camel/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIntegrationKits implements IntegrationKitInterface
type FakeIntegrationKits struct {
	Fake *FakeCamelV1alpha1
	ns   string
}

var integrationkitsResource = schema.GroupVersionResource{Group: "camel.apache.org", Version: "v1alpha1", Resource: "integrationkits"}

var integrationkitsKind = schema.GroupVersionKind{Group: "camel.apache.org", Version: "v1alpha1", Kind: "IntegrationKit"}

// Get takes name of the integrationKit, and returns the corresponding integrationKit object, and an error if there is any.
func (c *FakeIntegrationKits) Get(name string, options v1.GetOptions) (result *v1alpha1.IntegrationKit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(integrationkitsResource, c.ns, name), &v1alpha1.IntegrationKit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IntegrationKit), err
}

// List takes label and field selectors, and returns the list of IntegrationKits that match those selectors.
func (c *FakeIntegrationKits) List(opts v1.ListOptions) (result *v1alpha1.IntegrationKitList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(integrationkitsResource, integrationkitsKind, c.ns, opts), &v1alpha1.IntegrationKitList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IntegrationKitList{ListMeta: obj.(*v1alpha1.IntegrationKitList).ListMeta}
	for _, item := range obj.(*v1alpha1.IntegrationKitList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested integrationKits.
func (c *FakeIntegrationKits) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(integrationkitsResource, c.ns, opts))

}

// Create takes the representation of a integrationKit and creates it.  Returns the server's representation of the integrationKit, and an error, if there is any.
func (c *FakeIntegrationKits) Create(integrationKit *v1alpha1.IntegrationKit) (result *v1alpha1.IntegrationKit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(integrationkitsResource, c.ns, integrationKit), &v1alpha1.IntegrationKit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IntegrationKit), err
}

// Update takes the representation of a integrationKit and updates it. Returns the server's representation of the integrationKit, and an error, if there is any.
func (c *FakeIntegrationKits) Update(integrationKit *v1alpha1.IntegrationKit) (result *v1alpha1.IntegrationKit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(integrationkitsResource, c.ns, integrationKit), &v1alpha1.IntegrationKit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IntegrationKit), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIntegrationKits) UpdateStatus(integrationKit *v1alpha1.IntegrationKit) (*v1alpha1.IntegrationKit, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(integrationkitsResource, "status", c.ns, integrationKit), &v1alpha1.IntegrationKit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IntegrationKit), err
}

// Delete takes name of the integrationKit and deletes it. Returns an error if one occurs.
func (c *FakeIntegrationKits) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(integrationkitsResource, c.ns, name), &v1alpha1.IntegrationKit{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIntegrationKits) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(integrationkitsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IntegrationKitList{})
	return err
}

// Patch applies the patch and returns the patched integrationKit.
func (c *FakeIntegrationKits) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IntegrationKit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(integrationkitsResource, c.ns, name, pt, data, subresources...), &v1alpha1.IntegrationKit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IntegrationKit), err
}