// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeProjects implements ProjectInterface
type FakeProjects struct {
	Fake *FakeCoreV1beta1
}

var projectsResource = schema.GroupVersionResource{Group: "core.gardener.cloud", Version: "v1beta1", Resource: "projects"}

var projectsKind = schema.GroupVersionKind{Group: "core.gardener.cloud", Version: "v1beta1", Kind: "Project"}

// Get takes name of the project, and returns the corresponding project object, and an error if there is any.
func (c *FakeProjects) Get(name string, options v1.GetOptions) (result *v1beta1.Project, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(projectsResource, name), &v1beta1.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Project), err
}

// List takes label and field selectors, and returns the list of Projects that match those selectors.
func (c *FakeProjects) List(opts v1.ListOptions) (result *v1beta1.ProjectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(projectsResource, projectsKind, opts), &v1beta1.ProjectList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ProjectList{ListMeta: obj.(*v1beta1.ProjectList).ListMeta}
	for _, item := range obj.(*v1beta1.ProjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested projects.
func (c *FakeProjects) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(projectsResource, opts))
}

// Create takes the representation of a project and creates it.  Returns the server's representation of the project, and an error, if there is any.
func (c *FakeProjects) Create(project *v1beta1.Project) (result *v1beta1.Project, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(projectsResource, project), &v1beta1.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Project), err
}

// Update takes the representation of a project and updates it. Returns the server's representation of the project, and an error, if there is any.
func (c *FakeProjects) Update(project *v1beta1.Project) (result *v1beta1.Project, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(projectsResource, project), &v1beta1.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Project), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProjects) UpdateStatus(project *v1beta1.Project) (*v1beta1.Project, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(projectsResource, "status", project), &v1beta1.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Project), err
}

// Delete takes name of the project and deletes it. Returns an error if one occurs.
func (c *FakeProjects) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(projectsResource, name), &v1beta1.Project{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProjects) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(projectsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.ProjectList{})
	return err
}

// Patch applies the patch and returns the patched project.
func (c *FakeProjects) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Project, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(projectsResource, name, pt, data, subresources...), &v1beta1.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Project), err
}
