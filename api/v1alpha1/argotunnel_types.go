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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ArgoTunnelSpec defines the desired state of ArgoTunnel
type ArgoTunnelSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	//+kubebuilder:validation:Minimum=0
	//+kubebuilder:default:=1
	//+kubebuilder:validation:Optional
	// Size defines the number of Daemon pods to run for this tunnel
	Size int32 `json:"size,omitempty"`

	//+kubebuilder:default:="cloudflare/cloudflared:2023.7.3"
	//+kubebuilder:validation:Optional
	// Image sets the Cloudflared Image to use. Defaults to the image set during the release of the operator.
	Image string `json:"image,omitempty"`

	//+kubebuilder:validation:Required
	// Cloudflare Credentials
	Cloudflare CloudflareDetails `json:"cloudflare,omitempty"`
}

// ArgoTunnelStatus defines the observed state of ArgoTunnel
type ArgoTunnelStatus struct {
	TunnelId   string `json:"tunnelId"`
	TunnelName string `json:"tunnelName"`
	AccountId  string `json:"accountId"`
	ZoneId     string `json:"zoneId"`
	// +operator-sdk:csv:customresourcedefinitions:type=status
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

// ArgoTunnel is the Schema for the argotunnels API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TunnelID",type=string,JSONPath=`.status.tunnelId`
type ArgoTunnel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ArgoTunnelSpec   `json:"spec,omitempty"`
	Status ArgoTunnelStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ArgoTunnelList contains a list of ArgoTunnel
type ArgoTunnelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ArgoTunnel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ArgoTunnel{}, &ArgoTunnelList{})
}
