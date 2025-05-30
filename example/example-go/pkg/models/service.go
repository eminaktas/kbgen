package models

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
}

func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

type Service struct {
	// +optional
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	// +optional
	ClusterIP string `json:"clusterIP,omitempty" yaml:"clusterIP,omitempty"`
	// +optional
	ExternalIPs []string `json:"externalIPs,omitempty" yaml:"externalIPs,omitempty"`
	// +optional
	ExternalName string `json:"externalName,omitempty" yaml:"externalName,omitempty"`
	// +optional
	ExternalTrafficPolicy string `json:"externalTrafficPolicy,omitempty" yaml:"externalTrafficPolicy,omitempty"`
	// +optional
	HealthCheckNodePort *int `json:"healthCheckNodePort,omitempty" yaml:"healthCheckNodePort,omitempty"`
	// +optional
	IpFamilyPolicy string `json:"ipFamilyPolicy,omitempty" yaml:"ipFamilyPolicy,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// +optional
	LoadBalancerIP string `json:"loadBalancerIP,omitempty" yaml:"loadBalancerIP,omitempty"`
	// +optional
	LoadBalancerSourceRanges []string `json:"loadBalancerSourceRanges,omitempty" yaml:"loadBalancerSourceRanges,omitempty"`
	// +optional
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// +optional
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	// +optional
	Ports []*ServicePort `json:"ports,omitempty" yaml:"ports,omitempty"`
	// +optional
	PublishNotReadyAddresses *bool `json:"publishNotReadyAddresses,omitempty" yaml:"publishNotReadyAddresses,omitempty"`
	// +optional
	Selector map[string]string `json:"selector,omitempty" yaml:"selector,omitempty"`
	// +optional
	SessionAffinity string `json:"sessionAffinity,omitempty" yaml:"sessionAffinity,omitempty"`
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	SessionAffinityConfig map[string]*apiextensionsv1.JSON `json:"sessionAffinityConfig,omitempty" yaml:"sessionAffinityConfig,omitempty"`
	// +optional
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
