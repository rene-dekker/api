// Copyright (c) 2022 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package v3

import (
	"context"
	"time"

	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	scheme "github.com/tigera/api/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// StagedGlobalNetworkPoliciesGetter has a method to return a StagedGlobalNetworkPolicyInterface.
// A group's client should implement this interface.
type StagedGlobalNetworkPoliciesGetter interface {
	StagedGlobalNetworkPolicies() StagedGlobalNetworkPolicyInterface
}

// StagedGlobalNetworkPolicyInterface has methods to work with StagedGlobalNetworkPolicy resources.
type StagedGlobalNetworkPolicyInterface interface {
	Create(ctx context.Context, stagedGlobalNetworkPolicy *v3.StagedGlobalNetworkPolicy, opts v1.CreateOptions) (*v3.StagedGlobalNetworkPolicy, error)
	Update(ctx context.Context, stagedGlobalNetworkPolicy *v3.StagedGlobalNetworkPolicy, opts v1.UpdateOptions) (*v3.StagedGlobalNetworkPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.StagedGlobalNetworkPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.StagedGlobalNetworkPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.StagedGlobalNetworkPolicy, err error)
	StagedGlobalNetworkPolicyExpansion
}

// stagedGlobalNetworkPolicies implements StagedGlobalNetworkPolicyInterface
type stagedGlobalNetworkPolicies struct {
	client rest.Interface
}

// newStagedGlobalNetworkPolicies returns a StagedGlobalNetworkPolicies
func newStagedGlobalNetworkPolicies(c *ProjectcalicoV3Client) *stagedGlobalNetworkPolicies {
	return &stagedGlobalNetworkPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the stagedGlobalNetworkPolicy, and returns the corresponding stagedGlobalNetworkPolicy object, and an error if there is any.
func (c *stagedGlobalNetworkPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.StagedGlobalNetworkPolicy, err error) {
	result = &v3.StagedGlobalNetworkPolicy{}
	err = c.client.Get().
		Resource("stagedglobalnetworkpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StagedGlobalNetworkPolicies that match those selectors.
func (c *stagedGlobalNetworkPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v3.StagedGlobalNetworkPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.StagedGlobalNetworkPolicyList{}
	err = c.client.Get().
		Resource("stagedglobalnetworkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested stagedGlobalNetworkPolicies.
func (c *stagedGlobalNetworkPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("stagedglobalnetworkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a stagedGlobalNetworkPolicy and creates it.  Returns the server's representation of the stagedGlobalNetworkPolicy, and an error, if there is any.
func (c *stagedGlobalNetworkPolicies) Create(ctx context.Context, stagedGlobalNetworkPolicy *v3.StagedGlobalNetworkPolicy, opts v1.CreateOptions) (result *v3.StagedGlobalNetworkPolicy, err error) {
	result = &v3.StagedGlobalNetworkPolicy{}
	err = c.client.Post().
		Resource("stagedglobalnetworkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(stagedGlobalNetworkPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a stagedGlobalNetworkPolicy and updates it. Returns the server's representation of the stagedGlobalNetworkPolicy, and an error, if there is any.
func (c *stagedGlobalNetworkPolicies) Update(ctx context.Context, stagedGlobalNetworkPolicy *v3.StagedGlobalNetworkPolicy, opts v1.UpdateOptions) (result *v3.StagedGlobalNetworkPolicy, err error) {
	result = &v3.StagedGlobalNetworkPolicy{}
	err = c.client.Put().
		Resource("stagedglobalnetworkpolicies").
		Name(stagedGlobalNetworkPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(stagedGlobalNetworkPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the stagedGlobalNetworkPolicy and deletes it. Returns an error if one occurs.
func (c *stagedGlobalNetworkPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("stagedglobalnetworkpolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *stagedGlobalNetworkPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("stagedglobalnetworkpolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched stagedGlobalNetworkPolicy.
func (c *stagedGlobalNetworkPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.StagedGlobalNetworkPolicy, err error) {
	result = &v3.StagedGlobalNetworkPolicy{}
	err = c.client.Patch(pt).
		Resource("stagedglobalnetworkpolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
