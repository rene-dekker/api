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

// RemoteClusterConfigurationsGetter has a method to return a RemoteClusterConfigurationInterface.
// A group's client should implement this interface.
type RemoteClusterConfigurationsGetter interface {
	RemoteClusterConfigurations() RemoteClusterConfigurationInterface
}

// RemoteClusterConfigurationInterface has methods to work with RemoteClusterConfiguration resources.
type RemoteClusterConfigurationInterface interface {
	Create(ctx context.Context, remoteClusterConfiguration *v3.RemoteClusterConfiguration, opts v1.CreateOptions) (*v3.RemoteClusterConfiguration, error)
	Update(ctx context.Context, remoteClusterConfiguration *v3.RemoteClusterConfiguration, opts v1.UpdateOptions) (*v3.RemoteClusterConfiguration, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.RemoteClusterConfiguration, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.RemoteClusterConfigurationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.RemoteClusterConfiguration, err error)
	RemoteClusterConfigurationExpansion
}

// remoteClusterConfigurations implements RemoteClusterConfigurationInterface
type remoteClusterConfigurations struct {
	client rest.Interface
}

// newRemoteClusterConfigurations returns a RemoteClusterConfigurations
func newRemoteClusterConfigurations(c *ProjectcalicoV3Client) *remoteClusterConfigurations {
	return &remoteClusterConfigurations{
		client: c.RESTClient(),
	}
}

// Get takes name of the remoteClusterConfiguration, and returns the corresponding remoteClusterConfiguration object, and an error if there is any.
func (c *remoteClusterConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.RemoteClusterConfiguration, err error) {
	result = &v3.RemoteClusterConfiguration{}
	err = c.client.Get().
		Resource("remoteclusterconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RemoteClusterConfigurations that match those selectors.
func (c *remoteClusterConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v3.RemoteClusterConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.RemoteClusterConfigurationList{}
	err = c.client.Get().
		Resource("remoteclusterconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested remoteClusterConfigurations.
func (c *remoteClusterConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("remoteclusterconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a remoteClusterConfiguration and creates it.  Returns the server's representation of the remoteClusterConfiguration, and an error, if there is any.
func (c *remoteClusterConfigurations) Create(ctx context.Context, remoteClusterConfiguration *v3.RemoteClusterConfiguration, opts v1.CreateOptions) (result *v3.RemoteClusterConfiguration, err error) {
	result = &v3.RemoteClusterConfiguration{}
	err = c.client.Post().
		Resource("remoteclusterconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(remoteClusterConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a remoteClusterConfiguration and updates it. Returns the server's representation of the remoteClusterConfiguration, and an error, if there is any.
func (c *remoteClusterConfigurations) Update(ctx context.Context, remoteClusterConfiguration *v3.RemoteClusterConfiguration, opts v1.UpdateOptions) (result *v3.RemoteClusterConfiguration, err error) {
	result = &v3.RemoteClusterConfiguration{}
	err = c.client.Put().
		Resource("remoteclusterconfigurations").
		Name(remoteClusterConfiguration.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(remoteClusterConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the remoteClusterConfiguration and deletes it. Returns an error if one occurs.
func (c *remoteClusterConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("remoteclusterconfigurations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *remoteClusterConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("remoteclusterconfigurations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched remoteClusterConfiguration.
func (c *remoteClusterConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.RemoteClusterConfiguration, err error) {
	result = &v3.RemoteClusterConfiguration{}
	err = c.client.Patch(pt).
		Resource("remoteclusterconfigurations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
