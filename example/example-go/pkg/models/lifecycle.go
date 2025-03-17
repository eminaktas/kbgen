package models

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type Lifecycle struct {
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	PostStart *apiextensionsv1.JSON `json:"postStart,omitempty" yaml:"postStart,omitempty"`
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	PreStop *apiextensionsv1.JSON `json:"preStop,omitempty" yaml:"preStop,omitempty"`
}

func (in *Lifecycle) DeepCopyInto(out *Lifecycle) {
	*out = *in
}

func (in *Lifecycle) DeepCopy() *Lifecycle {
	if in == nil {
		return nil
	}
	out := new(Lifecycle)
	in.DeepCopyInto(out)
	return out
}
