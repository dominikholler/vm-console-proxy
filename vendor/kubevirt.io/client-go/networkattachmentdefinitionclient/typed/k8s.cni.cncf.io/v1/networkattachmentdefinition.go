/*
Copyright The KubeVirt Authors.

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

package v1

import (
	"context"

	v1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	scheme "kubevirt.io/client-go/networkattachmentdefinitionclient/scheme"
)

// NetworkAttachmentDefinitionsGetter has a method to return a NetworkAttachmentDefinitionInterface.
// A group's client should implement this interface.
type NetworkAttachmentDefinitionsGetter interface {
	NetworkAttachmentDefinitions(namespace string) NetworkAttachmentDefinitionInterface
}

// NetworkAttachmentDefinitionInterface has methods to work with NetworkAttachmentDefinition resources.
type NetworkAttachmentDefinitionInterface interface {
	Create(ctx context.Context, networkAttachmentDefinition *v1.NetworkAttachmentDefinition, opts metav1.CreateOptions) (*v1.NetworkAttachmentDefinition, error)
	Update(ctx context.Context, networkAttachmentDefinition *v1.NetworkAttachmentDefinition, opts metav1.UpdateOptions) (*v1.NetworkAttachmentDefinition, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.NetworkAttachmentDefinition, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.NetworkAttachmentDefinitionList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NetworkAttachmentDefinition, err error)
	NetworkAttachmentDefinitionExpansion
}

// networkAttachmentDefinitions implements NetworkAttachmentDefinitionInterface
type networkAttachmentDefinitions struct {
	*gentype.ClientWithList[*v1.NetworkAttachmentDefinition, *v1.NetworkAttachmentDefinitionList]
}

// newNetworkAttachmentDefinitions returns a NetworkAttachmentDefinitions
func newNetworkAttachmentDefinitions(c *K8sCniCncfIoV1Client, namespace string) *networkAttachmentDefinitions {
	return &networkAttachmentDefinitions{
		gentype.NewClientWithList[*v1.NetworkAttachmentDefinition, *v1.NetworkAttachmentDefinitionList](
			"network-attachment-definitions",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.NetworkAttachmentDefinition { return &v1.NetworkAttachmentDefinition{} },
			func() *v1.NetworkAttachmentDefinitionList { return &v1.NetworkAttachmentDefinitionList{} }),
	}
}
