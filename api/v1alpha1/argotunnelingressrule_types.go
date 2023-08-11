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

type ArgoTunnelRef struct {

	// The name of the referenced ArgoTunnel
	//+kubebuilder:validation:Required
	Name string `json:"name"`
	// The namespace of the referenced ArgoTunnel (optional)
	//+kubebuilder:validation:Optional
	NameSpace string `json:"namespace,omitempty"`
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ArgoTunnelIngressRuleSpec defines the desired state of ArgoTunnelIngressRule
type ArgoTunnelIngressRuleSpec struct {
	// Fqdn specifies the DNS name to access this service from.
	// Defaults to the service.metadata.name + tunnel.spec.domain.
	// If specifying this, make sure to use the same domain that the tunnel belongs to.
	// This is not validated and used as provided
	//+kubebuilder:validation:Optional
	Fqdn string `json:"fqdn,omitempty"`

	// ArgoTunnelRef specifies the referenced ArgoTunnel.
	// Defaults to the service.metadata.name + tunnel.spec.domain.
	// If specifying this, make sure to use the same domain that the tunnel belongs to.
	// This is not validated and used as provided
	//+kubebuilder:validation:Optional
	ArgoTunnelRef ArgoTunnelRef `json:"argoTunnelRef,omitempty"`

	// Protocol specifies the protocol for the service. Should be one of http, https, tcp, udp, ssh or rdp.
	// Defaults to http, with the exceptions of https for 443, smb for 139 and 445, rdp for 3389 and ssh for 22 if the service has a TCP port.
	// The only available option for a UDP port is udp, which is default.
	//+kubebuilder:validation:Optional
	Protocol string `json:"protocol,omitempty"`

	// Target specified where the tunnel should proxy to.
	// Defaults to the form of <protocol>://<service.metadata.name>.<service.metadata.namespace>.svc:<port>
	//+kubebuilder:validation:Optional
	Target string `json:"target,omitempty"`

	// CaPool trusts the CA certificate referenced by the key in the secret specified in tunnel.spec.originCaPool.
	// tls.crt is trusted globally and does not need to be specified. Only useful if the protocol is HTTPS.
	//+kubebuilder:validation:Optional
	CaPool string `json:"caPool,omitempty"`

	// NoTlsVerify disables TLS verification for this service.
	// Only useful if the protocol is HTTPS.
	//+kubebuilder:validation:Optional
	//+kubebuilder:default:=false
	NoTlsVerify bool `json:"noTlsVerify"`

	// cloudflared starts a proxy server to translate HTTP traffic into TCP when proxying, for example, SSH or RDP.

	// ProxyAddress configures the listen address for that proxy
	//+kubebuilder:validation:Optional
	//+kubebuilder:default:="127.0.0.1"
	//+kubebuilder:validation:Pattern="((^((([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]))$)|(^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$))"
	ProxyAddress string `json:"proxyAddress,omitempty"`

	// ProxyPort configures the listen port for that proxy
	//+kubebuilder:validation:Optional
	//+kubebuilder:default:=0
	//+kubebuilder:validation:Minimum:=0
	//+kubebuilder:validation:Maximum:=65535
	ProxyPort uint `json:"proxyPort,omitempty"`

	// ProxyType configures the proxy type.
	//+kubebuilder:validation:Optional
	//+kubebuilder:default:=""
	//+kubebuilder:validation:Enum:="";"socks"
	ProxyType string `json:"proxyType,omitempty"`
}

// ArgoTunnelIngressRuleStatus defines the observed state of ArgoTunnelIngressRule
type ArgoTunnelIngressRuleStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ArgoTunnelIngressRule is the Schema for the argotunnelingressrules API
type ArgoTunnelIngressRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ArgoTunnelIngressRuleSpec   `json:"spec,omitempty"`
	Status ArgoTunnelIngressRuleStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ArgoTunnelIngressRuleList contains a list of ArgoTunnelIngressRule
type ArgoTunnelIngressRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ArgoTunnelIngressRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ArgoTunnelIngressRule{}, &ArgoTunnelIngressRuleList{})
}
