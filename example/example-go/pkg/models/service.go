package models

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type Service struct {
	// +kubebuilder:validation:Optional
	ClusterIP string `json:"clusterIP,omitempty" yaml:"clusterIP,omitempty"`
	// +kubebuilder:validation:Optional
	SessionAffinity string `json:"sessionAffinity,omitempty" yaml:"sessionAffinity,omitempty"`
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// +kubebuilder:validation:Optional
	LoadBalancerIP string `json:"loadBalancerIP,omitempty" yaml:"loadBalancerIP,omitempty"`
	// +kubebuilder:validation:Optional
	ExternalIPs []string `json:"externalIPs,omitempty" yaml:"externalIPs,omitempty"`
	// +kubebuilder:validation:Optional
	Selector map[string]string `json:"selector,omitempty" yaml:"selector,omitempty"`
	// +kubebuilder:validation:Optional
	IpFamilyPolicy string `json:"ipFamilyPolicy,omitempty" yaml:"ipFamilyPolicy,omitempty"`
	// +kubebuilder:validation:Optional
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	// +kubebuilder:validation:Optional
	Ports []*ServicePort `json:"ports,omitempty" yaml:"ports,omitempty"`
	// +kubebuilder:validation:Optional
	ExternalName string `json:"externalName,omitempty" yaml:"externalName,omitempty"`
	// +kubebuilder:validation:Optional
	PublishNotReadyAddresses *bool `json:"publishNotReadyAddresses,omitempty" yaml:"publishNotReadyAddresses,omitempty"`
	// +kubebuilder:validation:Optional
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
	// +kubebuilder:validation:Optional
	ExternalTrafficPolicy string `json:"externalTrafficPolicy,omitempty" yaml:"externalTrafficPolicy,omitempty"`
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Optional
	SessionAffinityConfig map[string]*apiextensionsv1.JSON `json:"sessionAffinityConfig,omitempty" yaml:"sessionAffinityConfig,omitempty"`
	// +kubebuilder:validation:Optional
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	// +kubebuilder:validation:Optional
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// +kubebuilder:validation:Optional
	HealthCheckNodePort *int `json:"healthCheckNodePort,omitempty" yaml:"healthCheckNodePort,omitempty"`
	// +kubebuilder:validation:Optional
	LoadBalancerSourceRanges []string `json:"loadBalancerSourceRanges,omitempty" yaml:"loadBalancerSourceRanges,omitempty"`
}
