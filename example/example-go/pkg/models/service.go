package models

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type Service struct {
	// +optional
	ExternalName string `json:"externalName,omitempty" yaml:"externalName,omitempty"`
	// +optional
	Selector map[string]string `json:"selector,omitempty" yaml:"selector,omitempty"`
	// +optional
	Ports []*ServicePort `json:"ports,omitempty" yaml:"ports,omitempty"`
	// +optional
	PublishNotReadyAddresses *bool `json:"publishNotReadyAddresses,omitempty" yaml:"publishNotReadyAddresses,omitempty"`
	// +optional
	LoadBalancerSourceRanges []string `json:"loadBalancerSourceRanges,omitempty" yaml:"loadBalancerSourceRanges,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// +optional
	ExternalTrafficPolicy string `json:"externalTrafficPolicy,omitempty" yaml:"externalTrafficPolicy,omitempty"`
	// +optional
	IpFamilyPolicy string `json:"ipFamilyPolicy,omitempty" yaml:"ipFamilyPolicy,omitempty"`
	// +optional
	ClusterIP string `json:"clusterIP,omitempty" yaml:"clusterIP,omitempty"`
	// +optional
	LoadBalancerIP string `json:"loadBalancerIP,omitempty" yaml:"loadBalancerIP,omitempty"`
	// +optional
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
	// +optional
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// +optional
	SessionAffinity string `json:"sessionAffinity,omitempty" yaml:"sessionAffinity,omitempty"`
	// +optional
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	// +optional
	ExternalIPs []string `json:"externalIPs,omitempty" yaml:"externalIPs,omitempty"`
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	SessionAffinityConfig map[string]*apiextensionsv1.JSON `json:"sessionAffinityConfig,omitempty" yaml:"sessionAffinityConfig,omitempty"`
	// +optional
	HealthCheckNodePort *int `json:"healthCheckNodePort,omitempty" yaml:"healthCheckNodePort,omitempty"`
	// +optional
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
}

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
