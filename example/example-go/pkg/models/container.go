package models

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type Main struct {
	// +required
	Name string `json:"name" yaml:"name"`
	// +optional
	ReadinessProbe *Probe `json:"readinessProbe,omitempty" yaml:"readinessProbe,omitempty"`
	// +optional
	Lifecycle *Lifecycle `json:"lifecycle,omitempty" yaml:"lifecycle,omitempty"`
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	SecurityContext map[string]*apiextensionsv1.JSON `json:"securityContext,omitempty" yaml:"securityContext,omitempty"`
	// +optional
	Ports []*ContainerPort `json:"ports,omitempty" yaml:"ports,omitempty"`
	// +optional
	Command []string `json:"command,omitempty" yaml:"command,omitempty"`
	// +optional
	StartupProbe *Probe `json:"startupProbe,omitempty" yaml:"startupProbe,omitempty"`
	// +optional
	Env *EnvMap `json:"env,omitempty" yaml:"env,omitempty"`
	// +optional
	WorkingDir string `json:"workingDir,omitempty" yaml:"workingDir,omitempty"`
	// +optional
	UseBuiltInEnv *bool `json:"useBuiltInEnv,omitempty" yaml:"useBuiltInEnv,omitempty"`
	// +optional
	EnvFrom []*EnvFromSource `json:"envFrom,omitempty" yaml:"envFrom,omitempty"`
	// +optional
	LivenessProbe *Probe `json:"livenessProbe,omitempty" yaml:"livenessProbe,omitempty"`
	// +optional
	Args []string `json:"args,omitempty" yaml:"args,omitempty"`
}

func (in *Main) DeepCopyInto(out *Main) {
	*out = *in
}

func (in *Main) DeepCopy() *Main {
	if in == nil {
		return nil
	}
	out := new(Main)
	in.DeepCopyInto(out)
	return out
}
