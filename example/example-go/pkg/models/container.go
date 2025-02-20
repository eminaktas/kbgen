package models

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type Main struct {
	// +kubebuilder:validation:Optional
	WorkingDir string `json:"workingDir,omitempty" yaml:"workingDir,omitempty"`
	// +kubebuilder:validation:Required
	Name string `json:"name" yaml:"name"`
	// +kubebuilder:validation:Optional
	UseBuiltInEnv *bool `json:"useBuiltInEnv,omitempty" yaml:"useBuiltInEnv,omitempty"`
	// +kubebuilder:validation:Optional
	Args []string `json:"args,omitempty" yaml:"args,omitempty"`
	// +kubebuilder:validation:Optional
	Ports []*ContainerPort `json:"ports,omitempty" yaml:"ports,omitempty"`
	// +kubebuilder:validation:Optional
	Env *EnvMap `json:"env,omitempty" yaml:"env,omitempty"`
	// +kubebuilder:validation:Optional
	StartupProbe *Probe `json:"startupProbe,omitempty" yaml:"startupProbe,omitempty"`
	// +kubebuilder:validation:Optional
	LivenessProbe *Probe `json:"livenessProbe,omitempty" yaml:"livenessProbe,omitempty"`
	// +kubebuilder:validation:Optional
	Command []string `json:"command,omitempty" yaml:"command,omitempty"`
	// +kubebuilder:validation:Optional
	EnvFrom []*EnvFromSource `json:"envFrom,omitempty" yaml:"envFrom,omitempty"`
	// +kubebuilder:validation:Optional
	Lifecycle *Lifecycle `json:"lifecycle,omitempty" yaml:"lifecycle,omitempty"`
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Optional
	SecurityContext map[string]*apiextensionsv1.JSON `json:"securityContext,omitempty" yaml:"securityContext,omitempty"`
	// +kubebuilder:validation:Optional
	ReadinessProbe *Probe `json:"readinessProbe,omitempty" yaml:"readinessProbe,omitempty"`
}
