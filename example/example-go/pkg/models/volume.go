package models

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type Volume struct {
	// +optional
	Mounts []*Mount `json:"mounts,omitempty" yaml:"mounts,omitempty"`
	// +optional
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	VolumeSource *apiextensionsv1.JSON `json:"volumeSource,omitempty" yaml:"volumeSource,omitempty"`
}

func (in *Volume) DeepCopyInto(out *Volume) {
	*out = *in
}

func (in *Volume) DeepCopy() *Volume {
	if in == nil {
		return nil
	}
	out := new(Volume)
	in.DeepCopyInto(out)
	return out
}

type Mount struct {
	// +optional
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
	// +optional
	Container string `json:"container,omitempty" yaml:"container,omitempty"`
	// +optional
	SubPath string `json:"subPath,omitempty" yaml:"subPath,omitempty"`
	// +optional
	ReadOnly *bool `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
}

func (in *Mount) DeepCopyInto(out *Mount) {
	*out = *in
}

func (in *Mount) DeepCopy() *Mount {
	if in == nil {
		return nil
	}
	out := new(Mount)
	in.DeepCopyInto(out)
	return out
}

type EmptyDir struct {
	// +kubebuilder:validation:Enum=;Memory
	// +required
	Medium string `json:"medium" yaml:"medium"`
	// +optional
	SizeLimit string `json:"sizeLimit,omitempty" yaml:"sizeLimit,omitempty"`
}

func (in *EmptyDir) DeepCopyInto(out *EmptyDir) {
	*out = *in
}

func (in *EmptyDir) DeepCopy() *EmptyDir {
	if in == nil {
		return nil
	}
	out := new(EmptyDir)
	in.DeepCopyInto(out)
	return out
}

type Secret struct {
	// +required
	SecretName string `json:"secretName" yaml:"secretName"`
	// +optional
	DefaultMode *int `json:"defaultMode,omitempty" yaml:"defaultMode,omitempty"`
	// +optional
	Items []map[string]string `json:"items,omitempty" yaml:"items,omitempty"`
}

func (in *Secret) DeepCopyInto(out *Secret) {
	*out = *in
}

func (in *Secret) DeepCopy() *Secret {
	if in == nil {
		return nil
	}
	out := new(Secret)
	in.DeepCopyInto(out)
	return out
}

type ConfigMap struct {
	// +optional
	DefaultMode *int `json:"defaultMode,omitempty" yaml:"defaultMode,omitempty"`
	// +optional
	Items []map[string]string `json:"items,omitempty" yaml:"items,omitempty"`
	// +required
	Name string `json:"name" yaml:"name"`
}

func (in *ConfigMap) DeepCopyInto(out *ConfigMap) {
	*out = *in
}

func (in *ConfigMap) DeepCopy() *ConfigMap {
	if in == nil {
		return nil
	}
	out := new(ConfigMap)
	in.DeepCopyInto(out)
	return out
}

type FlexVolume struct {
	// +optional
	ReadOnly *bool `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	// +optional
	Options map[string]string `json:"options,omitempty" yaml:"options,omitempty"`
	// +required
	Driver string `json:"driver" yaml:"driver"`
	// +optional
	FsType string `json:"fsType,omitempty" yaml:"fsType,omitempty"`
}

func (in *FlexVolume) DeepCopyInto(out *FlexVolume) {
	*out = *in
}

func (in *FlexVolume) DeepCopy() *FlexVolume {
	if in == nil {
		return nil
	}
	out := new(FlexVolume)
	in.DeepCopyInto(out)
	return out
}

type HostPath struct {
	// +required
	Path string `json:"path" yaml:"path"`
	// +optional
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

func (in *HostPath) DeepCopyInto(out *HostPath) {
	*out = *in
}

func (in *HostPath) DeepCopy() *HostPath {
	if in == nil {
		return nil
	}
	out := new(HostPath)
	in.DeepCopyInto(out)
	return out
}

type DownwardAPI struct {
	// +optional
	DefaultMode *int `json:"defaultMode,omitempty" yaml:"defaultMode,omitempty"`
	// +optional
	Items []map[string]*apiextensionsv1.JSON `json:"items,omitempty" yaml:"items,omitempty"`
}

func (in *DownwardAPI) DeepCopyInto(out *DownwardAPI) {
	*out = *in
}

func (in *DownwardAPI) DeepCopy() *DownwardAPI {
	if in == nil {
		return nil
	}
	out := new(DownwardAPI)
	in.DeepCopyInto(out)
	return out
}

type CSI struct {
	// +optional
	FsType string `json:"fsType,omitempty" yaml:"fsType,omitempty"`
	// +required
	Driver string `json:"driver" yaml:"driver"`
	// +optional
	VolumeAttributes map[string]string `json:"volumeAttributes,omitempty" yaml:"volumeAttributes,omitempty"`
	// +optional
	ReadOnly *bool `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
}

func (in *CSI) DeepCopyInto(out *CSI) {
	*out = *in
}

func (in *CSI) DeepCopy() *CSI {
	if in == nil {
		return nil
	}
	out := new(CSI)
	in.DeepCopyInto(out)
	return out
}
