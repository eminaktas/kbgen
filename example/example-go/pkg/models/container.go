package models

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

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

type Main struct {
	// +optional
	Args []string `json:"args,omitempty" yaml:"args,omitempty"`
	// +optional
	Command []string `json:"command,omitempty" yaml:"command,omitempty"`
	// +optional
	Env *EnvMap `json:"env,omitempty" yaml:"env,omitempty"`
	// +optional
	EnvFrom []*EnvFromSource `json:"envFrom,omitempty" yaml:"envFrom,omitempty"`
	// +optional
	Lifecycle *Lifecycle `json:"lifecycle,omitempty" yaml:"lifecycle,omitempty"`
	// +optional
	LivenessProbe *Probe `json:"livenessProbe,omitempty" yaml:"livenessProbe,omitempty"`
	// +required
	Name string `json:"name" yaml:"name"`
	// +optional
	Ports []*ContainerPort `json:"ports,omitempty" yaml:"ports,omitempty"`
	// +optional
	ReadinessProbe *Probe `json:"readinessProbe,omitempty" yaml:"readinessProbe,omitempty"`
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	SecurityContext map[string]*apiextensionsv1.JSON `json:"securityContext,omitempty" yaml:"securityContext,omitempty"`
	// +optional
	StartupProbe *Probe `json:"startupProbe,omitempty" yaml:"startupProbe,omitempty"`
	// +optional
	UseBuiltInEnv *bool `json:"useBuiltInEnv,omitempty" yaml:"useBuiltInEnv,omitempty"`
	// +optional
	WorkingDir string `json:"workingDir,omitempty" yaml:"workingDir,omitempty"`
}
