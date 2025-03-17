package models

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func (in *Exec) DeepCopyInto(out *Exec) {
	*out = *in
}

func (in *Exec) DeepCopy() *Exec {
	if in == nil {
		return nil
	}
	out := new(Exec)
	in.DeepCopyInto(out)
	return out
}

func (in *Http) DeepCopyInto(out *Http) {
	*out = *in
}

func (in *Http) DeepCopy() *Http {
	if in == nil {
		return nil
	}
	out := new(Http)
	in.DeepCopyInto(out)
	return out
}

func (in *Probe) DeepCopyInto(out *Probe) {
	*out = *in
}

func (in *Probe) DeepCopy() *Probe {
	if in == nil {
		return nil
	}
	out := new(Probe)
	in.DeepCopyInto(out)
	return out
}

func (in *Tcp) DeepCopyInto(out *Tcp) {
	*out = *in
}

func (in *Tcp) DeepCopy() *Tcp {
	if in == nil {
		return nil
	}
	out := new(Tcp)
	in.DeepCopyInto(out)
	return out
}

type Exec struct {
	// +required
	Command []string `json:"command" yaml:"command"`
}

type Http struct {
	// +required
	Path string `json:"path" yaml:"path"`
	// +required
	Port *int `json:"port" yaml:"port"`
	// +kubebuilder:validation:Enum=HTTP;HTTPS
	// +required
	Scheme string `json:"scheme" yaml:"scheme"`
}

type Probe struct {
	// +optional
	FailureThreshold *int `json:"failureThreshold,omitempty" yaml:"failureThreshold,omitempty"`
	// +kubebuilder:pruning:PreserveUnknownFields
	// +required
	Handler *apiextensionsv1.JSON `json:"handler" yaml:"handler"`
	// +optional
	InitialDelaySeconds *int `json:"initialDelaySeconds,omitempty" yaml:"initialDelaySeconds,omitempty"`
	// +optional
	PeriodSeconds *int `json:"periodSeconds,omitempty" yaml:"periodSeconds,omitempty"`
	// +optional
	SuccessThreshold *int `json:"successThreshold,omitempty" yaml:"successThreshold,omitempty"`
	// +optional
	TimeoutSeconds *int `json:"timeoutSeconds,omitempty" yaml:"timeoutSeconds,omitempty"`
}

type Tcp struct {
	// +required
	TcpSocket *int `json:"tcpSocket" yaml:"tcpSocket"`
}
