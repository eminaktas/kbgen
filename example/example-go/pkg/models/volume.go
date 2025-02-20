package models

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type Volume struct {
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Optional
	VolumeSource *apiextensionsv1.JSON `json:"volumeSource,omitempty" yaml:"volumeSource,omitempty"`
	// +kubebuilder:validation:Optional
	Mounts []*Mount `json:"mounts,omitempty" yaml:"mounts,omitempty"`
	// +kubebuilder:validation:Optional
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

type ConfigMap struct {
	// +kubebuilder:validation:Optional
	Items []map[string]string `json:"items,omitempty" yaml:"items,omitempty"`
	// +kubebuilder:validation:Optional
	DefaultMode *int `json:"defaultMode,omitempty" yaml:"defaultMode,omitempty"`
	// +kubebuilder:validation:Required
	Name string `json:"name" yaml:"name"`
}

type FlexVolume struct {
	// +kubebuilder:validation:Optional
	ReadOnly *bool `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	// +kubebuilder:validation:Optional
	Options map[string]string `json:"options,omitempty" yaml:"options,omitempty"`
	// +kubebuilder:validation:Optional
	FsType string `json:"fsType,omitempty" yaml:"fsType,omitempty"`
	// +kubebuilder:validation:Required
	Driver string `json:"driver" yaml:"driver"`
}

type HostPath struct {
	// +kubebuilder:validation:Required
	Path string `json:"path" yaml:"path"`
	// +kubebuilder:validation:Optional
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

type DownwardAPI struct {
	// +kubebuilder:validation:Optional
	DefaultMode *int `json:"defaultMode,omitempty" yaml:"defaultMode,omitempty"`
	// +kubebuilder:validation:Optional
	Items []map[string]*apiextensionsv1.JSON `json:"items,omitempty" yaml:"items,omitempty"`
}

type CSI struct {
	// +kubebuilder:validation:Optional
	FsType string `json:"fsType,omitempty" yaml:"fsType,omitempty"`
	// +kubebuilder:validation:Optional
	VolumeAttributes map[string]string `json:"volumeAttributes,omitempty" yaml:"volumeAttributes,omitempty"`
	// +kubebuilder:validation:Optional
	ReadOnly *bool `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	// +kubebuilder:validation:Required
	Driver string `json:"driver" yaml:"driver"`
}

type EmptyDir struct {
	// +kubebuilder:validation:Enum=;Memory
	// +kubebuilder:validation:Required
	Medium string `json:"medium" yaml:"medium"`
	// +kubebuilder:validation:Optional
	SizeLimit string `json:"sizeLimit,omitempty" yaml:"sizeLimit,omitempty"`
}

type Secret struct {
	// +kubebuilder:validation:Required
	SecretName string `json:"secretName" yaml:"secretName"`
	// +kubebuilder:validation:Optional
	Items []map[string]string `json:"items,omitempty" yaml:"items,omitempty"`
	// +kubebuilder:validation:Optional
	DefaultMode *int `json:"defaultMode,omitempty" yaml:"defaultMode,omitempty"`
}

type Mount struct {
	// +kubebuilder:validation:Optional
	SubPath string `json:"subPath,omitempty" yaml:"subPath,omitempty"`
	// +kubebuilder:validation:Optional
	Container string `json:"container,omitempty" yaml:"container,omitempty"`
	// +kubebuilder:validation:Optional
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
	// +kubebuilder:validation:Optional
	ReadOnly *bool `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
}
