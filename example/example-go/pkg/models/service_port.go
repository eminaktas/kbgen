package models

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

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

type ServicePort struct {
	// +optional
	AppProtocol string `json:"appProtocol,omitempty" yaml:"appProtocol,omitempty"`
	// +optional
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// +optional
	NodePort *int `json:"nodePort,omitempty" yaml:"nodePort,omitempty"`
	// +optional
	Port *int `json:"port,omitempty" yaml:"port,omitempty"`
	// +optional
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	TargetPort *apiextensionsv1.JSON `json:"targetPort,omitempty" yaml:"targetPort,omitempty"`
}
