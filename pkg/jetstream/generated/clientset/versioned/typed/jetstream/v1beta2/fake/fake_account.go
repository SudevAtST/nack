// Copyright 2020 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
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

	v1beta2 "github.com/nats-io/nack/pkg/jetstream/apis/jetstream/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAccounts implements AccountInterface
type FakeAccounts struct {
	Fake *FakeJetstreamV1beta2
	ns   string
}

var accountsResource = schema.GroupVersionResource{Group: "jetstream.nats.io", Version: "v1beta2", Resource: "accounts"}

var accountsKind = schema.GroupVersionKind{Group: "jetstream.nats.io", Version: "v1beta2", Kind: "Account"}

// Get takes name of the account, and returns the corresponding account object, and an error if there is any.
func (c *FakeAccounts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.Account, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(accountsResource, c.ns, name), &v1beta2.Account{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.Account), err
}

// List takes label and field selectors, and returns the list of Accounts that match those selectors.
func (c *FakeAccounts) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.AccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(accountsResource, accountsKind, c.ns, opts), &v1beta2.AccountList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.AccountList{ListMeta: obj.(*v1beta2.AccountList).ListMeta}
	for _, item := range obj.(*v1beta2.AccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested accounts.
func (c *FakeAccounts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(accountsResource, c.ns, opts))

}

// Create takes the representation of a account and creates it.  Returns the server's representation of the account, and an error, if there is any.
func (c *FakeAccounts) Create(ctx context.Context, account *v1beta2.Account, opts v1.CreateOptions) (result *v1beta2.Account, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(accountsResource, c.ns, account), &v1beta2.Account{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.Account), err
}

// Update takes the representation of a account and updates it. Returns the server's representation of the account, and an error, if there is any.
func (c *FakeAccounts) Update(ctx context.Context, account *v1beta2.Account, opts v1.UpdateOptions) (result *v1beta2.Account, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(accountsResource, c.ns, account), &v1beta2.Account{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.Account), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAccounts) UpdateStatus(ctx context.Context, account *v1beta2.Account, opts v1.UpdateOptions) (*v1beta2.Account, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(accountsResource, "status", c.ns, account), &v1beta2.Account{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.Account), err
}

// Delete takes name of the account and deletes it. Returns an error if one occurs.
func (c *FakeAccounts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(accountsResource, c.ns, name, opts), &v1beta2.Account{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAccounts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(accountsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta2.AccountList{})
	return err
}

// Patch applies the patch and returns the patched account.
func (c *FakeAccounts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.Account, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(accountsResource, c.ns, name, pt, data, subresources...), &v1beta2.Account{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.Account), err
}
