/*
Copyright 2023.

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
	"context"

	v1 "github.com/apache/doris-operator/api/disaggregated/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDorisDisaggregatedClusters implements DorisDisaggregatedClusterInterface
type FakeDorisDisaggregatedClusters struct {
	Fake *FakeDisaggregatedV1
	ns   string
}

var dorisdisaggregatedclustersResource = v1.SchemeGroupVersion.WithResource("dorisdisaggregatedclusters")

var dorisdisaggregatedclustersKind = v1.SchemeGroupVersion.WithKind("DorisDisaggregatedCluster")

// Get takes name of the dorisDisaggregatedCluster, and returns the corresponding dorisDisaggregatedCluster object, and an error if there is any.
func (c *FakeDorisDisaggregatedClusters) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.DorisDisaggregatedCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dorisdisaggregatedclustersResource, c.ns, name), &v1.DorisDisaggregatedCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DorisDisaggregatedCluster), err
}

// List takes label and field selectors, and returns the list of DorisDisaggregatedClusters that match those selectors.
func (c *FakeDorisDisaggregatedClusters) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DorisDisaggregatedClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dorisdisaggregatedclustersResource, dorisdisaggregatedclustersKind, c.ns, opts), &v1.DorisDisaggregatedClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.DorisDisaggregatedClusterList{ListMeta: obj.(*v1.DorisDisaggregatedClusterList).ListMeta}
	for _, item := range obj.(*v1.DorisDisaggregatedClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dorisDisaggregatedClusters.
func (c *FakeDorisDisaggregatedClusters) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dorisdisaggregatedclustersResource, c.ns, opts))

}

// Create takes the representation of a dorisDisaggregatedCluster and creates it.  Returns the server's representation of the dorisDisaggregatedCluster, and an error, if there is any.
func (c *FakeDorisDisaggregatedClusters) Create(ctx context.Context, dorisDisaggregatedCluster *v1.DorisDisaggregatedCluster, opts metav1.CreateOptions) (result *v1.DorisDisaggregatedCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dorisdisaggregatedclustersResource, c.ns, dorisDisaggregatedCluster), &v1.DorisDisaggregatedCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DorisDisaggregatedCluster), err
}

// Update takes the representation of a dorisDisaggregatedCluster and updates it. Returns the server's representation of the dorisDisaggregatedCluster, and an error, if there is any.
func (c *FakeDorisDisaggregatedClusters) Update(ctx context.Context, dorisDisaggregatedCluster *v1.DorisDisaggregatedCluster, opts metav1.UpdateOptions) (result *v1.DorisDisaggregatedCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dorisdisaggregatedclustersResource, c.ns, dorisDisaggregatedCluster), &v1.DorisDisaggregatedCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DorisDisaggregatedCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDorisDisaggregatedClusters) UpdateStatus(ctx context.Context, dorisDisaggregatedCluster *v1.DorisDisaggregatedCluster, opts metav1.UpdateOptions) (*v1.DorisDisaggregatedCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dorisdisaggregatedclustersResource, "status", c.ns, dorisDisaggregatedCluster), &v1.DorisDisaggregatedCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DorisDisaggregatedCluster), err
}

// Delete takes name of the dorisDisaggregatedCluster and deletes it. Returns an error if one occurs.
func (c *FakeDorisDisaggregatedClusters) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(dorisdisaggregatedclustersResource, c.ns, name, opts), &v1.DorisDisaggregatedCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDorisDisaggregatedClusters) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dorisdisaggregatedclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.DorisDisaggregatedClusterList{})
	return err
}

// Patch applies the patch and returns the patched dorisDisaggregatedCluster.
func (c *FakeDorisDisaggregatedClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DorisDisaggregatedCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dorisdisaggregatedclustersResource, c.ns, name, pt, data, subresources...), &v1.DorisDisaggregatedCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DorisDisaggregatedCluster), err
}
