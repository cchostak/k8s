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

package v1alpha1

import (
	v1alpha1 "github.com/rook/rook/pkg/apis/edgefs.rook.io/v1alpha1"
	scheme "github.com/rook/rook/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// S3XsGetter has a method to return a S3XInterface.
// A group's client should implement this interface.
type S3XsGetter interface {
	S3Xs(namespace string) S3XInterface
}

// S3XInterface has methods to work with S3X resources.
type S3XInterface interface {
	Create(*v1alpha1.S3X) (*v1alpha1.S3X, error)
	Update(*v1alpha1.S3X) (*v1alpha1.S3X, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.S3X, error)
	List(opts v1.ListOptions) (*v1alpha1.S3XList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.S3X, err error)
	S3XExpansion
}

// s3Xs implements S3XInterface
type s3Xs struct {
	client rest.Interface
	ns     string
}

// newS3Xs returns a S3Xs
func newS3Xs(c *EdgefsV1alpha1Client, namespace string) *s3Xs {
	return &s3Xs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the s3X, and returns the corresponding s3X object, and an error if there is any.
func (c *s3Xs) Get(name string, options v1.GetOptions) (result *v1alpha1.S3X, err error) {
	result = &v1alpha1.S3X{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s3xs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of S3Xs that match those selectors.
func (c *s3Xs) List(opts v1.ListOptions) (result *v1alpha1.S3XList, err error) {
	result = &v1alpha1.S3XList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s3xs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested s3Xs.
func (c *s3Xs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("s3xs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a s3X and creates it.  Returns the server's representation of the s3X, and an error, if there is any.
func (c *s3Xs) Create(s3X *v1alpha1.S3X) (result *v1alpha1.S3X, err error) {
	result = &v1alpha1.S3X{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("s3xs").
		Body(s3X).
		Do().
		Into(result)
	return
}

// Update takes the representation of a s3X and updates it. Returns the server's representation of the s3X, and an error, if there is any.
func (c *s3Xs) Update(s3X *v1alpha1.S3X) (result *v1alpha1.S3X, err error) {
	result = &v1alpha1.S3X{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("s3xs").
		Name(s3X.Name).
		Body(s3X).
		Do().
		Into(result)
	return
}

// Delete takes name of the s3X and deletes it. Returns an error if one occurs.
func (c *s3Xs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s3xs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *s3Xs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s3xs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched s3X.
func (c *s3Xs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.S3X, err error) {
	result = &v1alpha1.S3X{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("s3xs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
