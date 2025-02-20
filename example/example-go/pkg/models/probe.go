package models

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type Probe struct {
	// +kubebuilder:validation:Optional
	PeriodSeconds *int `json:"periodSeconds,omitempty" yaml:"periodSeconds,omitempty"`
	// +kubebuilder:validation:Optional
	FailureThreshold *int `json:"failureThreshold,omitempty" yaml:"failureThreshold,omitempty"`
	// +kubebuilder:validation:Optional
	SuccessThreshold *int `json:"successThreshold,omitempty" yaml:"successThreshold,omitempty"`
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Required
	Handler *apiextensionsv1.JSON `json:"handler" yaml:"handler"`
	// +kubebuilder:validation:Optional
	InitialDelaySeconds *int `json:"initialDelaySeconds,omitempty" yaml:"initialDelaySeconds,omitempty"`
	// +kubebuilder:validation:Optional
	TimeoutSeconds *int `json:"timeoutSeconds,omitempty" yaml:"timeoutSeconds,omitempty"`
}

type Exec struct {
	// +kubebuilder:validation:Required
	Command []string `json:"command" yaml:"command"`
}

type Http struct {
	// +kubebuilder:validation:Required
	Port *int `json:"port" yaml:"port"`
	// +kubebuilder:validation:Required
	Path string `json:"path" yaml:"path"`
	// +kubebuilder:validation:Enum=HTTP;HTTPS
	// +kubebuilder:validation:Required
	Scheme string `json:"scheme" yaml:"scheme"`
}

type Tcp struct {
	// +kubebuilder:validation:Required
	TcpSocket *int `json:"tcpSocket" yaml:"tcpSocket"`
}
