package models

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type Lifecycle struct {
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Optional
	PostStart *apiextensionsv1.JSON `json:"postStart,omitempty" yaml:"postStart,omitempty"`
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Optional
	PreStop *apiextensionsv1.JSON `json:"preStop,omitempty" yaml:"preStop,omitempty"`
}
