package models

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type ServicePort struct {
	// +kubebuilder:validation:Optional
	NodePort *int `json:"nodePort,omitempty" yaml:"nodePort,omitempty"`
	// +kubebuilder:validation:Optional
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	// +kubebuilder:validation:Optional
	AppProtocol string `json:"appProtocol,omitempty" yaml:"appProtocol,omitempty"`
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Optional
	TargetPort *apiextensionsv1.JSON `json:"targetPort,omitempty" yaml:"targetPort,omitempty"`
	// +kubebuilder:validation:Optional
	Port *int `json:"port,omitempty" yaml:"port,omitempty"`
	// +kubebuilder:validation:Optional
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
