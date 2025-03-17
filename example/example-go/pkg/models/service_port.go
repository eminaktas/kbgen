package models

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type ServicePort struct {
	// +optional
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	// +optional
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	TargetPort *apiextensionsv1.JSON `json:"targetPort,omitempty" yaml:"targetPort,omitempty"`
	// +optional
	AppProtocol string `json:"appProtocol,omitempty" yaml:"appProtocol,omitempty"`
	// +optional
	Port *int `json:"port,omitempty" yaml:"port,omitempty"`
	// +optional
	NodePort *int `json:"nodePort,omitempty" yaml:"nodePort,omitempty"`
}

func (in *ServicePort) DeepCopyInto(out *ServicePort) {
	*out = *in
}

func (in *ServicePort) DeepCopy() *ServicePort {
	if in == nil {
		return nil
	}
	out := new(ServicePort)
	in.DeepCopyInto(out)
	return out
}
