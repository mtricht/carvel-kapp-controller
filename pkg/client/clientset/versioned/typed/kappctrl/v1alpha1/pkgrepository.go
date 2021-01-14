// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/kappctrl/v1alpha1"
	scheme "github.com/vmware-tanzu/carvel-kapp-controller/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PkgRepositoriesGetter has a method to return a PkgRepositoryInterface.
// A group's client should implement this interface.
type PkgRepositoriesGetter interface {
	PkgRepositories(namespace string) PkgRepositoryInterface
}

// PkgRepositoryInterface has methods to work with PkgRepository resources.
type PkgRepositoryInterface interface {
	Create(*v1alpha1.PkgRepository) (*v1alpha1.PkgRepository, error)
	Update(*v1alpha1.PkgRepository) (*v1alpha1.PkgRepository, error)
	UpdateStatus(*v1alpha1.PkgRepository) (*v1alpha1.PkgRepository, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PkgRepository, error)
	List(opts v1.ListOptions) (*v1alpha1.PkgRepositoryList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PkgRepository, err error)
	PkgRepositoryExpansion
}

// pkgRepositories implements PkgRepositoryInterface
type pkgRepositories struct {
	client rest.Interface
	ns     string
}

// newPkgRepositories returns a PkgRepositories
func newPkgRepositories(c *KappctrlV1alpha1Client, namespace string) *pkgRepositories {
	return &pkgRepositories{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the pkgRepository, and returns the corresponding pkgRepository object, and an error if there is any.
func (c *pkgRepositories) Get(name string, options v1.GetOptions) (result *v1alpha1.PkgRepository, err error) {
	result = &v1alpha1.PkgRepository{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pkgrepositories").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PkgRepositories that match those selectors.
func (c *pkgRepositories) List(opts v1.ListOptions) (result *v1alpha1.PkgRepositoryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PkgRepositoryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pkgrepositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pkgRepositories.
func (c *pkgRepositories) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("pkgrepositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a pkgRepository and creates it.  Returns the server's representation of the pkgRepository, and an error, if there is any.
func (c *pkgRepositories) Create(pkgRepository *v1alpha1.PkgRepository) (result *v1alpha1.PkgRepository, err error) {
	result = &v1alpha1.PkgRepository{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("pkgrepositories").
		Body(pkgRepository).
		Do().
		Into(result)
	return
}

// Update takes the representation of a pkgRepository and updates it. Returns the server's representation of the pkgRepository, and an error, if there is any.
func (c *pkgRepositories) Update(pkgRepository *v1alpha1.PkgRepository) (result *v1alpha1.PkgRepository, err error) {
	result = &v1alpha1.PkgRepository{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pkgrepositories").
		Name(pkgRepository.Name).
		Body(pkgRepository).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *pkgRepositories) UpdateStatus(pkgRepository *v1alpha1.PkgRepository) (result *v1alpha1.PkgRepository, err error) {
	result = &v1alpha1.PkgRepository{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pkgrepositories").
		Name(pkgRepository.Name).
		SubResource("status").
		Body(pkgRepository).
		Do().
		Into(result)
	return
}

// Delete takes name of the pkgRepository and deletes it. Returns an error if one occurs.
func (c *pkgRepositories) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pkgrepositories").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pkgRepositories) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pkgrepositories").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched pkgRepository.
func (c *pkgRepositories) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PkgRepository, err error) {
	result = &v1alpha1.PkgRepository{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("pkgrepositories").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
