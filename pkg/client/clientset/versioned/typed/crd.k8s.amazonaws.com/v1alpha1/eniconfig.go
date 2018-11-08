// Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/aws/amazon-vpc-cni-k8s/pkg/apis/crd.k8s.amazonaws.com/v1alpha1"
	scheme "github.com/aws/amazon-vpc-cni-k8s/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ENIConfigsGetter has a method to return a ENIConfigInterface.
// A group's client should implement this interface.
type ENIConfigsGetter interface {
	ENIConfigs(namespace string) ENIConfigInterface
}

// ENIConfigInterface has methods to work with ENIConfig resources.
type ENIConfigInterface interface {
	Create(*v1alpha1.ENIConfig) (*v1alpha1.ENIConfig, error)
	Update(*v1alpha1.ENIConfig) (*v1alpha1.ENIConfig, error)
	UpdateStatus(*v1alpha1.ENIConfig) (*v1alpha1.ENIConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ENIConfig, error)
	List(opts v1.ListOptions) (*v1alpha1.ENIConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ENIConfig, err error)
	ENIConfigExpansion
}

// eNIConfigs implements ENIConfigInterface
type eNIConfigs struct {
	client rest.Interface
	ns     string
}

// newENIConfigs returns a ENIConfigs
func newENIConfigs(c *CrdV1alpha1Client, namespace string) *eNIConfigs {
	return &eNIConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the eNIConfig, and returns the corresponding eNIConfig object, and an error if there is any.
func (c *eNIConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.ENIConfig, err error) {
	result = &v1alpha1.ENIConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eniconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ENIConfigs that match those selectors.
func (c *eNIConfigs) List(opts v1.ListOptions) (result *v1alpha1.ENIConfigList, err error) {
	result = &v1alpha1.ENIConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eniconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested eNIConfigs.
func (c *eNIConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("eniconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a eNIConfig and creates it.  Returns the server's representation of the eNIConfig, and an error, if there is any.
func (c *eNIConfigs) Create(eNIConfig *v1alpha1.ENIConfig) (result *v1alpha1.ENIConfig, err error) {
	result = &v1alpha1.ENIConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("eniconfigs").
		Body(eNIConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a eNIConfig and updates it. Returns the server's representation of the eNIConfig, and an error, if there is any.
func (c *eNIConfigs) Update(eNIConfig *v1alpha1.ENIConfig) (result *v1alpha1.ENIConfig, err error) {
	result = &v1alpha1.ENIConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("eniconfigs").
		Name(eNIConfig.Name).
		Body(eNIConfig).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *eNIConfigs) UpdateStatus(eNIConfig *v1alpha1.ENIConfig) (result *v1alpha1.ENIConfig, err error) {
	result = &v1alpha1.ENIConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("eniconfigs").
		Name(eNIConfig.Name).
		SubResource("status").
		Body(eNIConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the eNIConfig and deletes it. Returns an error if one occurs.
func (c *eNIConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eniconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *eNIConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eniconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched eNIConfig.
func (c *eNIConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ENIConfig, err error) {
	result = &v1alpha1.ENIConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("eniconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
